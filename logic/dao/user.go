package dao

import (
	"chat/db"
	"errors"
	"time"
)

var dbSrc = db.GetDB("chat")

type User struct {
	Id         int `gorm:"primary_key"`
	UserName   string
	Password   string
	CreateTime time.Time
	db.DBChat
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Add() (userId int, err error) {

    // UserName and passowrd can't be null
	if u.UserName == "" || u.Password == "" {
		return 0, errors.New("user_name or password empty!")
	}

    // Check have the user name
	oUser := u.CheckHaveUserName(u.UserName)
	if oUser.Id > 0 {
		return oUser.Id, nil
	}

    // Create user
	u.CreateTime = time.Now()
	if err = dbSrc.Table(u.TableName()).Create(&u).Error; err != nil {
		return 0, err
	}
	return u.Id, nil
}

func (u *User) CheckHaveUserName(userName string) (data User) {
	dbSrc.Table(u.TableName()).Where("user_name=?", userName).Take(&data)
	return
}

func (u *User) GetUserNameByUserId(userId int) (userName string) {
	var data User
	dbSrc.Table(u.TableName()).Where("id=?", userId).Take(&data)
	return data.UserName
}


