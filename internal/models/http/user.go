package models_http

type UserRequest struct {
	Username        string  `json:"username" binding:"required,alphanum,min=5,max=30"`
	DisplayUsername *string `json:"displayusername" binding:"omitempty"`
	Password        string  `json:"password" binding:"required,min=8,max=30"`
}

type UserResponse struct {
	ID              uint64 `json:"id"`
	Username        string `json:"username"`
	DisplayUsername string `json:"displayusername"`
}
