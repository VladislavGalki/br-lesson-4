package models

type Task struct {
	UserId      string `json:"user_id"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
