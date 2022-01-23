package model

import (
	"log"
	"yblog/constant"

	tconstant "github.com/MonkeyIsMe/MyTool/constant"
)

func (tag Tag) TableName() string {
	return constant.TableTag
}

func (tb TagBlog) TableName() string {
	return constant.TableTagBlog
}

// AddTag 插入一个分类标签
func (tag Tag) AddTag() error {
	err := DBClient.Create(&tag).Error
	if err != nil {
		log.Fatalf("Add Tag Erorr: [%+v]", err)
		return err
	}
	return nil
}

// DeleteTag 删除一个分类标签
func (tag Tag) DeleteTag() error {
	err := DBClient.Delete(&tag).Error
	if err != nil {
		log.Fatalf("Delete Tag Error: [%+v]", err)
		return err
	}
	return nil
}

// QueryTagPageSize 分页查询tag
func QueryTagPageSize(page int) ([]Tag, error) {
	var tagList []Tag
	err := DBClient.Offset((page - 1) * tconstant.PageSize).Limit(tconstant.PageSize).Find(&tagList).Error
	if err != nil {
		log.Fatalf("Query Tag PageSize Error: [%+v]", err)
		return nil, err
	}
	return tagList, nil
}

// AddTagBlog 插入一个博客和标签的关联
func (tb TagBlog) AddTagBlog() error {
	err := DBClient.Create(&tb).Error
	if err != nil {
		log.Fatalf("Add TagBlog Erorr: [%+v]", err)
		return err
	}
	return nil
}

// DeleteTagBlog 删除一个博客和标签的关联
func (tb TagBlog) DeleteTagBlog() error {
	err := DBClient.Delete(&tb).Error
	if err != nil {
		log.Fatalf("Delete TagBlog Erorr: [%+v]", err)
		return err
	}
	return nil
}
