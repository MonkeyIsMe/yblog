package model

import (
	"log"
	"yblog/constant"

	tconstant "github.com/MonkeyIsMe/MyTool/constant"
)

// TableName 返回一个评论表名
func (comment Comment) TableName() string {
	return constant.TableComment
}

// AddComment 增加一条评论
func (comment Comment) AddComment() error {
	err := DBClient.Create(&comment).Error
	if err != nil {
		log.Fatalf("Add Comment Error: [%+v]", err)
		return err
	}

	return nil
}

// QueryCommentPageSize 分页查询评论
func QueryCommentPageSize(page int) ([]Comment, error) {
	var commentList []Comment
	err := DBClient.Offset((page - 1) * tconstant.PageSize).Limit(tconstant.PageSize).Find(&commentList).Error
	if err != nil {
		log.Fatalf("Query Comment PageSize Error: [%+v]", err)
		return nil, err
	}

	return commentList, nil
}

// QueryCommentByStatus 分页不合法查询评论
func QueryCommentByStatus(page, valid int) ([]Comment, error) {
	var commentList []Comment
	err := DBClient.Where("valid = ?", valid).Offset((page - 1) *
		tconstant.PageSize).Limit(tconstant.PageSize).Find(&commentList).Error
	if err != nil {
		log.Fatalf("Query Invalid Comment PageSize Error: [%+v]", err)
		return nil, err
	}

	return commentList, nil
}

// SetCommentStatus 这是评论的状态
func (comment Comment) SetCommentStatus(valid int) error {
	err := DBClient.Model(&comment).Update("valid = ?", valid).Error
	if err != nil {
		log.Fatalf("Set Comment Valid Error: [%+v]", err)
		return err
	}

	return nil
}

// CountComment 查询所有评论的数量
func CountComment() (int64, error) {
	var count int64
	err := DBClient.Table(constant.TableComment).Count(&count).Error
	if err != nil {
		log.Fatalf("Count Comment Error: [%+v]", err)
		return 0, err
	}

	return count, nil
}

// QueryBlogCommentPageSize 查询某个博客的合法评论
func QueryBlogCommentPageSize(page int, blogID int) ([]Comment, error) {
	var commentList []Comment
	err := DBClient.Where("valid = ? AND blog_id = ?", "1", blogID).Offset((page - 1) *
		tconstant.PageSize).Limit(tconstant.PageSize).Find(&commentList).Error
	if err != nil {
		log.Fatalf("Query Comment PageSize Error: [%+v]", err)
		return nil, err
	}

	return commentList, nil
}

// DeleteComment 删除一条评论
func (comment Comment) DeleteComment() error {
	err := DBClient.Delete(&comment).Error
	if err != nil {
		log.Fatalf("Add Comment Error: [%+v]", err)
		return err
	}

	return nil
}

// QueryCommentByID 根据评论主键查询评论的信息
func QueryCommentByID(commentID int) (Comment, error) {
	commentInfo := Comment{}
	err := DBClient.First(&commentInfo, commentID).Error

	if err != nil {
		log.Fatalf("Query Single Comment Error: [%+v]", err)
		return commentInfo, err
	}

	return commentInfo, nil
}
