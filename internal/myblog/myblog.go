// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package myblog

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewMyBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "myblog",

		Short: "The firsr project(blog)",
		Long: `A good Go practical project, used to create user with basic information.

Find more miniblog information at:
        https://github.com/lanyoumeng/myblog.git`,

		SilenceUsage: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments ,got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	return cmd

}

func run() error {
	fmt.Println("Hello")
	return nil
}
