package models

import "regexp"

type FilterConfig struct {
	Filters []Filter `yaml:"Filters"`
}

// Filter struct
type Filter struct {
	FilterName   string         `yaml:"filter-name"`           // FilterName is the name of the filter
	Watchers     []string       `yaml:"watchers"`              // Watchers is the list of watchers to be filtered
	Target       []Target       `yaml:"target"`                // Target is the key-value pair to be matched
	PlainReplace []PlainReplace `yaml:"plain-replace"`         // Replace is the key-value pair to be replaced
	RegexReplace []RegexReplace `yaml:"regex-replace"`         // Replace is the key-value pair to be replaced
	Enable       bool           `yaml:"enable" default:"true"` // Enabled is the value to enable or disable the filter
	Drop         bool           `yaml:"drop"`                  // Drop is the flag to drop the event if the filter matches
	Category     string         `yaml:"category"`              // Category is the category of the metric
}

type Target struct {
	Key   string         `yaml:"key"`
	Value *regexp.Regexp `yaml:"value"`
}

type PlainReplace struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type RegexReplace struct {
	Key        string         `yaml:"key"`
	Expression *regexp.Regexp `yaml:"from"`
	Value      string         `yaml:"value"`
}
