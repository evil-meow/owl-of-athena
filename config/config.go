package config

type Config struct {
	Version  string `yaml:"version"`
	Services []struct {
		image string `yaml:"image"`
		url   string `yaml:"url"`
	} `yaml:"services"`
}
