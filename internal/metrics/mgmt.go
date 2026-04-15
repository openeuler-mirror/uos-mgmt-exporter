// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
package metrics

import (
	"fmt"
	"errors"
	"sync"
	"strconv"
	"uos-mgmt-exporter/internal/exporter"
	"github.com/prometheus/client_golang/prometheus"
)

type ResState int

const (
	ResStateOK ResState = iota
	ResStateSoftFail
	ResStateHardFail
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

// updateFailingGauge 更新失败资源指标
func (p *Prometheus) updateFailingGauge(ch chan<- prometheus.Metric) {
	softFails := make(map[string]float64)
	hardFails := make(map[string]float64)

	for _, entry := range p.resourcesState {
		switch entry.state {
		case ResStateSoftFail:
			softFails[entry.kind]++
		case ResStateHardFail:
			hardFails[entry.kind]++
		}
	}

	for kind, count := range softFails {
		p.failedResources.WithLabelValues(kind, "soft").Set(count)
	}
	for kind, count := range hardFails {
		p.failedResources.WithLabelValues(kind, "hard").Set(count)
	}
}

// updateManagedResources 更新 managedResources 指标
func (p *Prometheus) updateManagedResources() {
	resourceCounts := make(map[string]float64)

	// 统计每种资源类型的数量
	for _, entry := range p.resourcesState {
		resourceCounts[entry.kind]++
	}

	// 更新 managedResources 指标
	for kind, count := range resourceCounts {
		p.managedResources.WithLabelValues(kind).Set(count)
	}
}

func (p *Prometheus) UpdateCheckApplyTotal(kind string, apply, eventful, errorful bool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	labels := prometheus.Labels{
		"kind":     kind,
		"apply":    strconv.FormatBool(apply),
		"eventful": strconv.FormatBool(eventful),
		"errorful": strconv.FormatBool(errorful),
	}
	p.checkApplyTotal.With(labels).Inc()
}

func (p *Prometheus) UpdateManagedResources(kind string, count int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.managedResources.WithLabelValues(kind).Set(float64(count))
}

func (p *Prometheus) UpdateState(resUUID string, rtype string, newState ResState) error {
	//defer p.updateFailingGauge(newState)
	if p == nil {
		return nil // happens when mgmt is launched without --prometheus
	}
	p.mutex.Lock()
	p.resourcesState[resUUID] = resStateWithKind{state: newState, kind: rtype}
	p.mutex.Unlock()

	if newState != ResStateOK {
		var strState string
		if newState == ResStateSoftFail {
			strState = "soft"
		} else if newState == ResStateHardFail {
			strState = "hard"
		} else {
			return errors.New("state should be soft or hard failure")
		}

		// 更新 failedResourcesTotal 指标
		p.failedResourcesTotal.WithLabelValues(rtype, strState).Inc()
	}
	return nil
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

	// 模拟资源状态更新
	if err := collector.UpdateState("resource1", "Pod", ResStateSoftFail); err != nil {
	    fmt.Printf("Warning: Failed to update state for resource1: %v", err)
	}
	if err := collector.UpdateState("resource2", "Deployment", ResStateHardFail); err != nil {
	    fmt.Printf("Warning: Failed to update state for resource2: %v", err)
	}
	if err := collector.UpdateState("resource3", "Pod", ResStateSoftFail); err != nil {
	    fmt.Printf("Warning: Failed to update state for resource3: %v", err)
	}

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
