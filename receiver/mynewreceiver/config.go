package mynewreceiver // import "github.com/codeboten/newcomponents/receiver/mynewreceiver"

import (
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/codeboten/newcomponents/receiver/mynewreceiver/internal/metadata"
)

// Config defines the configuration for the various elements of the receiver agent.
type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	confighttp.HTTPClientSettings           `mapstructure:",squash"`
	Metrics                                 metadata.MetricsSettings `mapstructure:"metrics"`
}

// Validate validates the configuration by checking for missing or invalid fields
func (cfg *Config) Validate() error {
	var err error

	// do some validation of the config here

	return err
}
