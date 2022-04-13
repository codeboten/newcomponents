package mynewreceiver // import "github.com/codeboten/newcomponents/receiver/mynewreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component" // zeromqScraper handles scraping of RabbitMQ metrics
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"

	"github.com/codeboten/newcomponents/receiver/mynewreceiver/internal/metadata"
)

type mynewScraper struct {
	logger   *zap.Logger
	cfg      *Config
	settings component.TelemetrySettings
	mb       *metadata.MetricsBuilder
}

// start starts the scraper by creating a new HTTP Client on the scraper
func (r *mynewScraper) start(ctx context.Context, host component.Host) (err error) {
	return nil
}

// scrape collects metrics from the RabbitMQ API
func (r *mynewScraper) scrape(ctx context.Context) (pdata.Metrics, error) {
	return r.mb.Emit(), nil
}

// newScraper creates a new scraper
func newScraper(logger *zap.Logger, cfg *Config, settings component.TelemetrySettings) *mynewScraper {
	return &mynewScraper{
		logger:   logger,
		cfg:      cfg,
		settings: settings,
		mb:       metadata.NewMetricsBuilder(cfg.Metrics),
	}
}
