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
	CommentID   int    `gorm:"column:comment_id"`   // 主键ID
	CommentInfo string `gorm:"column:comment_info"` // 评论信息
	CreateTime  string `gorm:"column:create_time"`  // 发表时间
	UserName    string `gorm:"column:user_name"`    // 发表用户
	IsValid     int    `gorm:"column:is_valid"`     // 是否合法
	BlogID      int    `gorm:"column:blog_id"`      // 博客编号
}

// User 用户类
type User struct {
	UserID       int    `gorm:"column:user_id"`       // 用户编号
	UserName     string `gorm:"column:user_name"`     // 用户姓名
	UserPassword string `gorm:"column:user_password"` // 密码
	CreateTime   string `gorm:"column:create_time"`   // 创建时间
	UserAccount  string `gorm:"column:user_account"`  // 账号
	IsValid      int    `gorm:"column:is_valid"`      // 账号是否可用
	UserFlag     int    `gorm:"column:user_flag"`     // 用户的权限
	UserEmail    string `gorm:"column:user_email"`    // 用户邮箱
	UserImage    string `gorm:"column:user_image"`    // 用户头像链接
	UserInfo     string `gorm:"column:user_info"`     // 用户简介
}

// Tag 标签信息
type Tag struct {
	TagID   int    `gorm:"column:tag_id"`   // 标签编号，主键
	TagName string `gorm:"column:tag_name"` // 标签名字
	TagInfo string `gorm:"column:tag_info"` // 标签简介
}

// TagBlog 分类和博文的关联表
type TagBlog struct {
	UnionID int `gorm:"column:union_id"` // 主键，关联编号
	BlogID  int `gorm:"column:blog_id"`  // 博客编号，博客的主键
	TagID   int `gorm:"column:tag_id"`   // 标签编号，标签的主键
}

// Blog 博客类
type Blog struct {
	BlogID       int    `gorm:"column:blog_id"`       // 博客编号，主键
	BlogName     string `gorm:"column:blog_name"`     // 博客名称
	BlogContent  string `gorm:"column:blog_content"`  // 博客内容
	CreateTime   string `gorm:"column:create_time"`   // 创建时间
	BlogAbstract string `gorm:"column:blog_abstract"` // 博客摘要
	BlogAuthor   string `gorm:"column:blog_author"`   // 博客作者
	IsValid      string `gorm:"column:is_valid"`      // 是否可见
}
