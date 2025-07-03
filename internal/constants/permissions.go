package constants

// PermissionDef defines a static permission's value and description.
type PermissionDef struct {
	Value       string
	Description string
}

var (
	// User permissions
	PermissionUserRead = PermissionDef{
		Value:       "user.read",
		Description: "Read users",
	}

	PermissionUserCreate = PermissionDef{
		Value:       "user.create",
		Description: "Create users",
	}

	PermissionUserUpdate = PermissionDef{
		Value:       "user.update",
		Description: "Update user info",
	}

	PermissionUserDelete = PermissionDef{
		Value:       "user.delete",
		Description: "Delete users",
	}

	PermissionUserRoleUpdate = PermissionDef{
		Value:       "user.role.update",
		Description: "Update user roles",
	}

	// Role permissions
	PermissionRoleRead = PermissionDef{
		Value:       "role.read",
		Description: "Read roles",
	}

	PermissionRoleCreate = PermissionDef{
		Value:       "role.create",
		Description: "Create roles",
	}

	PermissionRoleUpdate = PermissionDef{
		Value:       "role.update",
		Description: "Update roles",
	}

	PermissionRoleDelete = PermissionDef{
		Value:       "role.delete",
		Description: "Delete roles",
	}
)
