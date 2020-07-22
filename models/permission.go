package models

import (
	"IrisAdminApi/sysinit"
	"github.com/jinzhu/gorm"
)

type Permission struct {
	gorm.Model
	Permission string `json:"permission"`
}

func (Permission) TableName() string {
	return "permissions"
}

func SelectPermission(permissionName string) (uint, bool) {
	permission := new(Permission)
	if err := sysinit.DB.Where(&Permission{Permission: permissionName}).First(&permission).Error; err != nil {
		return 0, false
	}
	return permission.ID, true
}
