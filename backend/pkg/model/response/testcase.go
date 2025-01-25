package response

type Testcase struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status int32  `json:"status"`
}

type BatchUpdateTestcaseResponse struct {
	*BaseResp
}
