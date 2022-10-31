package models

type Predict struct {
	ID       int64  `json:"id"`
	Sentence string `json:"sentence"`
	Intent   string `json:"intent"`
	UserID   int64  `json:"userID"`
}
