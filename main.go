package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Allan-Jacobs/go-start/plugin"
	"github.com/Allan-Jacobs/go-start/plugins/git_plugin"
	"github.com/Allan-Jacobs/go-start/plugins/ide_plugin"
	"github.com/Allan-Jacobs/go-start/plugins/template_plugin"
	"github.com/manifoldco/promptui"
)

func main() {
	engine := plugin.NewEngine(git_plugin.Git, ide_plugin.IdePlugin, template_plugin.TemplatePlugin)

	if len(os.Args) == 2 {
		if os.Args[1] == "--config" {
			interactive_config()
		} else {
			var project_dir = os.Args[1]

			err := engine.Run(project_dir)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println("Usage: [project dir] or --config")
	}
}

func interactive_config() {
	config, err := plugin.GetConfig()
	if err != nil {
		fmt.Printf("An error occurred when trying to open the config: %v\n", err)
		return
	}
	// Todo: menu
	prompt := promptui.Prompt{
		Label:     "Module URL Base: ",
		Default:   config.ModuleUrlBase,
		AllowEdit: true,
	}

	res, err := prompt.Run()
	if err != nil {
		if errors.Is(err, promptui.ErrAbort) || errors.Is(err, promptui.ErrInterrupt) {
			fmt.Println("Changes Canceled")
			return
		}
		fmt.Printf("A Validation error occurred: %v\n", err)
		return
	}
	config.ModuleUrlBase = res
	err = plugin.SetConfig(config)
	if err != nil {
		fmt.Printf("An error occurred while trying to save the config: %v\n", err)
		return
	}
	fmt.Println("Config Saved Successfully")
}
