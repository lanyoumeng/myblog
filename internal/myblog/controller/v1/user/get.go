// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package user

import (
	"blog/internal/pkg/core"
	"blog/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) Get(c *gin.Context) {
	log.C(c).Infow("Get user function called")
	user, err := ctrl.b.Users().Get(c, c.Param("name"))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, user)
}
