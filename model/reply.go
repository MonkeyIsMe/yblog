package model

import (
	"log"
	"yblog/constant"
)

// TableName 回复表的名字
func (reply Reply) TableName() string {
	return constant.TableReply
}

// AddReply 添加一个回复
func (reply Reply) AddReply() error {
	err := DBClient.Create(&reply).Error
	if err != nil {
		log.Fatalf("Add Reply Error: [%+v]", err)
		return err
	}

	return nil
}

// DeleteReply 删除一个回复
func (reply Reply) DeleteReply() error {
	err := DBClient.Delete(&reply).Error
	if err != nil {
		log.Fatalf("Delete Reply Error: [%+v]", err)
		return err
	}

	return nil
}

// SetReplyStatus 设置回复的状态
func (reply Reply) SetReplyStatus(valid int) error {
	err := DBClient.Model(&reply).Update("valid = ?", valid).Error
	if err != nil {
		log.Fatalf("Set Relpy Valid Error: [%+v]", err)
		return err
	}

	return nil
}
