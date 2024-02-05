package domain

type Banner struct {
	ID          int64  `json:"id,omitempty"`
	Cover       string `json:"cover,omitempty"`
	Title       string `json:"title,omitempty"`
	Status      string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	DetailId    int    `json:"detailId,omitempty"`
}
