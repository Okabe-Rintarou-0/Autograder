package response

import "time"

type SubmitAppResponse struct {
	*BaseResp
}

type AppRunTask struct {
	UUID        string    `json:"uuid"`
	UserID      uint      `json:"user_id"`
	Username    string    `json:"username"`
	Email       string    `json:"user_email"`
	RealName    string    `json:"real_name"`
	Status      int32     `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	Pass        int32     `json:"pass"`
	Total       int32     `json:"total"`
	TestResults *string   `json:"test_results"`
}

type ListAppTasksResponse struct {
	*BaseResp
	Total int64         `json:"total"`
	Data  []*AppRunTask `json:"data"`
}
