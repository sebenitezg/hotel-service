package config

import (
	"log"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Configurations Application wide configurations
type Configurations struct {
	Server   ServerConfigurations   `koanf:"server"`
	Database DatabaseConfigurations `koanf:"database"`
}

type ServerConfigurations struct {
	Port      string `koanf:"port"`
	DebugMode bool   `koanf:"debug-mode"`
}

type DatabaseConfigurations struct {
	Host       string `koanf:"host"`
	Port       int    `koanf:"port"`
	DbName     string `koanf:"db-name"`
	Username   string `koanf:"user"`
	Password   string `koanf:"password"`
	PoolSize   int    `koanf:"pool-size"`
	LogQueries bool   `koanf:"log-queries"`
}

// LoadConfig Loads configurations depending upon the environment
func LoadConfig() (*Configurations, error) {
	k := koanf.New(".")
	err := k.Load(file.Provider("resources/config.yaml"), yaml.Parser())
	if err != nil {
		return nil, err
	}

	// Searches for env variables and will transform them into koanf format
	// e.g. SERVER_PORT variable will be server.port: value
	err = k.Load(env.Provider("", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(s), "_", ".")
	}), nil)
	if err != nil {
		return nil, err
	}

	var configuration Configurations

	err = k.Unmarshal("", &configuration)
	if err != nil {
		log.Fatalf("Failed to load configurations. %v", err)
	}

	return &configuration, nil
}
