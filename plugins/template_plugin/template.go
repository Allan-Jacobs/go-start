package template_plugin

import (
	"errors"
	"regexp"
	"strings"
	"text/template"

	"github.com/Allan-Jacobs/go-start/plugin"
	"github.com/lithammer/dedent"
	"github.com/manifoldco/promptui"
)

var r, _ = regexp.Compile(`^[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,4}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$`)
var errInvalidModuleURL = errors.New("invalid module URL")

var TemplatePlugin = plugin.Builder().
	TemplateFeature().
	WithName("default").
	WithDescription("the default template").
	WithGetTemplateData(func() (any, error) {
		prompt := promptui.Prompt{
			Label: "Module URL",
			Validate: func(s string) error {
				if !r.MatchString(s) {
					return errInvalidModuleURL
				}
				return nil
			},
		}
		module, err := prompt.Run()
		if err != nil {
			return nil, err
		}
		data := struct {
			Version, Module string
		}{
			Version: "1.18",
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
	WithAvailabilityFilter(func() bool { return true }).
	AddFeature().
	Build()
