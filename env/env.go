package env

import (
	"encoding/json"
	"os"
)

type Config struct {
	CredentialsFile string   `json:"credentialsFile"`
	TokenFile       string   `json:"tokenFile"`
	Scopes          []string `json:"scopes"`
	Actions         []string `json:"actions"`
}

func NewLoadConfig() (*Config, error) {
	return LoadConfig("config.json")
}

func LoadConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	config := &Config{}
	if err := json.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
