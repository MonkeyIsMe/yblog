package model

import (
	"log"
	"yblog/constant"
)

func (classify Classify) TableName() string {
	return constant.TableClassify
}

func (cb ClassifyBlog) TableName() string {
	return constant.TableClassifyBlog
}

// AddClassify 增加一个分类
func (classify Classify) AddClassify() error {
	err := DBClient.Create(&classify).Error
	if err != nil {
		log.Fatalf("Add Classify Error: [%+v]", err)
		return err
	}

	return nil
}

// DeleteClassify 删除一个分类
func (classify Classify) DeleteClassify() error {
	err := DBClient.Delete(&classify).Error
	if err != nil {
		log.Fatalf("Delete Classify Error: [%+v]", err)
		return err
	}

	return nil
}

// AddClassifyBlog 增加一个博客的分类
func (cb ClassifyBlog) AddClassifyBlog() error {
	err := DBClient.Create(&cb).Error
	if err != nil {
		log.Fatalf("Add ClassifyBlog Error: [%+v]", err)
		return err
	}

	return nil
}

// DeleteClassifyBlog 删除一个博客的分类
func (cb ClassifyBlog) DeleteClassifyBlog() error {
	err := DBClient.Delete(&cb).Error
	if err != nil {
		log.Fatalf("Delete ClassifyBlog Error: [%+v]", err)
		return err
	}

	return nil
}

// QueryClassifyByID 根据分类主键查询分类的信息
func QueryClassifyByID(classifyID int) (Classify, error) {
	classifyInfo := Classify{}
	err := DBClient.First(&classifyInfo, classifyID).Error

	if err != nil {
		log.Fatalf("Query Single Classfiy Error: [%+v]", err)
		return classifyInfo, err
	}

	return classifyInfo, nil
}

// CountClassfiy 查询所有分类总数
func CountClassfiy() (int64, error) {
	var count int64
	err := DBClient.Table(constant.TableClassify).Count(&count).Error
	if err != nil {
		log.Fatalf("Count Classfiy Error: [%+v]", err)
		return 0, err
	}

	return count, nil
}
