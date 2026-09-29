package models

import "time"

type PostEvent struct {
	ID          int
	Title       string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
	BannerImage string
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func saveEvent(event PostEvent) {
	//TODO: Save event to database
}
