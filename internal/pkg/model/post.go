// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package model

import (
	"blog/pkg/util/id"
	"time"

	"gorm.io/gorm"
)

type PostM struct {
	ID        int64     `gorm:"column:id;primary_key"` //
	Username  string    `gorm:"column:username;not null"`
	PostID    string    `gorm:"column:postID;not null"`
	Title     string    `gorm:"column:title;not null"`
	Content   string    `gorm:"column:content"`   //
	CreatedAt time.Time `gorm:"column:createdAt"` //
	UpdatedAt time.Time `gorm:"column:updatedAt"` //
}

// TableName sets the insert table name for this struct type
func (p *PostM) TableName() string {
	return "post"
}

// BeforeCreate 在创建数据库记录之前生成 postID.
func (p *PostM) BeforeCreate(tx *gorm.DB) error {
	p.PostID = "post-" + id.GenShortID()

	return nil
}
