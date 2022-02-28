package service

import (
	"encoding/json"
	"log"
	"yblog/constant"
	"yblog/model"
)

// UpdateBlog 更新博客
func UpdateBlog(blog model.Blog) error {
	blogID := blog.BlogID
	oldBlog, err := model.QueryBlogByID(blogID)
	if err != nil {
		log.Fatalf("Query Blog In Update Blog err : [%+v]", err)
		return err
	}
	oldBlogJson, err := json.Marshal(oldBlog)
	if err != nil {
		log.Fatalf("Marshal Old Blog err : [%+v]", err)
		return err
	}
	newBlogJson, err := json.Marshal(blog)
	if err != nil {
		log.Fatalf("Marshal Blog err : [%+v]", err)
		return err
	}
	err = AddLog(oldBlogJson, newBlogJson, constant.LogUpdate, constant.LogDetail)
	if err != nil {
		log.Fatalf("Add Log In Update Blog err : [%+v]", err)
		return err
	}
	return nil
}

// QueryValidBlog 查询可见的博客
func QueryValidBlog(page, size int) ([]model.Blog, error) {
	blogList, err := model.QueryBlogByStatus(page, constant.Normal, size)
	if err != nil {
		log.Fatalf("BlogService Query Valid Blog err: [%+v]", err)
		return nil, err
	}
	return blogList, nil
}

// QueryInvalidBlog 查询不可见的博客
func QueryInvalidBlog(page, size int) ([]model.Blog, error) {
	blogList, err := model.QueryBlogByStatus(page, constant.UnSee, size)
	if err != nil {
		log.Fatalf("BlogService Query Invalid Blog err: [%+v]", err)
		return nil, err
	}
	return blogList, nil
}

// QueryAllBlog 查询所有的博客
func QueryAllBlog(page, size int) ([]model.Blog, error) {
	blogList, err := model.QueryBlogPageSize(page, size)
	if err != nil {
		log.Fatalf("BlogService Query Blog PageSize err: [%+v]", err)
		return nil, err
	}
	return blogList, nil
}
