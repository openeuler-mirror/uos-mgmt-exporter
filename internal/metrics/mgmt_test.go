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

	// 启动 HTTP 服务以暴露指标
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9233", nil)
	}()

	// 等待 HTTP 服务启动
	time.Sleep(2 * time.Second)

	// 验证指标是否正确更新
	validateMetrics(t, collector)
}

// validateMetrics 验证指标是否正确更新
func validateMetrics(t *testing.T, collector *Prometheus) {
	// 收集所有注册的指标
	metricFamilies, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		t.Fatalf("Failed to gather metrics: %v", err)
	}

	// 定义预期的指标值
	expectedMetrics := map[string]float64{
		`mgmt_checkapply_total{kind="Pod",apply="true",eventful="false",errorful="true"}`: 4,
		`mgmt_checkapply_total{kind="Deployment",apply="false",eventful="true",errorful="false"}`: 1,
		`mgmt_resources{kind="Pod"}`: 5,
		`mgmt_failures_total{kind="Pod",failure="soft"}`: 2,
		`mgmt_failures_total{kind="Deployment",failure="hard"}`: 1,
	}

	// 遍历指标，验证预期值
	for _, mf := range metricFamilies {
		for _, metric := range mf.Metric {
			// 获取指标名称和标签
			labels := make(map[string]string)
			for _, label := range metric.Label {
				labels[*label.Name] = *label.Value
			}

