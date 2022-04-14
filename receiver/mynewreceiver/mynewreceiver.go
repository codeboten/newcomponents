package mynewreceiver // import "github.com/codeboten/newcomponents/receiver/mynewreceiver"

import "go.opentelemetry.io/collector/component"

type mynewReceiver struct {
	component.StartFunc
	component.ShutdownFunc
}
