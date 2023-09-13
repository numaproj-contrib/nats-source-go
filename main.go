package main

import (
	"context"
	"fmt"
	"os"

	"github.com/numaproj/numaflow-go/pkg/sourcer"
	"github.com/numaproj/numaflow/pkg/apis/numaflow/v1alpha1"

	"nats-source-go/pkg/nats"
	"nats-source-go/pkg/utils"
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

	config, err := getConfigFromFile(format)
	if err != nil {
		logger.Panic("Failed to parse config file : ", err)
	} else {
		logger.Info("Successfully parsed config file")
	}

	natsSrc, err := nats.New(config, nats.WithLogger(logger))
	if err != nil {
		logger.Panic("Failed to create nats source : ", err)
	}
	defer natsSrc.Close()

	err = sourcer.NewServer(natsSrc).Start(context.Background())
	if err != nil {
		logger.Panic("Failed to start source server : ", err)
	}
}

func getConfigFromFile(format string) (*v1alpha1.NatsSource, error) {
	if format == "yaml" {
		parser := &utils.YAMLConfigParser{}
		content, err := os.ReadFile(fmt.Sprintf("%s/nats-config.yaml", utils.ConfigVolumePath))
		if err != nil {
			return nil, err
		}
		return parser.Parse(string(content))
	} else if format == "json" {
		parser := &utils.JSONConfigParser{}
		content, err := os.ReadFile(fmt.Sprintf("%s/nats-config.json", utils.ConfigVolumePath))
		if err != nil {
			return nil, err
		}
		return parser.Parse(string(content))
	} else {
		return nil, fmt.Errorf("invalid config format %s", format)
	}
}
