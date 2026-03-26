package location

type LocationRepo interface {
	GetLocationsForArtist(artistId int) ([]*Location, error)
}