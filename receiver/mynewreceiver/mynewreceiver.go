package mynewreceiver // import "github.com/codeboten/newcomponents/receiver/mynewreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
)

type mynewReceiver struct {
}

func (r *mynewReceiver) Start(_ context.Context, host component.Host) error {
	return nil
}

func (r *mynewReceiver) Shutdown(_ context.Context) error {
	return nil
}
