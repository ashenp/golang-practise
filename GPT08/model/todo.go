package model

type Todo struct {
	ID               int64  `json:"id"`
	Title            string `json:"title"`
	Description      string `json:"description,omitempty"`
	IsCompleted      bool   `json:"isCompleted"`
	DueTimestamp     int    `json:"dueTimestamp,omitempty"`
	CreatedTimestamp int    `json:"createdTimestamp,omitempty"`
}
