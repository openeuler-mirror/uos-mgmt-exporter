// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT

package server

import (
	"uos-mgmt-exporter/pkg/ratelimit"
	"github.com/alecthomas/kingpin"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	rateLimitInterval *time.Duration
	rateLimitSize     *int
	UseRatelimit      *bool
)

func Ratelimit(ratelimiter *ratelimit.RateLimiter) HandlerFunc {
	logrus.Info("ratelimit middleware init")
	logrus.Debugf("ratelimit middleware init rateLimitInterval: %v, rateLimitSize: %v\n", *rateLimitInterval, *rateLimitSize)
	return func(req *Request) {
		if err := ratelimiter.Get(); err != nil {
			req.Error = err
			req.Fail(429)
		}
	}
}
