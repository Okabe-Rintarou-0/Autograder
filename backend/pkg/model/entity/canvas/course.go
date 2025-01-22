package canvas

const (
	StudentEnrollment  = "StudentEnrollment"
	TaEnrollment       = "TaEnrollment"
	TeacherEnrollment  = "TeacherEnrollment"
	ObserverEnrollment = "ObserverEnrollment"
	DesignerEnrollment = "DesignerEnrollment"
)

type Enrollment struct {
	EnrollmentType  string `json:"type,omitempty"`
	Role            string `json:"role,omitempty"`
	RoleID          int64  `json:"role_id,omitempty"`
	UserID          int64  `json:"user_id,omitempty"`
	EnrollmentState string `json:"enrollment_state,omitempty"`
}

type Term struct {
	ID            int64   `json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	StartAt       *string `json:"start_at,omitempty"`
	EndAt         *string `json:"end_at,omitempty"`
	CreatedAt     *string `json:"created_at,omitempty"`
	WorkflowState string  `json:"workflow_state,omitempty"`
}

type Teacher struct {
	ID          int64  `json:"id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

type Course struct {
	ID                     int64         `json:"id,omitempty"`
	UUID                   string        `json:"uuid,omitempty"`
	Name                   string        `json:"name,omitempty"`
	CourseCode             string        `json:"course_code,omitempty"`
	Enrollments            []*Enrollment `json:"enrollments,omitempty"`
	Teachers               []*Teacher    `json:"teachers,omitempty"`
	AccessRestrictedByDate *bool         `json:"access_restricted_by_date,omitempty"`
	Term                   *Term         `json:"term,omitempty"`
}
