package entity

type User struct {
	UserID   uint   `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}
