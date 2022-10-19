package state

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Path string `yaml:"-"`

	Telegram struct {
		BotToken        string  `yaml:"bot_token"`
		ApiURL          string  `yaml:"api_url"`
		OwnerID         int64   `yaml:"owner_id"`
		SudoUsers       []int64 `yaml:"sudo_users"`
		AuthorizedChats []int64 `yaml:"authorized_chats"`
	} `yaml:"telegram"`

	Database map[string]string `yaml:"database"`
}

func (cfg *Config) LoadConfig() {
	configPath := cfg.Path

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalln("config file not found : ", err)
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatalln("could not open config file for reading : ", err)
	}
	defer configFile.Close()

	configBody, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Fatalln("could not read config file : ", err)
	}

	err = yaml.Unmarshal(configBody, cfg)
	if err != nil {
		log.Fatalln("could not parse config file : ", err)
	}
}

func (cfg *Config) SaveConfig() error {
	configPath := cfg.Path

	configFile, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer configFile.Close()

	newConfigBody, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	_, err = configFile.Write(newConfigBody)
	if err != nil {
		return err
	}

	return nil
}
