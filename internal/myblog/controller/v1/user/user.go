// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package user

import (
	"blog/internal/myblog/biz"
	"blog/internal/myblog/store"
	"blog/pkg/auth"
	pb "blog/pkg/proto/myblog/v1"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type UserController struct {
	b biz.IBiz
	a *auth.Authz
	pb.UnimplementedMyBlogServer
}

// New 创建一个 user controller.
func New(ds store.IStore, authz *auth.Authz) *UserController {
	return &UserController{b: biz.NewBiz(ds), a: authz}
}
