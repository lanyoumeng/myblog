// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package middleware

import (
	"blog/internal/pkg/core"
	"blog/internal/pkg/errno"
	"blog/pkg/known"
	"blog/pkg/token"

	"github.com/gin-gonic/gin"
)

func Authn() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		username, err := token.ParseRequest(ctx)
		if err != nil {
			core.WriteResponse(ctx, errno.ErrTokenInvalid, nil)
			ctx.Abort()
		}
		ctx.Set(known.XUsernameKey, username)
		ctx.Next()
	}
}
