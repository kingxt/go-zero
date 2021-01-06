package javagen

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
	"github.com/tal-tech/go-zero/tools/goctl/util"
)

const getSetTemplate = `
{{.indent}}{{.decorator}}
{{.indent}}public {{.returnType}} get{{.property}}() {
{{.indent}}	return this.{{.propertyValue}};
{{.indent}}}

{{.indent}}public void set{{.property}}({{.type}} {{.propertyValue}}) {
{{.indent}}	this.{{.propertyValue}} = {{.propertyValue}};
{{.indent}}}
`

func writeProperty(writer io.Writer, member spec.Member, indent int) error {
	writeIndent(writer, indent)
	ty, err := goTypeToJava(member.Type)
	ty = strings.Replace(ty, "*", "", 1)
	if err != nil {
		return err
	}

	name, err := member.GetPropertyName()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(writer, "private %s %s", ty, name)
	if err != nil {
		return err
	}

	writeDefaultValue(writer, member)
	fmt.Fprint(writer, ";\n")
	return err
}

func writeDefaultValue(writer io.Writer, member spec.Member) error {
	javaType, err := goTypeToJava(member.Type)
	if err != nil {
		return err
	}

	if javaType == "String" {
		_, err := fmt.Fprintf(writer, " = \"\"")
		return err
	}
	return nil
}

func writeIndent(writer io.Writer, indent int) {
	for i := 0; i < indent; i++ {
		fmt.Fprint(writer, "\t")
	}
}

func indentString(indent int) string {
	var result = ""
	for i := 0; i < indent; i++ {
		result += "\t"
	}
	return result
}

func goTypeToJava(tp spec.Type) (string, error) {
	switch v := tp.(type) {
	case spec.DefineStruct:
		return util.Title(tp.Name()), nil
	case spec.PrimitiveType:
		r, ok := primitiveType(tp.Name())
		if !ok {
			return "", errors.New("unsupported primitive type " + tp.Name())
		}
		return r, nil
	case spec.MapType:
		valueType, err := goTypeToJava(v.Value)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("java.util.HashMap<String, %s>", util.Title(valueType)), nil
	case spec.ArrayType:
		valueType, err := goTypeToJava(v.Value)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("java.util.ArrayList<%s>", util.Title(valueType)), nil
	case spec.InterfaceType:
		return "Object", nil
	case spec.PointerType:
		return goTypeToJava(v.Type)
	}

	return "", errors.New("unsupported primitive type " + tp.Name())
}

func primitiveType(tp string) (string, bool) {
	switch tp {
	case "string":
		return "String", true
	case "int64":
		return "long", true
	case "int", "int8", "int32":
		return "int", true
	case "float", "float32", "float64":
		return "double", true
	case "bool":
		return "boolean", true
	}

	return "", false
}

func genGetSet(writer io.Writer, defineStruct spec.DefineStruct, indent int) error {
	t := template.Must(template.New("getSetTemplate").Parse(getSetTemplate))
	for _, member := range defineStruct.Members {
		var tmplBytes bytes.Buffer

		javaType, err := goTypeToJava(member.Type)
		if err != nil {
			return err
		}

		tyString := javaType
		decorator := ""
		javaPrimitiveType := []string{"int", "long", "boolean", "float", "double", "short"}
		if !stringx.Contains(javaPrimitiveType, javaType) {
			if member.IsOptional() {
				decorator = "@Nullable "
			} else {
				decorator = "@NotNull "
			}
			tyString = decorator + tyString
		}

		err = t.Execute(&tmplBytes, map[string]string{
			"property":      util.Title(member.Name),
			"propertyValue": util.Untitle(member.Name),
			"type":          tyString,
			"decorator":     decorator,
			"returnType":    javaType,
			"indent":        indentString(indent),
		})
		if err != nil {
			return err
		}

		r := tmplBytes.String()
		r = strings.Replace(r, " boolean get", " boolean is", 1)
		writer.Write([]byte(r))
	}
	return nil
}
