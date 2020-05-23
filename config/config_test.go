package config

import (
	"os"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	setDummyEnvVar()
	expectedOutput := getDummyConfig()

	config := Read()

	if !reflect.DeepEqual(config, expectedOutput) {
		t.Fatal("unexpected output")
	}
}

func setDummyEnvVar() {
	dummyConfig := getDummyConfig()

	os.Setenv("LOG_FILE_PATH", dummyConfig.Log.FilePath)
}

func getDummyConfig() Config {
	dummyConfig := Config{
		Log: Log{
			FilePath: "dummy.log",
		},
	}

	return dummyConfig
}
