package config

import (
	"log"

	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

func LoadConfig() {
	yamlPath := "config.yaml"
	if err := k.Load(file.Provider(yamlPath), yaml.Parser()); err != nil {
		log.Fatalf("error loading config %v", err)
	}
	if err := k.Load(env.Provider("PORTFOLIO_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(strings.TrimPrefix(s, "PORTFOLIO_")), "_", ".", -1)
	}), nil); err != nil {
		log.Printf("error loading env config %v", err)
	}

}

func GetConfig() *Config {
	var c Config
	if err := k.Unmarshal("", &c); err != nil {
		log.Fatal("error unmarshalling error %w", err)
	}
	return &c
}

func GetString(key string) string {
	return k.String(key)
}

func GetInt(key string) int {
	return k.Int(key)
}

func GetBool(key string) bool {
	return k.Bool(key)
}
