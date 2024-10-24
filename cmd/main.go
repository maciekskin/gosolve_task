package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/maciekskin/gosolve_task/pkg/api"
	"github.com/maciekskin/gosolve_task/pkg/numbers"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

const (
	defaultConformationLevel = 10
	defaultPort              = "8888"
	defaultLogLevel          = "Info"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("failed to create logger: ", err.Error())
		os.Exit(1)
	}
	err = runApp(logger)
	if err != nil {
		logger.Error("error running app", zap.String("error_message", err.Error()))
		os.Exit(1)
	}
}

func runApp(logger *zap.Logger) error {
	cfg, err := LoadConfig(logger)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	logCfg := zap.NewProductionConfig()
	switch cfg.LogLevel {
	case "Debug":
		logCfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "Info":
		logCfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "Error":
		logCfg.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		logCfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
	logger, err = logCfg.Build()
	if err != nil {
		return fmt.Errorf("failed to create logger with custom config: %w", err)
	}

	portNumber, err := strconv.Atoi(cfg.Port)
	if err != nil {
		return fmt.Errorf("failed to parse server port: %w", err)
	}

	input, err := loadInput()
	if err != nil {
		return fmt.Errorf("failed to load input: %w", err)
	}
	repository := numbers.NewNumbersSliceRepository(input, defaultConformationLevel, logger)
	service := numbers.NewIndexService(repository, logger)

	apiServices := api.ApiSevices{
		IndexService: service,
	}
	logger.Info("starting HTTP server", zap.Int("port", portNumber))
	return api.StartHttpServer(apiServices, portNumber)
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

type Config struct {
	Port     string `yaml:"port" envconfig:"PORT"`
	LogLevel string `yaml:"logLevel" envconfig:"LOG_LEVEL"`
}

func LoadConfig(logger *zap.Logger) (*Config, error) {
	f, err := os.Open("config.yml")
	if err != nil {
		logger.Error("failed to open config.yml file, using defaults")
		return &Config{Port: defaultPort, LogLevel: defaultLogLevel}, nil
	}
	defer f.Close()

	cfg := &Config{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return nil, err
	}

	err = envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
