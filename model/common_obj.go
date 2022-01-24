package model

import (
	"fmt"
	"yblog/constant"

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
	UserFlag     int    `gorm:"column:user_flag"`     // 用户的权限 -1代表账号不可用，0代表普通用户只有看和评论的权利，1代表管理员，有任何权利
	UserEmail    string `gorm:"column:user_email"`    // 用户邮箱
	HeadImage    string `gorm:"column:head_image"`    // 用户头像链接
	UserInfo     string `gorm:"column:user_info"`     // 用户简介
}

// Tag 标签信息
type Tag struct {
	TagID   int    `gorm:"column:tag_id"`   // 标签编号，主键
	TagName string `gorm:"column:tag_name"` // 标签名字
	TagInfo string `gorm:"column:tag_info"` // 标签简介
}

// TagBlog 标签和博文的关联表
type TagBlog struct {
	UnionID int `gorm:"column:union_id"` // 主键，关联编号
	BlogID  int `gorm:"column:blog_id"`  // 博客编号，博客的主键
	TagID   int `gorm:"column:tag_id"`   // 标签编号，标签的主键
}

// Blog 博客类
type Blog struct {
	BlogID       int    `gorm:"column:blog_id"`       // 博客编号，主键
	BlogTitle    string `gorm:"column:blog_title"`    // 博客名称
	BlogContent  string `gorm:"column:blog_content"`  // 博客内容
	CreateTime   string `gorm:"column:create_time"`   // 创建时间
	BlogAbstract string `gorm:"column:blog_abstract"` // 博客摘要
	BlogAuthor   string `gorm:"column:blog_author"`   // 博客作者
	IsValid      string `gorm:"column:is_valid"`      // 是否可见
}

// Classify 分类信息
type Classify struct {
	ClassifyID   int    `gorm:"column:classify_id"`   // 分类编号，主键
	ClassifyName int    `gorm:"column:classify_name"` // 分类名称
	ClassifyInfo string `gorm:"column:calssify_info"` // 分类信息
}

// ClassifyBlog 分类和博文的关联表
type ClassifyBlog struct {
	UnionID    int `gorm:"column:union_id"`    // 主键，关联编号
	BlogID     int `gorm:"column:blog_id"`     // 博客编号，博客的主键
	ClassifyID int `gorm:"column:classify_id"` // 分类编号，分类的主键
}

// Log 操作日志的基础类
type Log struct {
	LogID        int                  `gorm:"column:log_id"`        // 主键，日志编号
	LogOperation int                  `gorm:"column:log_operation"` // 日志操作
	LogDetail    constant.LogOpDetail `gorm:"column:log_detail"`    // 日志描述
	LogOperator  string               `gorm:"column:log_operator"`  // 操作人
	LogTime      string               `gorm:"column:log_time"`      // 操作时间
	LogObj       int                  `gorm:"column:log_obj"`       // 操作对象
	LogFlag      int                  `gorm:"column:log_flag"`      // 操作成功与否
	ErrorMsg     string               `gorm:"column:error_msg"`     // 错误信息
}

// System 系统信息
type System struct {
	SystemID   int    `gorm:"column:system_id"`   // 主键，系统编号，一个博客只有一个SystemID
	SystemName string `gorm:"column:system_name"` // 博客的名称
	HeadImage  string `gorm:"column:head_image"`  // 博客的头像
	SystemInfo string `gorm:"column:system_info"` // 博客的介绍
}
