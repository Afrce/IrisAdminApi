package models

import (
	"IrisAdminApi/sysinit"
)

type UserHasPermission struct {
	ID           uint
	PermissionID uint
	UserID       uint
}

func (UserHasPermission) TableName() string {
	return "user_has_permission"
}

func CheckUserHasPermission(permissionID uint, UserId uint) bool {
	userHasPermission := new(UserHasPermission)
	if err := sysinit.DB.Where(&UserHasPermission{PermissionID: permissionID, UserID: UserId}).First(&userHasPermission).Error; err != nil {
		return false
	} else {
		return true
	}
}
