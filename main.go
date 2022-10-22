package main

import (
	"fmt"
	"os"

	"github.com/Allan-Jacobs/go-start/plugin"
	"github.com/Allan-Jacobs/go-start/plugins/git_plugin"
	"github.com/Allan-Jacobs/go-start/plugins/ide_plugin"
	"github.com/Allan-Jacobs/go-start/plugins/template_plugin"
)

func main() {
	engine := plugin.NewEngine(git_plugin.Git, ide_plugin.IdePlugin, template_plugin.TemplatePlugin)

	if len(os.Args) == 2 {
		var project_dir = os.Args[1]

		err := engine.Run(project_dir)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		fmt.Println("Usage: [project dir]")
	}

}
