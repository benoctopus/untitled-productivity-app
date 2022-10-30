package config

import "util"

type Config struct {
	Port           string
	TaskServiceUrl string
}

func New() *Config {
	return &Config{
		Port:           util.MustGetEnv("PORT"),
		TaskServiceUrl: util.MustGetEnv("TASK_SERVICE_URL"),
	}
}
