package db

type UserData struct {
	userId    int
	firstName string
	lastName  string
	photoUrl  string
}

type TourData struct {
	tourId int

	tourName  string
	tourPrice float64

	tourDateFrom string
	tourDateTo   string
}

type CityTourData struct {
	cityTourId int

	countryName   string
	cityName      string
	cityTourPrice float64

	dateFrom string
	dateTo   string

	ticketArrivalId   int
	ticketDepartureId int

	hotelName string
}

type CityTourTicketIdData struct {
	ticketArrivalId   int
	ticketDepartureId int
}

type TicketData struct {
	ticketId      int
	transportType string
	price         float64
	date          string

	origCountryName string
	origCityName    string
	destCountryName string
	destCityName    string

	companyName   string
	companyRating float64
}

type CityTourTicketData struct {
	arrivalTicket   TicketData
	departureTicket TicketData
}

type HotelData struct {
	hotelId   int
	hotelName string

	stars       int
	hotelRating float64
}

type EventData struct {
	eventId   int
	eventName string

	eventStart string
	eventEnd   string
	price      float64
	rating     float64

	maxPersons int
	curPersons int
}

type RestaurantBookingData struct {
	restaurantBookingId int
	restaurantId        int

	bookingTime    string
	restaurantName string

	averagePrice float64
	rating       float64
}
