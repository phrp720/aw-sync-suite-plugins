package script

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/phrp720/aw-sync-agent-plugins/models"
	"github.com/phrp720/aw-sync-agent-plugins/util"
	"log"
	"os/exec"
	"time"
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
		log.Printf("###Scripts Prompt###")
	}
	// Convert events to JSON
	eventsToJSON, err := json.Marshal(events)
	if err != nil {
		log.Printf("Error marshalling events: %v", err)
		return events
	}

	for _, script := range config.Scripts {

		if !util.FileExists(script.Path) {
			log.Printf("Error running %s : %s", script.Name, "No such file or directory.")
			continue
		} else {
			log.Printf("Running %s", script.Name)

		}

		// Set a default timeout of 30 seconds
		if script.Timeout == 0 {
			script.Timeout = 30
		}

		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(script.Timeout)*time.Second)
		defer cancel()

		cmd := exec.CommandContext(ctx, script.Path)
		// Set up stdin and stdout
		cmd.Stdin = bytes.NewReader(eventsToJSON)
		var stdout bytes.Buffer
		cmd.Stdout = &stdout

		// Run the command
		err = cmd.Run()
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Printf("Error running %s : %s", script.Name, "Timeout exceeded")
			return events

		}
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
		events = EventsBuffer
		log.Printf("Finished %s ", script.Name)

	}
	if config.Scripts != nil {
		log.Printf("####################")
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
