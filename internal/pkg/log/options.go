// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package log

import "go.uber.org/zap/zapcore"

type Options struct {
	Level             string
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	EncoderConfig     zapcore.EncoderConfig
	OutputPaths       []string
	ErrorOutputPaths  []string
	InitialFields     map[string]interface{}
	Format            string
}

func NewOptions() *Options {
	return &Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             zapcore.InfoLevel.String(),
		Format:            "console",
		OutputPaths:       []string{"stdout"},
	}
}
