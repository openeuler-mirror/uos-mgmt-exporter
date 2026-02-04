// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
package server

import (
        "os"
        "uos-mgmt-exporter/internal/exporter"
        "uos-mgmt-exporter/pkg/logger"
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

func (s *Server) SetUp() error {
        defer func() {
                if s.Error != nil {
                        logrus.Errorf("SetUp error: %v", s.Error)
                }
        }()
        err := s.parse()
        if err != nil {
                logrus.Errorf("Parsing command line arguments failed: %v", err)
                return err
        }

        err = s.loadConfig()
        if err != nil {
                logrus.Errorf("Loading config file failed: %v", err)
                return err
        }
        err = s.setupLog()
        if err != nil {
                logrus.Errorf("SetUp error: %v", err)
                return err
        }
        return nil
}

func (s *Server) loadConfig() error {
    content, err := os.ReadFile(*exporter.Configfile)
	if err != nil {
		logrus.Errorf("Failed to read config file: %v", err)
		logrus.Info("Use default config")
		return nil
	}
	err = yaml.Unmarshal(content, &s.CommonConfig)
	if err != nil {
		logrus.Errorf("Failed to parse config file: %v", err)
		logrus.Info("Use default config")
		return nil
	}
	logrus.Infof("Loaded config file from: %s", *exporter.Configfile)
	logrus.Info("CommonConfig file loaded")
	return nil
}

func (s *Server) setupLog() error {
	size, err := humanize.ParseBytes(s.CommonConfig.Logging.MaxSize)
	if err != nil {
		logrus.Errorf("Parsing log size failed: %v", err)
		return err
	}
	logConfig := logger.NewConfig(s.CommonConfig.Logging.Level, s.CommonConfig.Logging.LogPath, safeUint64ToInt64(size), s.CommonConfig.Logging.MaxAge)
	logger.Init(logConfig)
	return nil
}
