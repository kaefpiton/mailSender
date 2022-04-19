package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	From struct {
		Email 		string `json:"email"`
		Password 	string `json:"password"`
	}

	SMTP struct{
		Host string `json:"host"`
		Port string	`json:"port"`
	}
}

func LoadConfig(file string)Config {
	var config Config

	configFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}
