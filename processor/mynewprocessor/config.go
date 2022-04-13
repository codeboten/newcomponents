package mynewprocessor // import "github.com/codeboten/newcomponents/processor/mynewprocessor"

import (
	"go.opentelemetry.io/collector/config"
)

// Config specifies the set options available.
type Config struct {
	config.ProcessorSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct
}

var _ config.Processor = (*Config)(nil)

// Validate checks if the processor configuration is valid
func (cfg *Config) Validate() error {
	return nil
}
