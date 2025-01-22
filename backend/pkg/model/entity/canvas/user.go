package canvas

type User struct {
	Id           int64   `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	CreatedAt    string  `json:"created_at,omitempty"`
	SortableName string  `json:"sortable_name,omitempty"`
	ShortName    string  `json:"short_name,omitempty"`
	LoginId      string  `json:"login_id,omitempty"`
	Email        *string `json:"email,omitempty"`
}
