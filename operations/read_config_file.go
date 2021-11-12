package operations

import (
	"errors"
	"evil-meow/owl-of-athena/github_api"
	"evil-meow/owl-of-athena/service_config"
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

func ReadConfigFile(serviceName *string) (*service_config.ServiceConfig, error) {
	configFileUrl := fmt.Sprintf("https://raw.githubusercontent.com/evil-meow/%s/main/owl.yml", *serviceName)
	configFile, err := github_api.ReadFile(&configFileUrl)
	if err != nil {
		log.Printf("Could not find owl.yml config file at: %s", configFileUrl)
		return nil, errors.New("owl.yml file not found")
	}

	conf := service_config.ServiceConfig{}

	yaml.Unmarshal([]byte(configFile), &conf)
	conf.RepoName = *serviceName + "-infra"

	return &conf, nil
}
