// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
package main

import (
    "github.com/sirupsen/logrus"
    "uos-mgmt-exporter/pkg/logger"
    "uos-mgmt-exporter/internal/server"
)

func Run(name string, version string) error {
    logger.InitDefaultLog()
    s := server.NewServer(name, version)

    s.PrintVersion()
    err := s.SetUp()
	if err != nil {
		logrus.Errorf("SetUp error: %v", err)
		return err
	}

    go func() {
        err := s.Run()
        if err != nil {
            logrus.Errorf("Run error: %v", err)
            s.Error = err
        }

        s.Exit()
    }()
    return nil
}

