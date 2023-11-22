// Copyright 2023 Innkeeper lanmengyou(殷家豪)&lt;2058818914@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lanyoumeng/myblog.git.

package user

import (
	"blog/internal/myblog/store"
	v1 "blog/pkg/api/myblog/v1"
	"context"
	"testing"
)

func Test_userBiz_Create(t *testing.T) {
	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx context.Context
		r   *v1.CreateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &userBiz{
				ds: tt.fields.ds,
			}
			if err := b.Create(tt.args.ctx, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("userBiz.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
