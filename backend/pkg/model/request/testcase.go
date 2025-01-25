package request

type Testcase struct {
	Name   string `json:"name"`
	Status int32  `json:"status"`
}

type BatchUpdateTestcaseRequest struct {
	Data []*Testcase `json:"data"`
}
