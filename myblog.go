// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package main

import (
	"flag"
	"fmt"
)

var (
	// GitVersion 是语义化的版本号.
	GitVersion = "v0.0.0-master+$Format:%h$"
	// BuildDate 是 ISO8601 格式的构建时间, $(date -u +'%Y-%m-%dT%H:%M:%SZ') 命令的输出.
	BuildDate = "1970-01-01T00:00:00Z"
)

func main() {
	version := flag.Bool("version", false, "Print version info.")
	flag.Parse()

	if *version {
		fmt.Println("GitVersion", GitVersion)
		fmt.Println("BuildDate", BuildDate)
	}

	fmt.Println("ok")
}
