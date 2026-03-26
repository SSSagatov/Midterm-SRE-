package artist

import (
	"groupie-tracker/internal/domain/date"
	"groupie-tracker/internal/domain/event"
	"groupie-tracker/internal/domain/location"
)

type Artist struct {
	id               int
	imageLink        string
	name             string
	members          []string
	creationDateYear int
	firstAlbumDate   string
	events           []*event.Event
	locations        []*location.Location
	dates            []*date.Date
}

func New(id int, imageLink string, name string, members []string, creationDateYear int, firstAlbumDate string) *Artist {
	return &Artist{
		id:               id,
		imageLink:        imageLink,
		name:             name,
		members:          members,
		creationDateYear: creationDateYear,
		firstAlbumDate:   firstAlbumDate,
	}
}

func (a *Artist) AddEvents(events []*event.Event) {
	a.events = events
}

func (a *Artist) AddLocations(locations []*location.Location) {
	a.locations = locations
}

func (a *Artist) AddDates(dates []*date.Date) {
	a.dates = dates
}

func (a *Artist) ID() int {
	return a.id
}

func (a *Artist) ImageLink() string {
	return a.imageLink
}

func (a *Artist) Name() string {
	return a.name
}

func (a *Artist) Members() []string {
	return a.members
}

func (a *Artist) CreationDateYear() int {
	return a.creationDateYear
}

func (a *Artist) FirstAlbumDate() string {
	return a.firstAlbumDate
}

func (a *Artist) Events() []*event.Event {
	return a.events
}

func (a *Artist) Locations() []*location.Location {
	return a.locations
}

func (a *Artist) Dates() []*date.Date {
	return a.dates
}
