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

// QueryStatisticsByTime 根据日期查询统计数据
func QueryStatisticsByTime(time string) (Statistics, error) {
	statisticsInfo := Statistics{}
	err := DBClient.First(&statisticsInfo, "statistics_day = ?", time).Error

	if err != nil {
		log.Fatalf("Query Single Statistics Error: [%+v]", err)
		return statisticsInfo, err
	}

	return statisticsInfo, nil
}
