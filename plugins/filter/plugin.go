package filter

import (
	"aw-sync-agent-plugins/models"
	"log"
	"sort"
)

type Plugin struct{}

var config models.FilterConfig

func (f *Plugin) Initialize() {
	config = LoadYAMLConfig("./config/" + f.Name())
	if config.Filters != nil {
		log.Print("Validating Filters...")
		var totalFilters, invalidFilters, disabledFilters int
		config.Filters, totalFilters, invalidFilters, disabledFilters = ValidateFilters(config.Filters)
		PrintFilters(totalFilters, invalidFilters, disabledFilters)

		log.Print("Extracting Categories from Filters...")
		categories := GetCategories(config.Filters)
		if len(categories) > 0 {
			PrintCategories(categories)
		} else {
			log.Print("No Categories found.")
		}
	}

}

func (f *Plugin) Execute(watcher string, events models.Events, userID string, includeHostName bool) {
	// Implementation
	//Apply the filters
	var modifiedEvents models.Events
	var watcherFilters []models.Filter
	if watcher != "aw-watcher-afk" {
		watcherFilters = GetMatchingFilters(config.Filters, watcher)
		// Sort watcherFilters so filters with a Category take priority
		sort.Slice(watcherFilters, func(i, j int) bool {
			return watcherFilters[i].Category != "" && watcherFilters[j].Category == ""
		})
	}

	var dropEvent bool
	for _, event := range events {

		// Here it will be the abstract run of each plugin.We can follow strict order of execution.Each plugin will have its own function and must return Event type.

		//Apply the filters
		if watcher != "aw-watcher-afk" {
			event.Data["category"] = "Other" //Default category
			event.Data, dropEvent = Apply(event.Data, watcherFilters)
		}

		// Drop the event if it matches the filter
		if dropEvent {
			continue
		}
		modifiedEvents = append(modifiedEvents, event)
	}
}

func (f *Plugin) ReplicateConfig(path string) {
	err := CreateConfigFile(path, f.Name())
	log.Print(err)
}

func (f *Plugin) Name() string {
	return "aw-plugin-" + f.RawName() + ".yml"
}

func (f *Plugin) RawName() string {
	return "filters"
}
