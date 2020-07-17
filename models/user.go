package models

import (
	"IrisAdminApi/sysinit"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(15);unique_index"`
	Password string `gorm:"type:varchar(225)"`
	Email    string `gorm:"ype:varchar(100);unique_index"`
}

func (u User) TableName() string {
	return "users"
}

func (u *User) GetUserByUsername() {
	sysinit.DB.Where("name = ?", u.Name).First(&u)
}
