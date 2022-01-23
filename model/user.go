package model

import (
	"log"
	"yblog/constant"

	tconstant "github.com/MonkeyIsMe/MyTool/constant"
)

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

// SetUserFlag 修改用户权限
func (user User) SetUserFlag(flag int) error {
	err := DBClient.Model(&user).Update("user_flag = ?", flag).Error
	if err != nil {
		log.Fatalf("Set User Flag Error: [%+v]", err)
		return err
	}

	return nil
}

// CountUser 查询所有用户的数量
func CountUser() (int64, error) {
	var count int64
	err := DBClient.Table(constant.TableUser).Count(&count).Error
	if err != nil {
		log.Fatalf("Count User Error: [%+v]", err)
		return 0, err
	}

	return count, nil
}

// QueryUserPageSize 分页查询用户
func QueryUserPageSize(page int) ([]User, error) {
	var userList []User
	err := DBClient.Offset((page - 1) * tconstant.PageSize).Limit(tconstant.PageSize).Find(&userList).Error
	if err != nil {
		log.Fatalf("Query User PageSize Error: [%+v]", err)
		return nil, err
	}

	return userList, nil
}

// todo 更新用户信息
