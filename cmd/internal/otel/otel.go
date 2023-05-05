package otel

import (
	"github.com/tsingsun/woocoo/contrib/telemetry"
	"github.com/tsingsun/woocoo/pkg/conf"
	"go.opentelemetry.io/contrib/propagators/b3"
)

// Apply 尝试注册otel,如果配置中有otel配置,则注册.并返回关闭函数
func Apply(cnf *conf.AppConfiguration) func() {
	if cnf.IsSet("otel") {
		otelCnf := cnf.Sub("otel")
		otelcfg := telemetry.NewConfig(otelCnf,
			telemetry.WithPropagator(b3.New()),
		)
		return func() {
			otelcfg.Shutdown()
		}
	}
	return func() {}
}
