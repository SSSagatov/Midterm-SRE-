package event

type EventRepo interface {
	GetEvents(artistId int) ([]*Event, error)
}
