package response

import "time"

type SubmitAppResponse struct {
	*BaseResp
}

type AppRunTask struct {
	UUID      string    `json:"uuid"`
	UserID    uint      `json:"user_id"`
	Status    int32     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type ListAppTasksResponse struct {
	*BaseResp
	Total int64         `json:"total"`
	Data  []*AppRunTask `json:"data"`
}
