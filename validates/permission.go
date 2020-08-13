package validates

type CreatePermissionRequest struct {
	Name string `json:"name" validate:"required"  comment:"权限名称"`
}
