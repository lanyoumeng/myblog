// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package main

import (
	"blog/internal/pkg/log"
	pb "blog/pkg/proto/myblog/v1"
	"context"
	"fmt"
	"time"

	"github.com/bytedance/sonic"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr  = pflag.String("addr", "localhost:9090", "The adress to connect to.")
	limit = pflag.Int64("limit", 10, "limit to list users.")
)

func main() {
	pflag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalw("Did not connect", "err", err)
	}
	defer conn.Close()

	c := pb.NewMyBlogClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ListUser(ctx, &pb.ListUserRequest{Offset: 0, Limit: *limit})
	if err != nil {
		log.Fatalw("could not greet: %v", err)

	}

	// 打印请求结果
	fmt.Println("TotalCount:", r.TotalCount)
	for _, u := range r.Users {
		d, _ := sonic.Marshal(u)
		fmt.Println(string(d))
	}
}
