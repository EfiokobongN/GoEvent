package models

import (
	"time"

	db "goapi.com/event/DB"
)

type PostEvent struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"user_id"`
	BannerImage string    `json:"banner_image" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var postEvents = []PostEvent{}

func (event *PostEvent) Save() error {
	query := `INSERT INTO events (title, description, location, date_time, user_id, banner_image, category)
	          VALUES (?, ?, ?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(event.Title, event.Description, event.Location,
		event.DateTime, event.UserID, event.BannerImage, event.Category)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	event.ID = id
	return nil
}

func (event PostEvent) Update() error {
	query := `
	UPDATE events
	SET title = ?, description = ?, location = ?, date_time = ?, banner_image = ?, category = ?
	WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Title, event.Description, event.Location, event.DateTime, event.BannerImage, event.Category, event.ID)

	return err

}

func (event PostEvent) Delete() error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}
