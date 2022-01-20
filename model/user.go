package model

import (
	"log"
	"yblog/constant"
)

// todo 需要大量补充

func (user User) TableName() string {
	return constant.TableUser
}

// AddUser 增加一个用户
func (user User) AddUser() error {
	err := DBClient.Create(&user).Error
	if err != nil {
		log.Fatalf("Add User Error: [%+v]", err)
		return err
	}

	return nil
}

// DeleteUser 删除一个用户
func (user User) DeleteUser() error {
	err := DBClient.Delete(&user).Error
	if err != nil {
		log.Fatalf("Delete User Error: [%+v]", err)
		return err
	}

	return nil
}
