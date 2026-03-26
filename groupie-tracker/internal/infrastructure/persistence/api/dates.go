package api

import (
	"groupie-tracker/internal/domain"
	"groupie-tracker/internal/domain/date"
	"strconv"
)

const pathDates = "/api/dates"

type dateJson struct {
	Dates []string `json:"dates"`
}

func (r *Repository) GetDatesForArtist(artistId int) ([]*date.Date, error) {
	var datesJson dateJson
	err := r.ParseJSONFromURL(r.apiDomain+pathDates+"/"+strconv.Itoa(artistId), &datesJson)
	if err != nil {
		return nil, err
	}
	if len(datesJson.Dates) == 0 {
		return nil, domain.ErrIncorrectId
	}
	var dates []*date.Date
	for _, d := range datesJson.Dates {
		dates = append(dates, date.New(d))
	}
	return dates, nil
}
