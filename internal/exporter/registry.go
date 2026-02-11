// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
package exporter

import (
    "sync"
)

type Registry struct {
	metrics []Metric
	mu      sync.RWMutex
}

func init() {
	defaultReg = NewRegistry()
}

func NewRegistry() *Registry {
	return &Registry{
		metrics: []Metric{},
	}
}

func RegisterPrometheus(reg *prometheus.Registry) {
	reg.MustRegister(defaultReg)
}

