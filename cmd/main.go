package main

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/marcin-mc/find-index/internal/service"
	"github.com/spf13/viper"
)

const (
	defaultPort     = 5000
	defaultLogLevel = "INFO"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	if err := run(service.DataFilepath); err != nil {
		panic(err)
	}
	wg.Wait()
}

func run(dataFilepath string) error {
	viper.AddConfigPath(".")
	viper.SetDefault("port", defaultPort)
	viper.SetDefault("log_level", defaultLogLevel)
	if err := readConfig(); err != nil {
		slog.Info("There was problem using confing file. Default values will apply.")
	}
	logger := getLogger()
	if viper.Get("log_level") != "DEBUG" {
		gin.SetMode(gin.ReleaseMode)
	}
	srv, err := service.NewServer(logger, dataFilepath)
	if err != nil {
		return fmt.Errorf("cannot start server: %w", err)
	}
	go func() {
		if err := srv.Serve(); err != nil {
			panic(err)
		}
	}()
	return nil
}

func readConfig() error {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("cannot find config file: %w", err)
		} else {
			return fmt.Errorf("error reading config file: %w", err)
		}
	}
	return nil
}

func getLogger() *slog.Logger {
	logLevel := viper.GetString("log_level")
	var slogLevel slog.Level
	switch logLevel {
	case "DEBUG":
		slogLevel = slog.LevelDebug
	case "INFO":
		slogLevel = slog.LevelInfo
	case "ERROR":
		slogLevel = slog.LevelError
	}
	opts := &slog.HandlerOptions{
		Level: slogLevel,
	}
	h := slog.NewTextHandler(os.Stdout, opts)
	return slog.New(h)
}
