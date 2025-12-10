// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
package exporter

import (
        "uos-mgmt-exporter/pkg/logger"
)
var (
        DefaultConfig = Config{
                Logging: logger.Config{
                        Level:   "debug",
                        LogPath: "/var/log/uos-exporter/mgmt-exporter.log",
                        MaxSize: "10MB",
                        MaxAge:  time.Hour * 24 * 7},
                Address:     "0.0.0.0",
                Port:        9098,
                MetricsPath: "/metrics",
        }
)

type Config struct {
        Logging     logger.Config `yaml:"log"`
        Address     string        `yaml:"address"`
        Port        int           `yaml:"port"`
        MetricsPath string        `yaml:"metricsPath"`
}

