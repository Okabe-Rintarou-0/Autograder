package entity

import "autograder/pkg/model/dbm"

type User struct {
	UserID   uint   `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	RealName string `json:"real_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     int32  `json:"role,omitempty"`
}

func (u *User) IsAdmin() bool {
	return u.Role == dbm.Administrator
}
