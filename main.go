package main

import (
	"context"
	"fmt"
	"os"

	"github.com/numaproj/numaflow-go/pkg/sourcer"

	"github.com/numaproj-contrib/nats-source-go/pkg/config"
	"github.com/numaproj-contrib/nats-source-go/pkg/nats"
	"github.com/numaproj-contrib/nats-source-go/pkg/utils"
)

func main() {
	logger := utils.NewLogger()
	// Get the config file path and format from env vars
	var format string
	format, ok := os.LookupEnv("CONFIG_FORMAT")
	if !ok {
		logger.Info("CONFIG_FORMAT not set, defaulting to yaml")
		format = "yaml"
	}

	var config *config.NatsConfig
	var err error

	config, err = getConfigFromEnvVars(format)
	if err != nil {
		config, err = getConfigFromFile(format)
		if err != nil {
			logger.Panic("Failed to parse config file : ", err)
		} else {
			logger.Info("Successfully parsed config file")
		}
	} else {
		logger.Info("Successfully parsed config from env vars")
	}

	natsSrc, err := nats.New(config)
	if err != nil {
		logger.Panic("Failed to create nats source : ", err)
	}
	defer natsSrc.Close()

	err = sourcer.NewServer(natsSrc).Start(context.Background())
	if err != nil {
		logger.Panic("Failed to start source server : ", err)
	}
}

func getConfigFromFile(format string) (*config.NatsConfig, error) {
	if format == "yaml" {
		parser := &config.YAMLConfigParser{}
		content, err := os.ReadFile(fmt.Sprintf("%s/nats-config.yaml", utils.ConfigVolumePath))
		if err != nil {
			return nil, err
		}
		return parser.Parse(string(content))
	} else if format == "json" {
		parser := &config.JSONConfigParser{}
		content, err := os.ReadFile(fmt.Sprintf("%s/nats-config.json", utils.ConfigVolumePath))
		if err != nil {
			return nil, err
		}
		return parser.Parse(string(content))
	} else {
		return nil, fmt.Errorf("invalid config format %s", format)
	}
}

func getConfigFromEnvVars(format string) (*config.NatsConfig, error) {
	var c string
	c, ok := os.LookupEnv("NATS_CONFIG")
	if !ok {
		return nil, fmt.Errorf("NATS_CONFIG environment variable is not set")
	}
	if format == "yaml" {
		parser := &config.YAMLConfigParser{}
		return parser.Parse(c)
	} else if format == "json" {
		parser := &config.JSONConfigParser{}
		return parser.Parse(c)
	} else {
		return nil, fmt.Errorf("invalid config format %s", format)
	}
}
