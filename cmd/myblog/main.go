// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package main

import (
	// "command-line-arguments/home/goer/workspace/lanmengyou/myblog/internal/myblog/myblog.go"

	"blog/internal/myblog"
	"os"

	_ "go.uber.org/automaxprocs"
	// "../internal/myblog"
	// "lanmengyou/myblog/internal/myblog"
)

// Go 程序的默认入口函数(主函数).

func main() {
	command := myblog.NewMyBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
