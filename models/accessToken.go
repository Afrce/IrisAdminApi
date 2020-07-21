package models

import (
	"IrisAdminApi/sysinit"
	"fmt"
	"time"
)

type AccessToken struct {
	ID        int64 `gorm:"primary_key"`
	UserID    uint
	Token     string
	ExpiresAt time.Time
}

func (u AccessToken) TableName() string {
	return "access_token"
}

// 新增token
func CreateToken(token AccessToken) {
	sysinit.DB.Create(&token)
}

// 移除token
func DeleteToken(token string, UserId uint) {
	sysinit.DB.Where("token = ?", token).Where("user_id = ?", UserId).Delete(&AccessToken{})
}

func SelectToken(token string, UserId uint) bool {
	fmt.Println(UserId)
	accesstoken := new(AccessToken)
	if err := sysinit.DB.Where(&AccessToken{Token: token, UserID: UserId}).Find(&accesstoken).Error; err != nil {
		return false
	}
	return true
}
