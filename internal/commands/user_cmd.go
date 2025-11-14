package commands

type CreateProfileCmd struct {
	Name string
}

type CreateUserCmd struct {
	Email    string
	Password string
	Profile  *CreateProfileCmd
}
