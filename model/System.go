package model

import (
	"log"
	"yblog/constant"
)

// TableName 返回系统参数表名
func (system System) TableName() string {
	return constant.TableSystem
}

// AddSystem 增加一个系统信息
func (system System) AddSystem() error {
	err := DBClient.Create(&system).Error
	if err != nil {
		log.Fatalf("Add System Error: [%+v]", err)
		return err
	}

	return nil
}

// UpdateSystem 更新一个系统信息
func (system System) UpdateSystem() error {
	err := DBClient.Model(&system).Updates(system).Error
	if err != nil {
		log.Fatalf("Update System Error: [%+v]", err)
		return err
	}

	return nil
}
