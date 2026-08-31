package config

import (
	"github.com/pelletier/go-toml/v2"
)

type Rule struct {
	FileFormat []string
	Dest       string
}

type Config struct {
	Source      string
	Delay       int
	AutoMakeDir bool
	Rules       map[string]Rule
}

func TomlUnmarshal(data []byte) (Config, error) {
	var config Config
	err := toml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func DestinyTable(regras map[string]Rule) map[string]string {
	routes := make(map[string]string)

	for _, regra := range regras { //categoria = _
		for _, extensao := range regra.FileFormat {
			routes[extensao] = regra.Dest
		}
	}
	return routes
}
