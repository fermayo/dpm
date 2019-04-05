package parser

// Command contains YAML definition for a command
type Command struct {
	Name        string   `yaml:"-"`
	Image       string   `yaml:"image"`
	Entrypoints []string `yaml:"entrypoints,omitempty"`
	Context     string   `yaml:"context,omitempty"`
	Volumes     []string `yaml:"volumes,omitempty"`
}

// dpmFile is the type for the entire YAML file
type dpmFile map[string]map[string]Command
