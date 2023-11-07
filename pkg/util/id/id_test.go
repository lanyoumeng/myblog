// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package id

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGenShortID(t *testing.T) {
	shortID := GenShortID()
	assert.NotEqual(t, "", shortID)
	assert.Equal(t, 6, len(shortID))
}

func BenchmarkGenShortID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}

func BenchmarkGenShortIDTimeConsuming(b *testing.B) {
	b.StopTimer()

	shortId := GenShortID()

	if shortId == "" {
		b.Error("Failed to generate short id")
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}
