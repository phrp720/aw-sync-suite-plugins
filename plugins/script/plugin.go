package script

import (
	"bytes"
	"encoding/json"
	"github.com/phrp720/aw-sync-agent-plugins/models"
	"github.com/phrp720/aw-sync-agent-plugins/util"
	"log"
	"os/exec"
)

var config models.ScriptConfig

type Plugin struct{}

func (f *Plugin) Initialize() {
	config = LoadYAMLConfig("./config/" + f.Name())
	if config.Scripts != nil {
		PrintScripts(GetScriptNames(config.Scripts))
	}
}

func (f *Plugin) Execute(events models.Events, watcher string, userID string, includeHostName bool) models.Events {

	if config.Scripts == nil {
		return events
	} else {
		log.Printf("Scripts")
	}
	// Convert events to JSON
	eventsToJSON, err := json.Marshal(events)
	if err != nil {
		log.Printf("Error marshalling events: %v", err)
		return events
	}

	for _, script := range config.Scripts {
		AbsPath := script.Path + script.Name

		if !util.FileExists(AbsPath) {
			log.Printf("Error running %s : %s", script.Name, "No such file or directory.")
			continue
		} else {
			log.Printf("Running %s", script.Name)

		}
		cmd := exec.Command(AbsPath)
		// Set up stdin and stdout
		cmd.Stdin = bytes.NewReader(eventsToJSON)
		var stdout bytes.Buffer
		cmd.Stdout = &stdout

		// Run the command
		err = cmd.Run()
		if err != nil {
			log.Printf("Error running %s : %v", script.Name, err)
			return events
		}
		// Read the output and convert it back to events
		var EventsBuffer models.Events
		err = json.Unmarshal(stdout.Bytes(), &EventsBuffer)
		if err != nil {
			log.Printf("Error unmarshalling script output: %v", err)
			return events
		}
		if EventsBuffer != nil {
			events = EventsBuffer
			log.Printf("Finished %s ", script.Name)

		}

	}

	return events
}

func (f *Plugin) ReplicateConfig(path string) {
	err := CreateConfigFile(path, f.Name())
	if err != nil {
		log.Print(err)
	}
}

func (f *Plugin) Name() string {
	return "aw-plugin-" + f.RawName() + ".yaml"
}

func (f *Plugin) RawName() string {
	return models.SCRIPT
}
