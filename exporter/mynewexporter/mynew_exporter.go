package mynewexporter // import "github.com/codeboten/newcomponents/exporter/mynewexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/model/pdata"
)

// mynewExporter is the implementation of an example exporter
type mynewExporter struct {
}

func (e *mynewExporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func (e *mynewExporter) ConsumeTraces(_ context.Context, td pdata.Traces) error {
	// do something with traces
	return nil
}

func (e *mynewExporter) ConsumeMetrics(_ context.Context, md pdata.Metrics) error {
	// do something with metrics
	return nil
}

func (e *mynewExporter) ConsumeLogs(_ context.Context, ld pdata.Logs) error {
	// do something with logs
	return nil
}

func (e *mynewExporter) Start(context.Context, component.Host) error {
	var err error
	return err
}

func (e *mynewExporter) Shutdown(context.Context) error {
	return nil
}
