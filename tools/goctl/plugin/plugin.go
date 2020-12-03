package plugin

import (
	"encoding/json"
	"errors"
	"github.com/urfave/cli"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser"
	"github.com/tal-tech/go-zero/tools/goctl/rpc/execx"
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/tal-tech/go-zero/tools/goctl/util/ctx"
)

const defaultPluginName = "_plugin"

func Do(c *cli.Context) error {
	apiFile := c.String("api")
	dir := c.String("dir")
	namingStyle := c.String("style")
	plgArg := c.String("plugin")

	pluginArgs, err := prepareArgs(apiFile, dir, namingStyle)
	if err != nil {
		return err
	}

	bin, download, err := getCommand(plgArg)
	if err != nil {
		return err
	}
	if download {
		defer func() {
			_ = os.Remove(bin)
		}()
	}

	var commands []string
	commands = append(commands, bin)
	commands = append(commands, pluginArgs...)

	cmd := strings.Join(commands, " ")
	println(cmd)
	_, err = execx.Run(cmd, "")
	return err
}

func prepareArgs(apiPath, dir, namingStyle string) ([]string, error) {
	var pluginArgs []string
	if len(apiPath) > 0 && util.FileExists(apiPath) {
		p, err := parser.NewParser(apiPath)
		if err != nil {
			return nil, err
		}

		api, err := p.Parse()
		if err != nil {
			return nil, err
		}

		data, err := json.Marshal(api)
		if err != nil {
			return nil, err
		}
		pluginArgs = append(pluginArgs, "-spec")
		pluginArgs = append(pluginArgs, string(data))
	}

	if len(dir) > 0 {
		abs, err := filepath.Abs(dir)
		if err != nil {
			return nil, err
		}

		projectCtx, err := ctx.Prepare(abs)
		if err != nil {
			return nil, err
		}

		data, err := json.Marshal(projectCtx)
		if err != nil {
			return nil, err
		}
		pluginArgs = append(pluginArgs, "-context")
		pluginArgs = append(pluginArgs, string(data))
	}

	return pluginArgs, nil
}

func getCommand(arg string) (string, bool, error) {
	if _, err := exec.LookPath(arg); err == nil {
		return arg, false, nil
	}

	if strings.HasPrefix(arg, "http") {
		err := downloadFile(defaultPluginName, arg)
		if err != nil {
			return "", false, err
		}
		return defaultPluginName, true, nil
	}
	return "", false, errors.New("invalid plugin value " + arg)
}

func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
