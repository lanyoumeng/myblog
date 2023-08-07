// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.
package main

import "go.uber.org/zap"

func main() {
	//1.创建Logger时使用Fields(fs ...Field)选项
	logger := zap.NewExample(zap.Fields(
		zap.Int("userID", 10),
		zap.String("requestID", "fbf54504"),
	))
	logger.Info("This is a info message")

	//2. 不更改原始日志记录器的情况下，为特定的日志记录添加
	//logger, _ := zap.NewProduction(zap.Fields(zap.String("name", "dj"))) 已被废弃，推荐使用
	logger = logger.With(zap.String("app", "myapp"))
	logger.Info("This is a info message")
}
