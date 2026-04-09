// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
package metrics

import (
	"fmt"
	"errors"
	"sync"
	"strconv"
	"mgmt_exporter/internal/exporter"
	"github.com/prometheus/client_golang/prometheus"
)

type resStateWithKind struct {
	state ResState
	kind  string
}

type Prometheus struct {
	// 指标
	checkApplyTotal        *prometheus.CounterVec
	pgraphStartTimeSeconds prometheus.Gauge
	managedResources       *prometheus.GaugeVec
	failedResourcesTotal   *prometheus.CounterVec
	failedResources        *prometheus.GaugeVec

	// 状态管理
	resourcesState map[string]resStateWithKind
}

func init() {

	collector := NewMgmtCollect()
	collector.UpdateCheckApplyTotal("Pod", true, false, true)
	collector.UpdateCheckApplyTotal("Pod", true, false, true)
	collector.UpdateCheckApplyTotal("Pod", true, false, true)
	collector.UpdateCheckApplyTotal("Pod", true, false, true)
	collector.UpdateCheckApplyTotal("Deployment", false, true, false)

	// 更新 managedResources
	collector.UpdateManagedResources("Pod", 5)

	// 模拟资源状态更新
	collector.resourcesState["resource1"] = resStateWithKind{state: ResStateSoftFail, kind: "Pod"}
	collector.resourcesState["resource2"] = resStateWithKind{state: ResStateHardFail, kind: "Deployment"}
	exporter.Register(collector)
}

func NewMgmtCollect() *Prometheus {
	checkApplyTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "mgmt_checkapply_total",
			Help: "Number of CheckApply that have run.",
		},
		[]string{"kind", "apply", "eventful", "errorful"},
	)

	pgraphStartTimeSeconds := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "mgmt_graph_start_time_seconds",
			Help: "Start time of the current graph since unix epoch in seconds.",
		},
	)

	managedResources := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "mgmt_resources",
			Help: "Number of managed resources.",
		},
		[]string{"kind"},
	)

	failedResourcesTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "mgmt_failures_total",
			Help: "Total of failed resources.",
		},
		[]string{"kind", "failure"},
	)

	failedResources := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "mgmt_failures",
			Help: "Number of failing resources.",
		},
		[]string{"kind", "failure"},
	)

	return &Prometheus{
		checkApplyTotal:        checkApplyTotal,
		pgraphStartTimeSeconds: pgraphStartTimeSeconds,
		managedResources:       managedResources,
		failedResourcesTotal:   failedResourcesTotal,
		failedResources:        failedResources,
		resourcesState:         make(map[string]resStateWithKind),
	}
}
