package filter

import (
	"aw-sync-agent-plugins/models"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"path/filepath"
)

// LoadYAMLConfig  Load the YAML config file
func LoadYAMLConfig(filename string) models.FilterConfig {
	file, err := os.Open(filename)

	if err != nil {
		log.Printf("No %s file found.", filename)
	} else {
		log.Printf("Loading filters from %s file.", filename)
		defer file.Close()
		decoder := yaml.NewDecoder(file)
		if err = decoder.Decode(&config); err != nil && err != io.EOF {
			log.Fatalf("Error: Failed to decode filters file: %v", err)
		}

	}

	return config
}

// CreateConfigFile creates a config file to a given path based on the settings
func CreateConfigFile(path string, name string) error {

	fullPath := path + name
	content, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}
	dir := filepath.Dir(fullPath)
	if err = os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(fullPath, content, 0644)
}
