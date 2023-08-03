// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/data", func(c *gin.Context) {
		data := gin.H{
			"key1": "value1",
			"key2": 123,
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":8080")

}
