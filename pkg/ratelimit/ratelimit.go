// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
package ratelimit

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrRateLimited   = errors.New("rate limited")
	ErrRateLimitSize = errors.New("limit must be greater than zero")
	ErrRateLimitTime = errors.New("invalid limit")
)

type RateLimiter struct {
	tokens chan struct{}
	limit  time.Duration
	ticker *time.Ticker
	stopOnce sync.Once
}

func NewRateLimiter(limit time.Duration, chanSize int) (*RateLimiter, error) {
	if chanSize <= 0 {
		return nil, ErrRateLimitSize
	}
	if limit <= 0 {
		return nil, ErrRateLimitTime
	}
	rl := &RateLimiter{
		tokens: make(chan struct{}, chanSize),
		limit:  limit,
		ticker: time.NewTicker(limit),
	}

	for i := 0; i < chanSize; i++ {
		rl.tokens <- struct{}{}
	}

	go rl.startRefreshTokens()
	return rl, nil
}

func (rl *RateLimiter) startRefreshTokens() {
	for range rl.ticker.C {
		select {
		case rl.tokens <- struct{}{}:
		default:
		}
	}
}

func (rl *RateLimiter) Get() error {
	select {
	case <-rl.tokens:
		return nil
	default:
		return ErrRateLimited
	}
}

func (rl *RateLimiter) Stop() {
	rl.ticker.Stop()
	rl.stopOnce.Do(func() {
		close(rl.tokens)
	})
}
