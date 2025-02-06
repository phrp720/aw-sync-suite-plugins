package models

type ScriptConfig struct {
	Scripts []Script `yaml:"Scripts"`
}

type Script struct {
	Name    string `yaml:"name"`    // ScriptName is the name of the script
	Path    string `yaml:"path"`    // Path is the path of the script
	Timeout int    `yaml:"timeout"` // Timeout is the timeout of the script in seconds
}
