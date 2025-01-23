package request

type LoginRequest struct {
	Identifier string `json:"identifier" form:"identifier" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username,omitempty" form:"username" binding:"required"`
	RealName string `json:"real_name,omitempty" form:"real_name" binding:"required"`
	Email    string `json:"email,omitempty" form:"email" binding:"required"`
	Password string `json:"password,omitempty" form:"password" binding:"required"`
}

type ImportCanvasUsersRequest struct {
	CourseID int64 `json:"course_id" form:"course_id" binding:"required"`
}
