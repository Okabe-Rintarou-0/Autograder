package response

type Testcase struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Path    string `json:"path"`
	Status  int32  `json:"status"`
	Content string `json:"content"`
}

type BatchUpdateTestcaseResponse struct {
	*BaseResp
}
