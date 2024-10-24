package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/maciekskin/gosolve_task/pkg/api"
	"github.com/maciekskin/gosolve_task/pkg/numbers"

	"go.uber.org/zap"
)

const defaultConformationLevel = 10

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("failed to initialize logger: ", err.Error())
		os.Exit(1)
	}
	err = runApp(logger)
	if err != nil {
		logger.Error("error running app", zap.String("error_message", err.Error()))
		os.Exit(1)
	}
}

func runApp(logger *zap.Logger) error {
	input, err := loadInput()
	if err != nil {
		return fmt.Errorf("failed to load input: %w", err)
	}

	repository := numbers.NewNumbersSliceRepository(input, defaultConformationLevel, logger)
	service := numbers.NewIndexService(repository, logger)

	apiServices := api.ApiSevices{
		IndexService: service,
	}
	logger.Info("starting HTTP server")
	return api.StartHttpServer(apiServices)
}

func loadInput() ([]int, error) {
	inputFile, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}

	inputRaw := strings.Split(string(inputFile), "\n")
	input := make([]int, len(inputRaw))
	for idx, line := range inputRaw {
		if line != "" {
			input[idx], err = strconv.Atoi(line)
			if err != nil {
				return nil, err
			}
		}
	}
	return input, nil
}

// TODO:
// - add configuration file with service port and log level (Debug, Info, Error)
// - add README.md with service description
// - upload to GitHub
