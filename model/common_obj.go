package model

import (
	"fmt"

	"gorm.io/gorm"
)

var DBClient *gorm.DB

// SetClient 初始化DB的client
func SetClient(input *gorm.DB) error {
	if input == nil {
		err := fmt.Errorf("input client is nil")
		return err
	}
	DBClient = input
	return nil
}

// Comment 评论类
type Comment struct {
	CommentID   int    // 主键ID
	CommentInfo string // 评论信息
	CreateTime  string // 发表时间
	UserName    string // 发表用户
	IsValid     int    // 是否合法
	BlogID      int    // 博客编号
}

// User 用户类
type User struct {
	UserID       int    // 用户编号
	UserName     string // 用户姓名
	UserPassword string // 密码
	CreateTime   string // 创建时间
	UserAccount  string // 账号
	IsValid      int    // 账号是否可用
	UserFlag     int    // 用户的权限
	UserEmail    string // 用户邮箱
	UserImage    string // 用户头像链接
	UserInfo     string // 用户简介
}

// Tag 标签信息
type Tag struct {
	TagID   int    `gorm:"column:tag_id"`
	TagName string `gorm:"column:tag_name"`
	TagInfo string `gorm:"column:tag_info"`
}
