package models

import "time"

type Post struct {
	Id          string    `json:"id"`
	PostContent string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UserId      string    `json:"id"`
}
