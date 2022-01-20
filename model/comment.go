package model

import (
	"log"
	"yblog/constant"

	tconstant "github.com/MonkeyIsMe/MyTool/constant"
)

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
		log.Fatalf("Query Tag PageSize Err: [%+v]", err)
		return nil, err
	}

	return commentList, nil
}

// QueryValidCommentPageSize 分页合法查询评论
func QueryValidCommentPageSize(page int) ([]Comment, error) {
	var commentList []Comment
	err := DBClient.Where("valid = ?", "1").Offset((page - 1) *
		tconstant.PageSize).Limit(tconstant.PageSize).Find(&commentList).Error
	if err != nil {
		log.Fatalf("Query Tag PageSize Err: [%+v]", err)
		return nil, err
	}

	return commentList, nil
}

// QueryInvalidCommentPageSize 分页不合法查询评论
func QueryInvalidCommentPageSize(page int) ([]Comment, error) {
	var commentList []Comment
	err := DBClient.Where("valid = ?", "0").Offset((page - 1) *
		tconstant.PageSize).Limit(tconstant.PageSize).Find(&commentList).Error
	if err != nil {
		log.Fatalf("Query Tag PageSize Err: [%+v]", err)
		return nil, err
	}

	return commentList, nil
}

// SetCommentValid 将评论置为合法
func (comment Comment) SetCommentValid() error {
	err := DBClient.Model(&comment).Update("valid = ?", "1").Error
	if err != nil {
		log.Fatalf("Set Comment Valid Err: [%+v]", err)
		return err
	}

	return nil
}

// CountComment 查询所有评论的数量
func CountComment() (int64, error) {
	var count int64
	err := DBClient.Table(constant.TableComment).Count(&count).Error
	if err != nil {
		log.Fatalf("Count Comment Err: [%+v]", err)
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
		log.Fatalf("Query Tag PageSize Err: [%+v]", err)
		return nil, err
	}

	return commentList, nil
}