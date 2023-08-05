// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction(zap.AddCaller()) // 创建一个Zap日志记录器,并增加行号和文件名
	defer logger.Sync()                             // 刷新磁盘

	url := "http://marmotedu.com"
	logger.Info("failed to fetch URL", // 结构化日志记录
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	logger.Info("hello world") // 使用日志记录器输出日志

	sugar := logger.Sugar()
	sugar = sugar.With(zap.String("name", "lmy")) // 添加公共字段。

	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)

	sugar.Info("HEllo ")                        //一个字符串参数
	sugar.Infof("Failed to fetch URL: %s", url) //f-->传参
	logger.Sugar().Infow("creat", "lmy")        //w--> 一个或多个 key-value 对
}
