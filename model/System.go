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

// QueryFirstSystem 查询第一个系统设置
func (system System) QueryFirstSystem() error {
	err := DBClient.First(&system).Error
	if err != nil {
		log.Fatalf("Query System Info Error: [%+v]", err)
		return err
	}

	return nil
}

// QuerySystemByID 根据系统信息主键查询系统的信息
func QuerySystemByID(systemID int) (System, error) {
	systemInfo := System{}
	err := DBClient.First(&systemInfo, systemID).Error

	if err != nil {
		log.Fatalf("Update System Error: [%+v]", err)
		return systemInfo, err
	}

	return systemInfo, nil
}