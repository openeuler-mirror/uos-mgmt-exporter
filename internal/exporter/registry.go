// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

var defaultReg *Registry

func init() {
	defaultReg = NewRegistry()
}

func Register(metric Metric) {
	defaultReg.Register(metric)
}

type Registry struct {
	metrics []Metric
	mu      sync.RWMutex
}

func NewRegistry() *Registry {
	return &Registry{
		metrics: []Metric{},
	}
}

func RegisterPrometheus(reg *prometheus.Registry) {
	reg.MustRegister(defaultReg)
}

func (r *Registry) Register(metrics Metric) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.metrics = append(r.metrics, metrics)
}

func (r *Registry) GetMetrics() []Metric {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.metrics
}

func (r *Registry) Describe(descs chan<- *prometheus.Desc) {
}

func (r *Registry) Collect(ch chan<- prometheus.Metric) {
	for _, m := range r.GetMetrics() {
		m.Collect(ch)
	}
}
