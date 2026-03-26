package artist

type ArtistRepo interface {
	GetArtist(id int) (*Artist, error)
	GetArtists() ([]*Artist, error)
}
