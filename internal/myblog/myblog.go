// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package myblog

import (
	"blog/internal/pkg/log"
	"blog/pkg/known"
	"blog/pkg/token"
	"blog/pkg/version/verflag"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	mw "blog/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

func NewMyBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "myblog",
		Short: "The firsr project(blog)",
		Long: `A good Go practical project, used to create user with basic information.

Find more miniblog information at:
        https://github.com/lanyoumeng/myblog.git`,

		SilenceUsage: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			// 如果 `--version=true`，则打印版本并退出
			verflag.PrintAndExitIfRequested()

			log.Init(logOptions())
			defer log.Sync()

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

	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.myblog.yaml)")

	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// 添加 --version 标志
	verflag.AddFlags(cmd.PersistentFlags())
	return cmd

}

func run() error {
	if err := initStore(); err != nil {
		return err
	}

	// 设置 token 包的签发密钥，用于 token 包 token 的签发和解析
	token.Init(viper.GetString("jwt-secret"), known.XUsernameKey)

	gin.SetMode(viper.GetString("runmode"))

	router := gin.New()

	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.Secure, mw.RequestID()}
	router.Use(mws...)

	if err := installRouters(router); err != nil {
		return err
	}

	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: router}
	log.Infow("Start to  listening the incoming request on  http address", "addr", viper.GetString("addr"))
	// if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	// 	log.Fatalw(err.Error())
	// }
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Infow("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}
	log.Infow("Server exiting")

	return nil
}
