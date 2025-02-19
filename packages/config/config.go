package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) ConfigGetDBURL() string {
	return c.DBUrl
}

func exportConfigJSON(c *Config) []byte {
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error encoding JSON config", err)
	}
	return jsonBytes
}

func ReadConfig() (Config, error) {
	file, err := GetConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	jsonData, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening config file", err)
		return Config{}, err
	}
	defer jsonData.Close()

	var newConfig Config
	decoder := json.NewDecoder(jsonData)
	err = decoder.Decode(&newConfig)
	if err != nil {
		fmt.Println("Error decoding config file", err)
		return Config{}, err
	}
	return newConfig, nil
}

func GetConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".gatorconfig.json"), nil
}

func (c *Config) SetUser(userName string) error {
	fmt.Println("Setting user to", userName)
	c.CurrentUserName = userName
	return write(*c)
}

func write(c Config) error {
	fullPath, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) GetCurrentUser() string {
	return c.CurrentUserName
}
