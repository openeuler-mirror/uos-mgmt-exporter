// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
package server

import (
        "uos-mgmt-exporter/internal/exporter"
)

var defaultSeverVersion = "1.0.0"

type Server struct {
        Name           string
        Version        string
        CommonConfig   exporter.Config
}

func NewServer(name, version string) *Server {
        if version == "" {
                version = defaultSeverVersion
        }
        s := &Server{
                Name:         name,
                Version:      version,
                CommonConfig: exporter.DefaultConfig,
        }
        return s
}

