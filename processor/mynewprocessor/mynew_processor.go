package mynewprocessor // import "github.com/codeboten/newcomponents/processor/mynewprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
)

type mynewProcessor struct {
	logger *zap.Logger
}

// newProcessor returns a processor.
func newProcessor(logger *zap.Logger) *mynewProcessor {
	return &mynewProcessor{
		logger: logger,
	}
}

func (a *mynewProcessor) processTraces(ctx context.Context, td pdata.Traces) (pdata.Traces, error) {
	return td, nil
}

func (a *mynewProcessor) processLogs(ctx context.Context, ld pdata.Logs) (pdata.Logs, error) {
	return ld, nil
}

func (a *mynewProcessor) processMetrics(ctx context.Context, md pdata.Metrics) (pdata.Metrics, error) {
	return md, nil
}
