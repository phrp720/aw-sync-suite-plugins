package plugins

import (
	"github.com/phrp720/aw-sync-agent-plugins/models"
	"github.com/phrp720/aw-sync-agent-plugins/plugins/filter"
)

func Initialize() []models.Plugin {
	var plugins []models.Plugin
	plugins = append(plugins, &filter.Plugin{})

	return plugins
}

func Select(plugins []models.Plugin, names []string) []models.Plugin {

	var selectedPlugins []models.Plugin
	nameSet := make(map[string]struct{})
	for _, name := range names {
		nameSet[name] = struct{}{}
	}

	for _, plugin := range plugins {
		if _, exists := nameSet[plugin.RawName()]; exists {
			selectedPlugins = append(selectedPlugins, plugin)
		}
	}
	return selectedPlugins
}
