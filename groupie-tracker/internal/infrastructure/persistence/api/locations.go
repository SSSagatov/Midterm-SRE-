package api

import (
	domainLocation "groupie-tracker/internal/domain/location"
	"strconv"
)

const pathLocations = "/api/locations"

type locationJson struct {
	Locations []string `json:"locations"`
}

func (r *Repository) GetLocationsForArtist(artistId int) ([]*domainLocation.Location, error) {
	var locationJson locationJson
	err := r.ParseJSONFromURL(r.apiDomain+pathLocations+"/"+strconv.Itoa(artistId), &locationJson)
	if err != nil {
		return nil, err
	}
	locations := make([]*domainLocation.Location, len(locationJson.Locations))
	for i, location := range locationJson.Locations {
		locations[i] = domainLocation.New(location)
	}
	return locations, nil
}
