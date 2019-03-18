package requests

// we will separate each request so it is easy to adjust later

// CreateUserRequest post
type CreateUserRequest struct {
	Name        string `json:"name" binding:"alpha,required"`
	Email       string `json:"email" binding:"email,required"`
	PhoneNumber string `json:"phone" binding:"numeric,required"`
}

// UpdateUserRequest put
type UpdateUserRequest struct {
	Name        string `json:"name" binding:"alpha,required"`
	Email       string `json:"email" binding:"email,required"`
	PhoneNumber string `json:"phone" binding:"numeric,required"`
}
