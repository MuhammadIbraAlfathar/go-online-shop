package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	App AppConfig `yaml:"app"`
	DB  DBConfig  `yaml:"db"`
}

type AppConfig struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type DBConfig struct {
	Host           string                 `yaml:"host"`
	Port           string                 `yaml:"port"`
	User           string                 `yaml:"user"`
	Password       string                 `yaml:"password"`
	Name           string                 `yaml:"name"`
	ConnectionPool DBConnectionPoolConfig `yaml:"connection_pool"`
}

type DBConnectionPoolConfig struct {
	MaxIdleConnection     uint8 `yaml:"max_idle_connection"`
	MaxOpenConnection     uint8 `yaml:"max_open_connection"`
	MaxLifetimeConnection uint8 `yaml:"max_lifetime_connection"`
	MaxIdleTimeConnection uint8 `yaml:"max_idleTime_connection"`
}

var Cfg Config

func LoadConfig(filename string) (err error) {

	readFile, err := os.ReadFile(filename)

	if err != nil {
		return
	}

	err = yaml.Unmarshal(readFile, &Cfg)

	if err != nil {
		return
	}

	return
}
