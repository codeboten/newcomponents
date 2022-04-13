package mynewexporter // import "github.com/codeboten/newcomponents/exporter/mynewexporter"

import (
	"go.opentelemetry.io/collector/config"
)

// Config defines configuration for file exporter.
type Config struct {
	config.ExporterSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct

	// Add additional fields here. See example below:
	// Path string `mapstructure:"path"`
}

var _ config.Exporter = (*Config)(nil)

// Validate checks if the exporter configuration is valid
func (cfg *Config) Validate() error {
	// do validation here
	return nil
}
