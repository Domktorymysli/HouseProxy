package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
	Weather struct {
		ApiKey string `json:"apiKey"`
		City   string `json:"city"`
	} `json:"weather"`
	Recuperator struct {
		Ip       string `json:"ip"`
		Login    string `json:"login"`
		Password string `json:"password"`
	} `json:"recuperator"`
	Clus []struct {
		Name string `json:"name"`
		Ip   string `json:"ip"`
		Port string `json:"port"`
		Key  string `json:"key"`
		Iv   string `json:"iv"`
	} `json:"clus"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		panic("missing config file: " + file + "!")
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
