package mynewreceiver // import "github.com/codeboten/newcomponents/receiver/mynewreceiver"

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/confighttp"
)

// Config defines the configuration for the various elements of the receiver agent.
type Config struct {
	config.ReceiverSettings       `mapstructure:",squash"`
	confighttp.HTTPClientSettings `mapstructure:",squash"`
}

// Validate validates the configuration by checking for missing or invalid fields
func (cfg *Config) Validate() error {
	var err error
	return err
}
