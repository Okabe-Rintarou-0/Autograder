package canvas

type GetAssignmentsRequest struct {
	CourseID int64 `json:"course_id" form:"course_id" binding:"required"`
}

type GetAssignmentSubmissionsRequest struct {
	CourseID     int64 `json:"course_id" form:"course_id" binding:"required"`
	AssignmentID int64 `json:"assignment_id" form:"assignment_id" binding:"required"`
}
