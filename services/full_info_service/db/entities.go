package db

type Hotel struct {
	hotelId          int
	hotelName        string
	hotelDescription string
	hotelAddr        string

	stars        int
	hotelRating  float64
	averagePrice float64

	nearSea bool

	countryName string
	cityName    string
}

type Event struct {
	eventId          int
	eventName        string
	eventDescription string
	eventAddr        string

	countryName string
	cityName    string

	eventStart string
	eventEnd   string
	price      float64
	rating     float64

	maxPersons int
	curPersons int
	languages  []string
}

type Restaurant struct {
	restaurantId int

	restaurantName        string
	restaurantDescription string
	restaurantAddr        string

	averagePrice float64
	rating       float64

	childMenu   bool
	smokingRoom bool

	countryName string
	cityName    string
}

type Ticket struct {
	ticketId int

	companyName   string
	companyRating string

	origStationName string
	origStationAddr string
	origCountryName string
	origCityName    string

	destStationName string
	destStationAddr string
	destCityName    string
	destCountryName string

	transportType string
	price         float64
	ticketDate    string
}
