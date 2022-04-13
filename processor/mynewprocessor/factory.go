package mynewprocessor // import "github.com/codeboten/newcomponents/processor/mynewprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

const (
	// typeStr is the value of "type" key in configuration.
	typeStr = "mynew"
)

var processorCapabilities = consumer.Capabilities{MutatesData: true}

// NewFactory returns a new factory for the Attributes processor.
func NewFactory() component.ProcessorFactory {
	return component.NewProcessorFactory(
		typeStr,
		createDefaultConfig,
		component.WithTracesProcessor(createTracesProcessor),
		component.WithLogsProcessor(createLogProcessor),
		component.WithMetricsProcessor(createMetricsProcessor))
}

// Note: This isn't a valid configuration because the processor would do no work.
func createDefaultConfig() config.Processor {
	return &Config{
		ProcessorSettings: config.NewProcessorSettings(config.NewComponentID(typeStr)),
	}
}

func createTracesProcessor(
	_ context.Context,
	params component.ProcessorCreateSettings,
	cfg config.Processor,
	nextConsumer consumer.Traces,
) (component.TracesProcessor, error) {
	return processorhelper.NewTracesProcessor(
		cfg,
		nextConsumer,
		newProcessor(params.Logger).processTraces,
		processorhelper.WithCapabilities(processorCapabilities))
}

func createLogProcessor(
	_ context.Context,
	params component.ProcessorCreateSettings,
	cfg config.Processor,
	nextConsumer consumer.Logs,
) (component.LogsProcessor, error) {
	return processorhelper.NewLogsProcessor(
		cfg,
		nextConsumer,
		newProcessor(params.Logger).processLogs,
		processorhelper.WithCapabilities(processorCapabilities))
}

func createMetricsProcessor(
	_ context.Context,
	params component.ProcessorCreateSettings,
	cfg config.Processor,
	nextConsumer consumer.Metrics,
) (component.MetricsProcessor, error) {
	return processorhelper.NewMetricsProcessor(
		cfg,
		nextConsumer,
		newProcessor(params.Logger).processMetrics,
		processorhelper.WithCapabilities(processorCapabilities))
}
