// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package core

import (
	"blog/internal/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func WriteResponse(ctx *gin.Context, err error, data interface{}) {
	if err != nil {
		hcode, code, message := errno.Decode(err)
		ctx.JSON(hcode, ErrResponse{
			Code:    code,
			Message: message,
		})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
