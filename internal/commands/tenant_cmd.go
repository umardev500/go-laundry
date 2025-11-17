package commands

type CreateTenantCmd struct {
	Name string
}

// --- Update commands ---

type UpdateTenantCmd struct {
	Name *string
}
