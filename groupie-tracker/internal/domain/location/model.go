package location

type Location struct {
	location string
}

func New(location string) *Location {
	return &Location{
		location: location,
	}
}

func (l *Location) Location() string {
	return l.location
}
