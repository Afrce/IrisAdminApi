package models

import (
	"IrisAdminApi/sysinit"
)

type UserHasRole struct {
	ID     uint
	UserId uint
	RoleId uint
}

func (UserHasRole) TableName() string {
	return "user_has_role"
}

func SelectUserRoles(userId uint) []uint {
	var roles []uint
	sysinit.DB.Table(UserHasRole{}.TableName()).Where("user_id = ?", userId).Pluck("role_id", &roles)
	return roles
}
