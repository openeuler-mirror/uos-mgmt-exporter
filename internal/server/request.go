// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
package server

import (
	"net/http"
)

type Request struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	Error          error
	handlers       []HandlerFunc
}

type HandlerFunc func(ctx *Request)

func NewRequest(w http.ResponseWriter, r *http.Request) *Request {
	return &Request{
		Request:        r,
		ResponseWriter: w,
	}
}

func (r *Request) Fail(status int) {
	http.Error(r.ResponseWriter, http.StatusText(status), status)
}

