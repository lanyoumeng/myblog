// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package middleware

import (
	"blog/pkg/known"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		requestID := ctx.Request.Header.Get(known.XRequestIDKey)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		ctx.Set(known.XRequestIDKey, requestID)
		ctx.Writer.Header().Set(known.XRequestIDKey, requestID)
		ctx.Next()

	}
}
