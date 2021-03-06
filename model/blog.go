package model

import (
	"log"
	"yblog/constant"
)

// TableName 博客表的名字
func (blog Blog) TableName() string {
	return constant.TableBlog
}

// AddBlog 添加一个博客
func (blog Blog) AddBlog() error {
	err := DBClient.Create(&blog).Error
	if err != nil {
		log.Fatalf("Add Blog Error: [%+v]", err)
		return err
	}

	return nil
}

// DeleteBlog 删除一个博客
func (blog Blog) DeleteBlog() error {
	err := DBClient.Delete(&blog).Error
	if err != nil {
		log.Fatalf("Delete Blog Error: [%+v]", err)
		return err
	}

	return nil
}

// QueryBlogPageSize 分页查询博客
func QueryBlogPageSize(page, size int) ([]Blog, error) {
	var blogList []Blog
	err := DBClient.Offset((page - 1) * size).Limit(size).Find(&blogList).Error
	if err != nil {
		log.Fatalf("Query Blog PageSize Error: [%+v]", err)
		return nil, err
	}

	return blogList, nil
}

// QueryBlogByStatus 根据状态分页查询博客
func QueryBlogByStatus(page, valid, size int) ([]Blog, error) {
	var blogList []Blog
	err := DBClient.Where("valid = ?", valid).Offset((page - 1) * size).Limit(size).Find(&blogList).Error
	if err != nil {
		log.Fatalf("Query Blog By Status Error: [%+v]", err)
		return nil, err
	}

	return blogList, nil
}

// SetBlogValid 修改博客状态
func (blog Blog) SetBlogValid(flag int) error {
	err := DBClient.Model(&blog).Update("valid = ?", flag).Error
	if err != nil {
		log.Fatalf("Set Blog Valid Error: [%+v]", err)
		return err
	}

	return nil
}

// CountBlog 查询所有博客的数量
func CountBlog() (int64, error) {
	var count int64
	err := DBClient.Table(constant.TableBlog).Count(&count).Error
	if err != nil {
		log.Fatalf("Count Blog Error: [%+v]", err)
		return 0, err
	}

	return count, nil
}

// UpdateBlog 更新一个博客
func (blog Blog) UpdateBlog() error {
	err := DBClient.Model(&blog).Updates(blog).Error
	if err != nil {
		log.Fatalf("Update Blog Error: [%+v]", err)
		return err
	}

	return nil
}

// QueryBlogByID 根据博客主键查询博客的信息
func QueryBlogByID(blogID int) (Blog, error) {
	blogInfo := Blog{}
	err := DBClient.First(&blogInfo, blogID).Error

	if err != nil {
		log.Fatalf("Query Single Blog Error: [%+v]", err)
		return blogInfo, err
	}

	return blogInfo, nil
}
