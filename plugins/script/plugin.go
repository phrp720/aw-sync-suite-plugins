package script

import (
	"github.com/phrp720/aw-sync-agent-plugins/models"
)

type Plugin struct{}

func (f *Plugin) Initialize() {

}

func (f *Plugin) Execute(events models.Events, watcher string, userID string, includeHostName bool) models.Events {

	return events
}

func (f *Plugin) ReplicateConfig(path string) {

}

func (f *Plugin) Name() string {
	return "aw-plugin-" + f.RawName() + ".yaml"
}

func (f *Plugin) RawName() string {
	return "script"
}
