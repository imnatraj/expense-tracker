package hello

// request response struct will be written here

type createHelloBody struct {
	Name string `json:"name" validate:"required"`
}


