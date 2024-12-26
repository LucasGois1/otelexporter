package src

import (
	"context"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type DigibeeExporter struct {
}

func NewDigibeeExporter() *DigibeeExporter {
	return &DigibeeExporter{}
}

func (s *DigibeeExporter) pushLogs(_ context.Context, ld plog.Logs) error {
	return nil
}

func (s *DigibeeExporter) pushMetrics(ctx context.Context, md pmetric.Metrics) error {
	return nil
}

func (s *DigibeeExporter) pushTraces(_ context.Context, td ptrace.Traces) error {
	return nil
}
