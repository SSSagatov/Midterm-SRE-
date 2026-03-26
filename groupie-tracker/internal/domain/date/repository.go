package date

type DateRepo interface {
	GetDatesForArtist(artistId int) ([]*Date, error)
}