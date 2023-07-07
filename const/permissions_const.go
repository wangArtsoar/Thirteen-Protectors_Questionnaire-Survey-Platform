package _const

// permissions
const (
	AdminCreate = "admin::create"
	AdminUpdate = "admin::update"
	AdminSelect = "admin::select"
	AdminDelete = "admin::delete"

	ManageCreate = "manage::create"
	ManageUpdate = "manage::update"
	ManageDelete = "manage::delete"
	ManageSelect = "manage::select"
)

// Super return SuperAdmin permissions
func Super() any {
	return map[string]any{
		"SUPER": []string{
			AdminCreate, AdminSelect, AdminUpdate, AdminDelete,
			ManageCreate, ManageSelect, ManageUpdate, ManageDelete},
	}
}

// Admin return admin permissions
func Admin() (string, any) {
	return "ADMIN", []string{
		AdminCreate, AdminSelect, AdminUpdate, AdminDelete,
	}
}

// Manage return manage permissions
func Manage() (string, any) {
	return "MANAGE", []string{
		ManageCreate, ManageSelect, ManageUpdate, ManageDelete,
	}
}

// User return user permissions
func User() any {
	return map[string]any{"USER": []string{}}
}

// Customize return customize permissions
// 自定义角色 (Super & Admin & Manage)
func Customize(roleName string, permissions ...string) any {
	return map[string]any{
		roleName: permissions,
	}
}
