package date

type Date struct {
	date string
}

func New(date string) *Date {
	return &Date{
		date: date,
	}
}

func (d *Date) Date() string {
	return d.date
}
