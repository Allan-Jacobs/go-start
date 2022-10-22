package ide_plugin

import "github.com/Allan-Jacobs/go-start/plugin"

var IdePlugin = plugin.Builder().
	PostGenerationFeature().
	WithName("Open In VSCode").
	WithDescription("Open the generated project in VSCode").
	WithPostGenerationAction(
		plugin.ConfirmAndThenActions("Open VSCode", plugin.CommandAction("code", ".", "-r", "-g", "main.go:4")),
	).
	WithAvailabilityFilter(plugin.HasExec("code")).
	AddFeature().
	Build()
