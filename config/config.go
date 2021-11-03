package config

type Config struct {
	RepoName string
	Version  string `yaml:"version"`
	Name     string `yaml:"name"`
	Replicas int8   `yaml:"replicas"`
	Image    string `yaml:"image"`
	Url      string `yaml:"url"`
}
