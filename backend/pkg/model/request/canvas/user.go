package canvas

type GetUsersRequest struct {
	CourseID int64 `json:"course_id" form:"course_id" binding:"required"`
}
