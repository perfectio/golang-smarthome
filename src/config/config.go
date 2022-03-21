package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

type (
	Rule struct {
		When string `yaml:"when"`
	}

	Device struct {
		Name      string `yaml:"name"`
		Endpoints struct {
			Event  string `yaml:"event"`
			Action string `yaml:"action"`
		} `yaml:"endponts"`
	}

	Config struct {
		Devices []Device `yaml:"devices"`
		Rules   []Rule   `yaml:"rules"`
	}
)

func Load() *Config {
	data, err := os.ReadFile("rules.yml")
	if err != nil {
		panic(err)
	}

	cfg := Config{}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}

	log.Debugf("%+v", cfg)

	return &cfg
}
