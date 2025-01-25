package response

type Testcase struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Status  int32  `json:"status"`
	Content string `json:"content"`
}

type BatchUpdateTestcaseResponse struct {
	*BaseResp
}
