package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is the struct of the env config
type Config struct {
	DBURL string `mapstructure:"DB_URL"`
}

// LoadConfig loads the env config
func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	if config.DBURL == "" {
		return nil, fmt.Errorf("missing DB_URL in configuration")
	}

	return &config, nil
}
