package service

import (
	"log"
	"yblog/model"
)

// AddBlogClassfiy 增加一个博客和分类的关联
func AddBlogClassfiy(cb model.ClassifyBlog) error {
	err := cb.AddClassifyBlog()
	if err != nil {
		log.Fatalf("Link Service ClassfiyBlog Error: [%+v", err)
	}
	return nil
}

// AddBlogTag 增加一个博客和标签的关联
func AddBlogTag(tb model.TagBlog) error {
	err := tb.AddTagBlog()
	if err != nil {
		log.Fatalf("Link Service TagBlog Error: [%+v", err)
	}
	return nil
}
