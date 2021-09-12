package db

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	PhotoUrl string `json:"photo_url"`
}

type Tour struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`

	TourName  string  `json:"tour_name"`
	TourPrice float64 `json:"tour_price"`

	TourDateFrom string `json:"tour_date_from"`
	TourDateTo   string `json:"tour_date_to"`
}

type CityTour struct {
	Id     int `json:"id"`
	TourId int `json:"tour_id"`
	CityId int `json:"city_id"`

	CityTourPrice float64 `json:"city_tour_price"`

	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`

	TicketArrivalId   int `json:"ticket_arrival_id"`
	TicketDepartureId int `json:"ticket_departure_id"`

	HotelId int `json:"hotel_id"`
}

type CityToursEvent struct {
	CtId    int `json:"ct_id"`
	EventId int `json:"event_id"`
}

type RestaurantBookingDTO struct {
	CtId         int    `json:"id"`
	RestaurantId int    `json:"restaurant_id"`
	BookingTime  string `json:"booking_time"`
}

type RestaurantBooking struct {
	Id           int    `json:"id"`
	RestaurantId int    `json:"restaurant_id"`
	BookingTime  string `json:"booking_time"`
}

type CityToursRestBooking struct {
	CtId int `json:"ct_id"`
	RbId int `json:"rb_id"`
}
