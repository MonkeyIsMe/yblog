package model

import (
	tconstant "github.com/MonkeyIsMe/MyTool/constant"
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
func QueryBlogPageSize(page int) ([]Blog, error) {
	var blogList []Blog
	err := DBClient.Offset((page - 1) * tconstant.PageSize).Limit(tconstant.PageSize).Find(&blogList).Error
	if err != nil {
		log.Fatalf("Query Blog PageSize Error: [%+v]", err)
		return nil, err
	}

	return blogList, nil
}

// QueryValidBlogPageSize 分页查询可见博客
func QueryValidBlogPageSize(page int) ([]Blog, error) {
	var blogList []Blog
	err := DBClient.Where("valid = ?", "1").Offset((page - 1) *
		tconstant.PageSize).Limit(tconstant.PageSize).Find(&blogList).Error
	if err != nil {
		log.Fatalf("Query Valid Blog PageSize Error: [%+v]", err)
		return nil, err
	}

	return blogList, nil
}

// SetBlogValid 将博客置为可见
func (blog Blog) SetBlogValid() error {
	err := DBClient.Model(&blog).Update("valid = ?", "1").Error
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
