package application

import (
	"groupie-tracker/internal/domain/artist"
	"groupie-tracker/internal/domain/event"
	"groupie-tracker/internal/domain/date"
	"groupie-tracker/internal/domain/location"
)

type Service struct {
	repo repoInterface
}

type repoInterface interface {
	artist.ArtistRepo
	event.EventRepo
	date.DateRepo
	location.LocationRepo
}

func NewService(repo repoInterface) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetArtist(id int) (*artist.Artist, error) {
	artist, err := s.repo.GetArtist(id)
	if err != nil {
		return nil, err
	}
	events, err := s.repo.GetEvents(id)
	if err != nil {
		return nil, err
	}
	dates, err := s.repo.GetDatesForArtist(id)
	if err != nil {
		return nil, err
	}
	locations, err := s.repo.GetLocationsForArtist(id)
	if err != nil {
		return nil, err
	}
	artist.AddEvents(events)
	artist.AddDates(dates)
	artist.AddLocations(locations)
	return artist, nil
}

func (s *Service) GetArtists() ([]*artist.Artist, error) {
	artists, err := s.repo.GetArtists()
	if err != nil {
		return nil, err
	}
	return artists, nil
}
