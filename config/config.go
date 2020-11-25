package config

import (
	"os"
)

type (
	Config struct {
		Log Log `json:"log"`
	}

	Log struct {
		FilePath string `json:"file_path"`
	}
)

var Instance Config

func Read() (config Config) {
	config = readFromEnvVar()

	return
}

func readFromEnvVar() (config Config) {
	config.Log.FilePath = readEnvVarWithDefaultValue("LOG_FILE_PATH", "http.log")

	return
}

func readEnvVarWithDefaultValue(key, defaultValue string) string {
	if envVarValue, ok := os.LookupEnv(key); ok {
		return envVarValue
	}
	return defaultValue
}
