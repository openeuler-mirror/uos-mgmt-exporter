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

	// 模拟资源操作
	collector.UpdateCheckApplyTotal("Pod", true, false, true)
	collector.UpdateCheckApplyTotal("Pod", true, false, true)
	collector.UpdateCheckApplyTotal("Pod", true, false, true)
	collector.UpdateCheckApplyTotal("Pod", true, false, true)
	collector.UpdateCheckApplyTotal("Deployment", false, true, false)

	// 更新 managedResources
	collector.UpdateManagedResources("Pod", 5)

	// 模拟资源状态更新
	collector.UpdateState("resource1", "Pod", ResStateSoftFail)
	collector.UpdateState("resource2", "Deployment", ResStateHardFail)
	collector.UpdateState("resource3", "Pod", ResStateSoftFail)

