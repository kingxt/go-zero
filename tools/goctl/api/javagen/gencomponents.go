package javagen

import (
	"errors"
	"fmt"
	"io"
	"path"
	"strings"
	"text/template"

	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
	apiutil "github.com/tal-tech/go-zero/tools/goctl/api/util"
	"github.com/tal-tech/go-zero/tools/goctl/util"
)

const (
	componentTemplate = `// Code generated by goctl. DO NOT EDIT.
package com.xhb.logic.http.packet.{{.packet}}.model;
import com.xhb.logic.http.DeProguardable;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
{{.componentType}}
`
)

func genComponents(dir, packetName string, api *spec.ApiSpec) error {
	types := api.Types
	if len(types) == 0 {
		return nil
	}

	for _, ty := range types {
		if err := createComponent(dir, packetName, ty); err != nil {
			return err
		}
	}

	return nil
}

func createComponent(dir, packetName string, ty spec.Type) error {
	defineStruct, ok := ty.(spec.DefineStruct)
	if !ok {
		return errors.New("unsupported type %s" + ty.Name())
	}

	if len(defineStruct.GetBodyMembers()) == 0 {
		return nil
	}

	modelFile := util.Title(ty.Name()) + ".java"
	filename := path.Join(dir, modelDir, modelFile)
	if err := util.RemoveOrQuit(filename); err != nil {
		return err
	}

	fp, created, err := apiutil.MaybeCreateFile(dir, modelDir, modelFile)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}
	defer fp.Close()

	tyString, err := buildType(defineStruct)
	if err != nil {
		return err
	}

	t := template.Must(template.New("componentType").Parse(componentTemplate))
	return t.Execute(fp, map[string]string{
		"componentType": tyString,
		"packet":        packetName,
	})
}

func buildType(ty spec.DefineStruct) (string, error) {
	var builder strings.Builder
	if err := writeType(&builder, ty); err != nil {
		return "", apiutil.WrapErr(err, "Type "+ty.Name()+" generate error")
	}
	return builder.String(), nil
}

func writeType(writer io.Writer, defineStruct spec.DefineStruct) error {
	fmt.Fprintf(writer, "public class %s extends HttpData {\n", util.Title(defineStruct.Name()))
	err := writeMembers(writer, defineStruct, 1)
	if err != nil {
		return err
	}

	genGetSet(writer, defineStruct, 1)
	fmt.Fprintf(writer, "}")
	return nil
}

func writeMembers(writer io.Writer, ty spec.DefineStruct, indent int) error {
	for _, member := range ty.Members {
		if member.IsInline {
			defineStruct, ok := member.Type.(spec.DefineStruct)
			if ok {
				err := writeMembers(writer, defineStruct, indent)
				if err != nil {
					return err
				}
				continue
			}
			return errors.New("unsupported inline type %s" + member.Type.Name())
		}

		if !member.IsBodyMember() {
			continue
		}

		if err := writeProperty(writer, member, indent); err != nil {
			return err
		}
	}
	return nil
}
