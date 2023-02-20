package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

const PasswordCyptLevel = 12

type User struct {
	BaseModel
	UserName     string `gorm:"not null;unique_index;" json:"username"`
	Password     string `                              json:"-"`
	IsSuperAdmin bool   `                              json:"is_super_admin"`
	IsActive     bool   `                              json:"is_active"`
	LastLogin    time.Time
}

// 获取表名
func (user *User) TableName() string {
	return "user_basic"
}

// 根据ID获取用户
func (user *User) GetUserByID(id uint64) *User {
	DB.Model(&User{}).First(user, id)
	if user.ID > 0 {
		return user
	} else {
		return nil
	}
}

// 加密密码
func (user *User) SetPWD(password string) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCyptLevel)
	if err != nil {
		return err
	}
	user.Password = string(pwd)
	return nil
}

// 校验密码
func (user *User) CheckPWD() bool {
	pwd := user.Password
	DB.Where("username=?", user.UserName).First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))
	return err == nil
}

// 验证登录
func (user *User) IsPWDEqual(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))
	return err == nil
}

// 验证用户名重复
func (user *User) CheckDuplicateUserName() bool {
	var count int64
	if DB.Model(&User{}).Where("username = ?", user.UserName).Count(&count); count > 0 {
		return false
	} else {
		return true
	}
}

func (user *User) CheckUserNameExist(username string) bool {
	var count int64
	if DB.Model(&User{}).Where("username = ?", username).Count(&count); count > 0 {
		return false
	} else {
		return true
	}
}

// 根据ID判断用户是否存在
