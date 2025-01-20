package response

type LoginResponse struct {
	*BaseResp

	Token string `json:"token"`
}

type GetMeResponse struct {
	UserID   uint   `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}
