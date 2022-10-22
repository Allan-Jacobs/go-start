package git_plugin

import "github.com/Allan-Jacobs/go-start/plugin"

var Git = plugin.Builder().
	PostGenerationFeature().
	WithName("Initialize Git").
	WithDescription("Initialize a git repository").
	WithPostGenerationAction(
		plugin.ConfirmAndThenActions("Initialize git repo", plugin.CommandAction("git", "init", "-q"), plugin.CommandAction("git", "add", ".")),
	).
	WithAvailabilityFilter(plugin.HasExec("git")).
	AddFeature().
	Build()
