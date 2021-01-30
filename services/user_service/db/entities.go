package db

type User struct {
	userId    int
	firstName string
	lastName  string
	photoUrl  string
}

type Tour struct {
	tourId int
	userId int

	tourName  string
	tourPrice float64

	tourDateFrom string
	tourDateTo   string
}

type CityTour struct {
	cityTourId int
	tourId     int

	countryName   string
	cityName      string
	cityTourPrice float64

	dateFrom string
	dateTo   string

	ticketArrivalId   int
	ticketDepartureId int

	hotelName string
}

type Ticket struct {
	ticketId      int
	transportType string
	price         float64
	date          string

	origStationName string
	origStationAddr string
	dstStationName  string
	dstStationAddr  string

	companyName   string
	companyRating float64
}

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

type RestaurantBooking struct {
	restaurantId int
	bookingTime  string

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
