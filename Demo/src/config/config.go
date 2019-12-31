package config

import (
	"os"
	"structdemo"

	yaml "gopkg.in/yaml.v2"
)

// Config type
type Config structdemo.Config

// GetEnv function
func GetEnv() (cfg *Config) {
	f, err := os.Open("./src/config/config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	return
}
