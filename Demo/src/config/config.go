package config

import (
	"constants"
	"os"
	"structdemo"

	yaml "gopkg.in/yaml.v2"
)

// Config type
type Config structdemo.Config

// GetEnv function
func GetEnv() (cfg *Config) {
	f, err := os.Open(constants.CfgDic)
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
