package models

import (
	"time"

	db "goapi.com/event/DB"
)

type PostEvent struct {
	ID          int64
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

func (event PostEvent) Save() error {
	query := `INSERT INTO events (title, description, location,datetime, user_id, bannerimage, category) VALUES(?,?,?,?,?,?,?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(event.Title, event.Description, event.Location, event.DateTime, event.UserID, event.BannerImage, event.Category)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	event.ID = id
	return err
}
