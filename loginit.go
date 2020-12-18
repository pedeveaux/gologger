package logger

import (
	"encoding/json"
	"fmt"
	"os"
)

//Config defines the struct for the application configuration
type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"database"`
	RabbitMQ struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Exchange string `json:"exchange"`
	} `json:"rabbitmq"`
	LogConfig struct {
		EnableConsole     bool   `json:"enableConsole"`
		ConsoleLevel      string `json:"consoleLevel"`
		ConsoleJSONFormat bool   `json:"consoleJSONformat"`
		EnableFile        bool   `json:"enableFile"`
		FileLevel         string `json:"fileLevel"`
		FileJSONFormat    bool   `json:"fileJSONformat"`
		FileLocation      string `json:"fileLocation"`
		MaxSize           int    `json:"maxSize"` //MB
		Compress          bool   `json:"compress"`
		MaxAge            int    `json:"maxAge"` //Days
	} `json:"logconfig"`
}

//LoadConfiguration loads the configuration for the application
func LoadConfiguration(filename string) (config Config, err error) {
	configfile, err := os.Open(filename)
	defer configfile.Close()
	if err != nil {
		return config, fmt.Errorf("LoadConfiguration failure: %s", err.Error())
	}
	jsonParser := json.NewDecoder(configfile)
	err = jsonParser.Decode(&config)
	return config, err
}

//InitLogger creates a zap logger for use in the program
func InitLogger() Logger {

	config, err := LoadConfiguration("config.json")
	log.Fatalf("Failed to load configuration: %s", err.Error())
	logconfig := config.LogConfig

	err = NewLogger(logconfig, InstanceZapLogger)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}

	// contextLogger := WithFields(Fields{"key1": "value1"})

	return log
}
