package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	App      AppConfig
}

type DatabaseConfig struct {
	URL string
}

type ServerConfig struct {
	Port    int
	Host    string
	GinMode string
}

type AppConfig struct {
	Name string
	Env  string
}

var cfg *Config

// Load reads configuration from .env file using Viper
func Load() (*Config, error) {
	v := viper.New()

	// Set the file name and path
	v.SetConfigFile(".env")
	v.SetConfigType("env")

	// AutomaticEnv will check for an environment variable any time
	// a viper.Get request is made
	v.AutomaticEnv()

	// Replace . with _ in environment variable names
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read the config file
	if err := v.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
	}

	cfg = &Config{
		Database: DatabaseConfig{
			URL: v.GetString("DATABASE_URL"),
		},
		Server: ServerConfig{
			Port:    v.GetInt("SERVER_PORT"),
			Host:    v.GetString("SERVER_HOST"),
			GinMode: v.GetString("GIN_MODE"),
		},
		App: AppConfig{
			Name: v.GetString("APP_NAME"),
			Env:  v.GetString("APP_ENV"),
		},
	}

	// Validate required configurations
	if cfg.Database.URL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	if cfg.Server.Port == 0 {
		return nil, fmt.Errorf("SERVER_PORT is required")
	}

	if cfg.Server.Host == "" {
		return nil, fmt.Errorf("SERVER_HOST is required")
	}

	return cfg, nil
}

// // Get returns the loaded configuration
// func Get() *Config {
// 	if cfg == nil {
// 		panic("Configuration not loaded. Call Load() first")
// 	}
// 	return cfg
// }
