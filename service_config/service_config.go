package service_config

type ServiceConfig struct {
	RepoName     string
	Version      string `yaml:"version"`
	Name         string `yaml:"name"`
	Replicas     int8   `yaml:"replicas"`
	Image        string `yaml:"image"`
	InternalPort int16  `yaml:"internal_port"`
	Url          string `yaml:"url"`
}
