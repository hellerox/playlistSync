package config

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Spotify Config
type spotifyConfig struct {
	ClientID     string   `yaml:"clientID"`
	ClientSecret string   `yaml:"clientSecret"`
	Users        []string `yaml:"users"`
}

// Config stores the service configurations.
type Config struct {
	dir     string
	Port    string         `yaml:"port"`
	Spotify *spotifyConfig `yaml:"spotify"`
}

// New returns a new Config object.
func New(dir string) (*Config, error) {
	log.Println("loading config from", dir)

	conf := &Config{
		dir: dir,
	}

	err := conf.ReloadConfigs()
	if err != nil {
		return nil, err
	}

	return conf, err
}

// ReloadConfigs reload the configurations stored in the configuration file.
func (c *Config) ReloadConfigs() error {
	yamlConfig, err := ioutil.ReadFile(c.dir + "/playlist-sync.yaml")
	if err != nil {
		return fmt.Errorf("error reading the config file: %+v", err)
	}

	err = yaml.Unmarshal(yamlConfig, c)
	if err != nil {
		return fmt.Errorf("error unmarshaling the config: %+v", err)
	}

	return nil
}
