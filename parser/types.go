package parser

type Command struct {
	Name       string   `yaml:"-"`
	Image      string   `yaml:"image"`
	Entrypoint string   `yaml:"entrypoint,omitempty"`
	Context    string   `yaml:"context,omitempty"`
	Volumes    []string `yaml:"volumes,omitempty"`
}

// type dpmFile map[string]map[string]Command
