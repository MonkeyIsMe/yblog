package model

import (
	"log"
	"yblog/constant"
)

// TableName 链接表的名字
func (link Links) TableName() string {
	return constant.TableLinks
}

// AddLinks 添加一个友情链接
func (link Links) AddLinks() error {
	err := DBClient.Create(&link).Error
	if err != nil {
		log.Fatalf("Add Links Error: [%+v]", err)
		return err
	}

	return nil
}

// UpdateLinks 更新一个友情链接
func (link Links) UpdateLinks() error {
	err := DBClient.Model(&link).Updates(link).Error
	if err != nil {
		log.Fatalf("Update Links Error: [%+v]", err)
		return err
	}

	return nil
}

// QueryLinksByID 根据主键查询友情链接
func (link Links) QueryLinksByID(id int) error {
	err := DBClient.First(&link, id).Error
	if err != nil {
		log.Fatalf("Query Links By ID Info Error: [%+v]", err)
		return err
	}

	return nil
}
