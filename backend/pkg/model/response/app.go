package response

import "time"

type SubmitAppResponse struct {
	*BaseResp
}

type UserProfile struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	RealName string `json:"real_name"`
	Email    string `json:"email"`
	Role     int32  `json:"role"`
}

type AppRunTask struct {
	UUID        string       `json:"uuid"`
	User        *UserProfile `json:"user"`
	Operator    *UserProfile `json:"operator"`
	Status      int32        `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	Pass        int32        `json:"pass"`
	Total       int32        `json:"total"`
	TestResults *string      `json:"test_results"`
}

type ListAppTasksResponse struct {
	*BaseResp
	Total int64         `json:"total"`
	Data  []*AppRunTask `json:"data"`
}
