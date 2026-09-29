package models

import "time"

type PostEvent struct {
	ID          int
	Title       string    `binding:"required`
	Description string    `binding:"required`
	Location    string    `binding:"required`
	DateTime    time.Time `binding:"required`
	UserID      int
	BannerImage string `binding:"required`
	Category    string `binding:"required`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var postEvents = []PostEvent{}

func (event PostEvent) Save() {
	//TODO: Save event to database

	postEvents = append(postEvents, event)
}
