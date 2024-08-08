package models

import (
	"api_developement/db"
	"database/sql"
	"fmt"
	"time"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserId      int
}

var eventsDetails []Event = []Event{}

func (e Event) Save() error {

	fmt.Println("working on save", e.Location)
	query := `insert into events(name, description, datetime, location, user_id) 
	values (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("Query executions faild")
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.DateTime, e.Location, e.UserId)

	if err != nil {
		return (err)
	}
	var ID int64
	ID, err = result.LastInsertId()

	fmt.Print("Inserted successfully ", ID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `select * from  events;`
	stmt, err := db.DB.Query(query)

	if err != nil {
		panic(err)
	}

	results := []Event{}
	defer stmt.Close()

	for stmt.Next() {
		var event Event
		stmt.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		results = append(results, event)
	}

	return results, nil
}

func GetEventById(eventId int64) (Event, error) {

	fmt.Println("Fetching data for:", eventId)
	query := `
	select * from events where id= ? `

	result, err := db.DB.Query(query, eventId)
	var eventDetails Event

	if err != nil {
		return eventDetails, err
	}

	defer result.Close()

	if result.Next() {
		err = result.Scan(&eventDetails.ID, &eventDetails.Name, &eventDetails.Description, &eventDetails.Location, &eventDetails.DateTime, &eventDetails.UserId)
		if err != nil {
			return eventDetails, err
		}
	} else {
		return eventDetails, sql.ErrNoRows
	}
	return eventDetails, err
}

func (e Event) UpdateEvent() (Event, error) {

	query := `
	update events set name= ? , description=?, datetime=?, location=? where id=?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println(err)
		return e, err
	}
	_, err = stmt.Exec(e.Name, e.Description, e.DateTime, e.Location, e.ID)
	if err != nil {
		return e, err
	}
	eventsDetails, err := GetEventById(e.ID)

	defer stmt.Close()

	if err != nil {
		return eventsDetails, err
	}
	return eventsDetails, err

}
