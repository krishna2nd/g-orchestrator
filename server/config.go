package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	PluginDir string   `json:"PluginDir"`
	Plugins   []string `json:"Plugins"`
}

func ReadConfig(cfgFile string) *Config {
	var config *Config = &Config{}
	configFile, err := os.Open(cfgFile)
	if err != nil {
		fmt.Println("Error in reading config", err)
	}
	defer configFile.Close()
	configValue, _ := ioutil.ReadAll(configFile)
	json.Unmarshal(configValue, config)
	fmt.Print(config)
	return config
}
