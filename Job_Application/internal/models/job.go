package models

import "time"

type Job struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Salary      string    `json:"salary"`
	Experience  int       `json:"experience"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      int       `json:"user_id"`
}
