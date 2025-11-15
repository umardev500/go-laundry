package commands

type CreateProfileCmd struct {
	Name string
}

type CreateUserCmd struct {
	Email    string
	Password string
	Profile  *CreateProfileCmd
}

// --- Update commands ---

type UpdateProfileCmd struct {
	Name *string
}

type UpdateUserCmd struct {
	Email    *string
	Password *string
}
