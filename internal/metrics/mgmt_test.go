package metrics

import (
	"net/http"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// TestPrometheusMetrics 测试 Prometheus 指标是否正确更新
func TestPrometheusMetrics(t *testing.T) {
	// 初始化 Prometheus 收集器
	collector := NewMgmtCollect()

