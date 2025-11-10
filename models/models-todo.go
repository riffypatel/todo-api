package models

type Todo struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}