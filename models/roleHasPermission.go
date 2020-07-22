package models

import "IrisAdminApi/sysinit"

type RoleHasPermission struct {
	ID           uint
	RoleID       uint
	PermissionID uint
}

func (RoleHasPermission) TableName() string {
	return "role_has_permission"
}

func CheckRoleHasPermission(roles []uint, permission uint) bool {
	roleHasPermission := new(RoleHasPermission)
	if err := sysinit.DB.Where("role_id in (?)", roles).Where("permission_id = ?", permission).Find(&roleHasPermission).Error; err != nil {
		return false
	} else {
		return true
	}

}
