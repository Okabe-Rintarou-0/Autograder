package response

import "autograder/pkg/model/entity"

type LoginResponse struct {
	*BaseResp

	Token string `json:"token"`
}

type GetMeResponse struct {
	entity.User
}
