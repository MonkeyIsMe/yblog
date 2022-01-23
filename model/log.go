package model

import (
	tconstant "github.com/MonkeyIsMe/MyTool/constant"
	"log"
	"yblog/constant"
)

// TableName 返回日志表名
func (log Log) TableName() string {
	return constant.TableLog
}

// AddLog 增加一个日志
func (l Log) AddLog() error {
	err := DBClient.Create(&l).Error
	if err != nil {
		log.Fatalf("Add Log Error: [%+v]", err)
		return err
	}

	return nil
}

// QueryLogPageSize 分页查询日志
func QueryLogPageSize(page int) ([]Log, error) {
	var logList []Log
	err := DBClient.Offset((page - 1) * tconstant.PageSize).Limit(tconstant.PageSize).Find(&logList).Error
	if err != nil {
		log.Fatalf("Query Log PageSize Error: [%+v]", err)
		return nil, err
	}

	return logList, nil
}

// CountLog 查询所有日志的数量
func CountLog() (int64, error) {
	var count int64
	err := DBClient.Table(constant.TableLog).Count(&count).Error
	if err != nil {
		log.Fatalf("Count Log Error: [%+v]", err)
		return 0, err
	}

	return count, nil
}
