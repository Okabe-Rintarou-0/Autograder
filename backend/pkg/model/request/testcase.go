package request

type Testcase struct {
	Name    string `json:"name"`
	Status  int32  `json:"status"`
	Content string `json:"content"`
}

type BatchUpdateTestcaseRequest struct {
	Data []*Testcase `json:"data"`
}
