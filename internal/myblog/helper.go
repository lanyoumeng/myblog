// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package myblog

import (
	"blog/internal/myblog/store"
	"blog/internal/pkg/log"
	"blog/pkg/db"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfigName = "myblog"
	// recommendedHomeDir = "workspace/lanmengyou/myblog/configs"

	recommendedHomeDir = ".myblog"
)

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))

		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(defaultConfigName)
	}
	viper.AutomaticEnv()

	viper.SetEnvPrefix("MYBLOG")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read viper configuration file", "err", err)
	}
	log.Infow("Using config file", "file", viper.ConfigFileUsed())
}

func logOptions() *log.Options {

	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log. disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}

func initStore() error {
	dbOptions := &db.MySQLOptions{Host: viper.GetString("db.host"),
		Username:              viper.GetString("db.username"),
		Password:              viper.GetString("db.password"),
		Database:              viper.GetString("db.database"),
		MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
		LogLevel:              viper.GetInt("db.log-level")}

	ins, err := db.NewMySQL(dbOptions)
	if err != nil {
		return err
	}
	_ = store.NewStore(ins)

	return nil
}
