// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
package config

import (
        "mgmt_exporter/pkg/utils"
        "github.com/alecthomas/kingpin"
        "github.com/sirupsen/logrus"
)

var (
        ScrapeUrl       *string
)

func init() {
        ScrapeUrl = kingpin.Flag("scrape_uri",
                "Scrape URI").
                Short('s').
                String()
        if *ScrapeUrl != "" {
                if err := utils.ValidateURI(*ScrapeUrl); err != nil {
                        logrus.Warnf("Invalid scrape uri: %s", err)
                }
        }
}

type Settings struct {
        ScrapeUri string `yaml:"scrape_uri"`
}

