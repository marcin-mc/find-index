package main

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	viper.AddConfigPath("./..")
	assert.NoError(t, Run("./../internal/server/test_data/input.txt"))
}

func TestRunFailsNoConfig(t *testing.T) {
	viper.Reset()
	assert.NoError(t, Run("./../internal/server/test_data/input.txt"))
	assert.Equal(t, defaultLogLevel, viper.GetString("log_level"))
	assert.Equal(t, defaultPort, viper.GetInt("port"))
}

func TestRunFailsBadDataFilepath(t *testing.T) {
	viper.AddConfigPath("./..")
	assert.ErrorContains(t, Run("./../bad/path/file.txt"), "cannot load data")
}
