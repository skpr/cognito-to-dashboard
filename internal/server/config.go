package server

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Region           string        `mapstructure:"SKPR_COGNITO_TO_DASHBOARD_REGION"`
	ClientID         string        `mapstructure:"SKPR_COGNITO_TO_DASHBOARD_CLIENT_ID"`
	Host             string        `mapstructure:"SKPR_COGNITO_TO_DASHBOARD_HOST"`
	Scope            []string      `mapstructure:"SKPR_COGNITO_TO_DASHBOARD_SCOPE"`
	Callback         string        `mapstructure:"SKPR_COGNITO_TO_DASHBOARD_CALLBACK"`
	IdentityPool     string        `mapstructure:"SKPR_COGNITO_TO_DASHBOARD_IDENTITY_POOL"`
	IdentityProvider string        `mapstructure:"SKPR_COGNITO_TO_DASHBOARD_IDENTITY_PROVIDER"`
	SessionDuration  int           `mapstructure:"SKPR_COGNITO_TO_DASHBOARD_SESSION_DURATION"`
	StorageRetention time.Duration `mapstructure:"SKPR_COGNITO_TO_DASHBOARD_STORAGE_RETENTION"`
	AllowedListPath  string        `mapstructure:"SKPR_COGNITO_TO_DASHBOARD_ALLOWED_LIST_PATH"`
}

// Validate validates the config.
func (c Config) Validate() error {
	var errs []error

	if c.Region == "" {
		errs = append(errs, fmt.Errorf("SKPR_COGNITO_TO_DASHBOARD_REGION is a required variable"))
	}

	if c.ClientID == "" {
		errs = append(errs, fmt.Errorf("SKPR_COGNITO_TO_DASHBOARD_CLIENT_ID is a required variable"))
	}

	if len(c.Scope) == 0 {
		errs = append(errs, fmt.Errorf("SKPR_COGNITO_TO_DASHBOARD_SCOPE is a required variable"))
	}

	if c.Callback == "" {
		errs = append(errs, fmt.Errorf("SKPR_COGNITO_TO_DASHBOARD_CALLBACK is a required variable"))
	}

	if c.IdentityPool == "" {
		errs = append(errs, fmt.Errorf("SKPR_COGNITO_TO_DASHBOARD_IDENTITY_POOL is a required variable"))
	}

	if c.IdentityProvider == "" {
		errs = append(errs, fmt.Errorf("SKPR_COGNITO_TO_DASHBOARD_IDENTITY_PROVIDER is a required variable"))
	}

	if c.StorageRetention == 0 {
		errs = append(errs, fmt.Errorf("SKPR_COGNITO_TO_DASHBOARD_STORAGE_RETENTION is a required variable"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("defaults")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	var config Config

	err := viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to read config: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return config, err
}
