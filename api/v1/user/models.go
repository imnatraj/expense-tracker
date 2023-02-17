package user

// request response struct will be written here

type createUserBody struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Admin    bool   `json:"admin"`
}
