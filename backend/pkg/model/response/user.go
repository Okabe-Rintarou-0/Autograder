package response

import "time"

type LoginResponse struct {
	*BaseResp

	Token string `json:"token"`
}

type RegisterResponse struct {
	*BaseResp
}

type ImportCanvasUsers struct {
	*BaseResp
}

type GetMeResponse struct {
	UserID   uint   `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

type User struct {
	ID        uint      `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	RealName  string    `json:"real_name,omitempty"`
	Role      int32     `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type ListUsersResponse struct {
	*BaseResp
	Total int64   `json:"total"`
	Data  []*User `json:"data"`
}
