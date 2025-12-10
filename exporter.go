// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
package main

import (
        "uos-mgmt-exporter/pkg/logger"
        "uos-mgmt-exporter/internal/server"
)

func Run(name string, version string) error {
        logger.InitDefaultLog()
        server.NewServer(name, version)
        return nil
}

