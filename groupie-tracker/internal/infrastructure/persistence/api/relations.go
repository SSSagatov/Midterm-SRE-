package api

import (
	"groupie-tracker/internal/domain"
	"groupie-tracker/internal/domain/event"
	"strconv"
)

const pathRelations = "/api/relation"

type eventJson struct {
	DatesLocation map[string][]string `json:"datesLocations"`
}

func (r *Repository) GetEvents(artistId int) ([]*event.Event, error) {
	var eventsJson eventJson
	err := r.ParseJSONFromURL(r.apiDomain+pathRelations+"/"+strconv.Itoa(artistId), &eventsJson)
	if err != nil {
		return nil, err
	}
	if eventsJson.DatesLocation == nil {
		return nil, domain.ErrIncorrectId
	}
	var events []*event.Event
	for location, dates := range eventsJson.DatesLocation {
		events = append(events, event.New(location, dates))
	}
	return events, nil
}
