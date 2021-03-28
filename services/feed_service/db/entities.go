package db

type Hotel struct {
	hotelId     int
	hotelName   string
	stars       int
	hotelRating float64

	countryName string
	cityName    string
}

type Event struct {
	eventId   int
	eventName string

	eventStart string
	eventEnd   string
	rating     float64

	maxPersons int
	curPersons int

	countryName string
	cityName    string
}

type Restaurant struct {
	restaurantId   int
	restaurantName string
	rating         float64

	countryName string
	cityName    string
}
