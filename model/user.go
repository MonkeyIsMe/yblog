package model

import (
	"log"
	"yblog/constant"
)

// TableName 返回用户表名
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
func QueryUserPageSize(page, size int) ([]User, error) {
	var userList []User
	err := DBClient.Offset((page - 1) * size).Limit(size).Find(&userList).Error
	if err != nil {
		log.Fatalf("Query User PageSize Error: [%+v]", err)
		return nil, err
	}

	return userList, nil
}

// UpdateUser 更新一个用户
func (user User) UpdateUser() error {
	err := DBClient.Model(&user).Updates(user).Error
	if err != nil {
		log.Fatalf("Update User Error: [%+v]", err)
		return err
	}

	return nil
}

// UpdateUserPassword 更新用户的密码
func (user User) UpdateUserPassword(newPassword string) error {
	err := DBClient.Update("user_password", newPassword).Error
	if err != nil {
		log.Fatalf("Update User Password Error: [%+v]", err)
		return err
	}

	return nil
}

// UpdateUserIcon 更新用户的头像
func (user User) UpdateUserIcon(iconURL string) error {
	err := DBClient.Update("head_image", iconURL).Error
	if err != nil {
		log.Fatalf("Update User Head Image Error: [%+v]", err)
		return err
	}

	return nil
}

// QueryUserByID 根据用户主键查询用户的信息
func QueryUserByID(userID int) (User, error) {
	userInfo := User{}
	err := DBClient.First(&userInfo, userID).Error

	if err != nil {
		log.Fatalf("Query Single User Error: [%+v]", err)
		return userInfo, err
	}

	return userInfo, nil
}
