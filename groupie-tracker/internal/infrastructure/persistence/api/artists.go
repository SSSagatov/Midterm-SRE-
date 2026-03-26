package api

import (
	"groupie-tracker/internal/domain"
	"groupie-tracker/internal/domain/artist"
	"strconv"
)

const pathArtists = "/api/artists"

type artistJson struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func (r *Repository) GetArtist(id int) (*artist.Artist, error) {
	var artistJson artistJson
	err := r.ParseJSONFromURL(r.apiDomain+pathArtists+"/"+strconv.Itoa(id), &artistJson)
	if err != nil {
		return nil, err
	}
	if artistJson.ID == 0 {
		return nil, domain.ErrIncorrectId
	}
	artist := artist.New(artistJson.ID, artistJson.Image, artistJson.Name, artistJson.Members, artistJson.CreationDate, artistJson.FirstAlbum)
	return artist, nil
}

func (r *Repository) GetArtists() ([]*artist.Artist, error) {
	var artistsJson []artistJson
	err := r.ParseJSONFromURL(r.apiDomain+pathArtists, &artistsJson)
	if err != nil {
		return nil, err
	}
	artists := make([]*artist.Artist, len(artistsJson))
	for i, artistJson := range artistsJson {
		artist := artist.New(artistJson.ID, artistJson.Image, artistJson.Name, artistJson.Members, artistJson.CreationDate, artistJson.FirstAlbum)
		artists[i] = artist
	}
	return artists, nil
}
