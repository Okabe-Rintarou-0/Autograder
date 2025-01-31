package request

type Testcase struct {
	ID      uint   `json:"id" form:"id" binding:"required"`
	Name    string `json:"name" form:"name" binding:"required"`
	Path    string `json:"path" form:"path" binding:"required"`
	Status  int32  `json:"status" form:"status" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
}

type BatchUpdateTestcaseRequest struct {
	Data []*Testcase `json:"data" form:"data" binding:"required"`
}
