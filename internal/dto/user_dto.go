package dto

// UserIdFromUri represents path parameter for user ID
type UserIdFromUri struct {
	ID uint `uri:"id" binding:"required" json:"id" example:"1"` // User ID from path
}

// UserBodyFromCreateRequest represents user creation payload
type UserBodyFromCreateRequest struct {
	Name  string  `form:"name" json:"name" example:"John Doe"`           // User name
	Tel   *string `form:"tel" json:"tel" example:"01012345678"`          // Phone number
	Email *string `form:"email" json:"email" example:"john@example.com"` // Email address
}

// UserBodyFromUpdateRequest represents update payload (path + body)
type UserBodyFromUpdateRequest struct {
	UserIdFromUri
	UserBodyFromCreateRequest
}
