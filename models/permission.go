package models

import (
	"IrisAdminApi/sysinit"
	"github.com/jinzhu/gorm"
)

type Permission struct {
	gorm.Model
	Permission string `json:"permission"`
}

type PermissionData struct {
	ID         uint   `json:"id"`
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

func GetPermissionList(page int, limit int) ([]PermissionData, int) {
	var permissions []PermissionData
	var count int
	sysinit.DB.Model(&Permission{}).Limit(limit).Offset((page - 1) * limit).Scan(&permissions)
	sysinit.DB.Model(&Permission{}).Count(&count)
	return permissions, count
}

func CreatePermission() {

}

func DeletePermission() {

}
