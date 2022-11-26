package template_plugin

import (
	"errors"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"text/template"

	"github.com/Allan-Jacobs/go-start/plugin"
	"github.com/lithammer/dedent"
	"github.com/manifoldco/promptui"
)

var module_regexp = *regexp.MustCompile(`^[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,4}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$`)
var version_regexp = *regexp.MustCompile(`go version go(?P<version>\d+\.\d+)\.\d+ (?:\w+)\/(?:\w+)`)
var errInvalidModuleURL = errors.New("invalid module URL")

var TemplatePlugin = plugin.Builder().
	TemplateFeature().
	WithName("default").
	WithDescription("the default template").
	WithGetTemplateData(func(ctx plugin.TemplateContext) (any, error) {

		name := path.Base(ctx.ProjectDir)

		default_module_url := path.Join(ctx.Config.ModuleUrlBase, name)

		prompt := promptui.Prompt{
			Label: "Module URL",
			Validate: func(s string) error {
				if !module_regexp.MatchString(s) {
					return errInvalidModuleURL
				}
				return nil
			},
			Default: default_module_url,
		}
		module, err := prompt.Run()
		if err != nil {
			return nil, err
		}

		cmd := exec.Command("go", "version")
		out, err := cmd.Output()
		if err != nil {
			return nil, err
		}

		matches := version_regexp.FindStringSubmatch(string(out))
		go_version := matches[version_regexp.SubexpIndex("version")]

		data := struct {
			Version, Module string
		}{
			Version: go_version,
			Module:  module,
		}
		return data, nil
	}).
	WithNewTemplate(*template.Must(template.New("main.go").
		Parse(strings.Trim(dedent.Dedent(`
		package main
	
		func main() {
			
		}
		`), "\n")))).
	WithNewTemplate(*template.Must(template.New("go.mod").
		Parse(strings.Trim(dedent.Dedent(`
		module {{.Module}}

		go {{.Version}}
		`), "\n")))).
	WithAvailabilityFilter(plugin.HasExec("go")).
	WithEntryPoint(func(ctx plugin.TemplateContext) plugin.EntryPoint {
		return plugin.EntryPoint{
			Path:  "main.go",
			Line:  4,
			IsDir: false,
		}
	}).
	AddFeature().
	Build()
