package mynewreceiver // import "github.com/codeboten/newcomponents/receiver/mynewreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
)

const typeStr = "mynew"

var errConfigNotMyNewConfig = errors.New("config was not a mynew receiver config")

// NewFactory creates a new receiver factory
func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultConfig,
		component.WithTracesReceiver(createTracesReceiver),
		component.WithMetricsReceiver(createMetricsReceiver),
		component.WithLogsReceiver(createLogsReceiver),
	)
}

func createDefaultConfig() config.Receiver {
	return &Config{}
}

func createTracesReceiver(ctx context.Context, params component.ReceiverCreateSettings, rConf config.Receiver, consumer consumer.Traces) (component.TracesReceiver, error) {
	return &mynewReceiver{}, nil
}

func createMetricsReceiver(ctx context.Context, params component.ReceiverCreateSettings, rConf config.Receiver, consumer consumer.Metrics) (component.MetricsReceiver, error) {
	return &mynewReceiver{}, nil
}

func createLogsReceiver(ctx context.Context, params component.ReceiverCreateSettings, rConf config.Receiver, consumer consumer.Logs) (component.LogsReceiver, error) {
	return &mynewReceiver{}, nil
}
