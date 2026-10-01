package models

import (
	db "goapi.com/event/DB"
)

var getEvents = []PostEvent{}

func GetAllEvent() ([]PostEvent, error) {
	query := "SELECT*FROM events"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []PostEvent

	for rows.Next() {
		var event PostEvent
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserID, &event.BannerImage, &event.Category, &event.CreatedAt, &event.UpdatedAt)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*PostEvent, error) {
	query := "SELECT *FROM events WHERE id= ?"
	row := db.DB.QueryRow(query, id)

	var event PostEvent

	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserID, &event.BannerImage, &event.Category, &event.CreatedAt, &event.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
