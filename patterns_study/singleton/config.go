package config

import "time"

type Config struct {
	InitTime time.Time
}

var instance *Config

func initConfig() {
	instance = &Config{
		InitTime: time.Now(),
	}
}

func GetConfig() *Config {
	if instance == nil {
		initConfig()
	}
	return instance
}
