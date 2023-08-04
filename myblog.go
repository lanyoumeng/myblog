// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version string // 这里定义版本信息
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A brief description of your application",
	Long:  "A longer description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, this is the root command!")
	},
}

func init() {
	rootCmd.Flags().StringVarP(&version, "version", "v", "", "Print the version number")

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
