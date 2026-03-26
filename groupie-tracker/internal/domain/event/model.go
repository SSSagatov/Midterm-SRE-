package event

type Event struct {
	dates    []string
	location string
}

func New(location string, dates []string) *Event {
	return &Event{
		location: location,
		dates:    dates,
	}
}

func (e *Event) Dates() []string {
	return e.dates
}

func (e *Event) Location() string {
	return e.location
}
