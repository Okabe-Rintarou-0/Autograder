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

type Assignment struct {
	ID                      int64       `json:"id"`
	Description             *string     `json:"description,omitempty"`
	DueAt                   *string     `json:"due_at,omitempty"`
	UnlockAt                *string     `json:"unlock_at,omitempty"`
	LockAt                  *string     `json:"lock_at,omitempty"`
	PointsPossible          *float64    `json:"points_possible,omitempty"`
	CourseID                int64       `json:"course_id"`
	Name                    string      `json:"name"`
	NeedsGradingCount       *int32      `json:"needs_grading_count,omitempty"`
	HtmlURL                 string      `json:"html_url"`
	SubmissionTypes         []string    `json:"submission_types"`
	AllowedExtensions       []string    `json:"allowed_extensions,omitempty"`
	HasSubmittedSubmissions bool        `json:"has_submitted_submissions"`
	Published               bool        `json:"published"`
	SubmissionsDownloadURL  string      `json:"submissions_download_url,omitempty"`
	Submission              *Submission `json:"submission,omitempty"`
}

type Submission struct {
	ID           int64         `json:"id"`
	SubmittedAt  *string       `json:"submitted_at,omitempty"`
	Grade        *string       `json:"grade,omitempty"`
	AssignmentID int64         `json:"assignment_id"`
	UserID       int64         `json:"user_id"`
	Late         bool          `json:"late"`
	Attachments  []*Attachment `json:"attachments"`
}

type Attachment struct {
	ID          int64  `json:"id"`
	UUID        string `json:"uuid"`
	FolderID    *int64 `json:"folder_id,omitempty"`
	DisplayName string `json:"display_name"`
	Filename    string `json:"filename"`
	URL         string `json:"url,omitempty"`
	Size        int64  `json:"size"`
	Locked      bool   `json:"locked"`
	MimeClass   string `json:"mime_class,omitempty"`
	ContentType string `json:"content-type,omitempty"`
}
