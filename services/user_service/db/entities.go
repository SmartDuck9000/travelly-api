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

type CityTourTicketID struct {
	ticketArrivalId   int
	ticketDepartureId int
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
	hotelId   int
	hotelName string

	stars       int
	hotelRating float64
}

type Event struct {
	eventId   int
	eventName string

	eventStart string
	eventEnd   string
	price      float64
	rating     float64

	maxPersons int
	curPersons int
}

type RestaurantBooking struct {
	restaurantBookingId int
	restaurantId        int

	bookingTime    string
	restaurantName string

	averagePrice float64
	rating       float64
}
