package db

type UserEntity struct {
	id       int
	email    string
	password string

	firstName string
	lastName  string

	photoUrl string
}

type TourEntity struct {
	id     int
	userId int

	tourName  string
	tourPrice float64

	tourDateFrom string
	tourDateTo   string
}

type CityTourEntity struct {
	id     int
	userId int
	cityId int

	cityTourPrice float64

	dateFrom string
	dateTo   string

	ticketArrivalId   int
	ticketDepartureId int

	hotelId int
}

type RestaurantBookingEntity struct {
	id           int
	restaurantId int
	bookingTime  string
}
