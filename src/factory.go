package src

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"log"
)

func NewFactory() exporter.Factory {
	newType, err := component.NewType("otelexporter")
	if err != nil {
		log.Fatal(err)
	}

	return exporter.NewFactory(
		newType,
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, component.StabilityLevelDevelopment),
		exporter.WithMetrics(createMetricsExporter, component.StabilityLevelDevelopment),
		exporter.WithLogs(createLogsExporter, component.StabilityLevelDevelopment),
	)
}

func createTracesExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config) (exporter.Traces, error) {

	cfg := config.(*Config)
	s := NewDigibeeExporter()
	return exporterhelper.NewTraces(ctx, set, cfg, s.pushTraces)
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config) (exporter.Metrics, error) {

	cfg := config.(*Config)
	s := NewDigibeeExporter()
	return exporterhelper.NewMetrics(ctx, set, cfg, s.pushMetrics)
}

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config) (exporter.Logs, error) {

	cfg := config.(*Config)
	s := NewDigibeeExporter()
	return exporterhelper.NewLogs(ctx, set, cfg, s.pushLogs)
}

type Config struct {
}

func createDefaultConfig() component.Config {
	return &Config{}
}
