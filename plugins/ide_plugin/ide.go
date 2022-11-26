package ide_plugin

import (
	"errors"
	"strconv"

	"github.com/Allan-Jacobs/go-start/plugin"
	"github.com/manifoldco/promptui"
)

var IdePlugin = plugin.Builder().
	PostGenerationFeature().
	WithName("Open In IDE").
	WithDescription("Open the generated project in The Selected IDE").
	WithPostGenerationAction(func(ctx plugin.PostGenerationContext) error {
		hasVSCode := plugin.HasExec("code")()
		hasGoLand := plugin.HasExec("goland.sh")()

		line := strconv.FormatInt(int64(ctx.EntryPoint.Line), 10)

		openInVSCode := plugin.CommandAction("code", ".", "-r", "-g", ctx.EntryPoint.Path+":"+line)
		openInGoLand := plugin.CommandAction("goland.sh", "--line", line, ctx.EntryPoint.Path)

		if hasVSCode && !hasGoLand {
			return plugin.ConfirmAndThenActions("Open VSCode", openInVSCode)()
		} else if !hasVSCode && hasGoLand {
			return plugin.ConfirmAndThenActions("Open GoLand", openInGoLand)()
		} else if hasVSCode && hasGoLand {
			prompt := promptui.Select{
				Label: "Open project in",
				Items: []string{"VSCode", "GoLand"},
			}
			_, value, err := prompt.Run()
			if err != nil {
				return err
			}
			if value == "VSCode" {
				return openInVSCode()
			} else if value == "GoLand" {
				return openInGoLand()
			} else {
				return errors.New("invalid IDE Selected")
			}
		}
		return errors.New("no IDE Found, but the filter returned true. This is probably a bug in go-start")
	}).
	WithAvailabilityFilter(plugin.HasExec("code").Or(plugin.HasExec("goland.sh"))). // vscode or goland
	AddFeature().
	Build()
