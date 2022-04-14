package mynewscraper // import "github.com/codeboten/newcomponents/receiver/mynewscraper"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

const typeStr = "mynewscraper"

var errConfigNotMyNewConfig = errors.New("config was not a mynewscraper receiver config")

// NewFactory creates a new receiver factory
func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultConfig,
		component.WithMetricsReceiver(createMetricsReceiver),
	)
}

func createDefaultConfig() config.Receiver {
	return &Config{}
}

func createMetricsReceiver(ctx context.Context, params component.ReceiverCreateSettings, rConf config.Receiver, consumer consumer.Metrics) (component.MetricsReceiver, error) {
	cfg, ok := rConf.(*Config)
	if !ok {
		return nil, errConfigNotMyNewConfig
	}

	mynewScraper := newScraper(params.Logger, cfg, params.TelemetrySettings)
	scraper, err := scraperhelper.NewScraper(typeStr, mynewScraper.scrape, scraperhelper.WithStart(mynewScraper.start))
	if err != nil {
		return nil, err
	}

	return scraperhelper.NewScraperControllerReceiver(&cfg.ScraperControllerSettings, params, consumer, scraperhelper.AddScraper(scraper))
}
