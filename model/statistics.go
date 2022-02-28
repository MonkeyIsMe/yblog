package model

import (
	"log"
	"yblog/constant"
)

// TableName 统计表的名字
func (statistics Statistics) TableName() string {
	return constant.TableStatistics
}

// AddStatistics 添加一个统计
func (statistics Statistics) AddStatistics() error {
	err := DBClient.Create(&statistics).Error
	if err != nil {
		log.Fatalf("Add Statistics Error: [%+v]", err)
		return err
	}

	return nil
}
