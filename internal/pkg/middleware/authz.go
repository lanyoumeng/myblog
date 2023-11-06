// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package middleware

import (
	"blog/internal/pkg/core"
	"blog/internal/pkg/errno"
	"blog/internal/pkg/log"
	"blog/pkg/known"

	"github.com/gin-gonic/gin"
)

type Auther interface {
	Authorize(sub, obj, act string) (bool, error)
}

func Authz(a Auther) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sub := ctx.GetString(known.XUsernameKey)
		obj := ctx.Request.URL.Path
		act := ctx.Request.Method
		log.Debugw("Build authorize context", "sub", sub, "obj", obj, "act", act)

		if ok, _ := a.Authorize(sub, obj, act); !ok {
			core.WriteResponse(ctx, errno.ErrUnauthorized, nil)
			ctx.Abort()
			return
		}

	}

}
