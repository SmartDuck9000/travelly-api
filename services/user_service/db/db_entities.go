package db

type UserEntity struct {
	Id       int
	Email    string
	Password string

	FirstName string
	LastName  string

	PhotoUrl string
}

type TourEntity struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`

	TourName  string  `json:"tour_name"`
	TourPrice float64 `json:"tour_price"`

	TourDateFrom string `json:"tour_date_from"`
	TourDateTo   string `json:"tour_date_to"`
}

type CityTourEntity struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	CityId int `json:"city_id"`

	CityTourPrice float64 `json:"city_tour_price"`

	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`

	TicketArrivalId   int `json:"ticket_arrival_id"`
	TicketDepartureId int `json:"ticket_departure_id"`

	HotelId int `json:"hotel_id"`
}

type RestaurantBookingEntity struct {
	Id           int    `json:"id"`
	RestaurantId int    `json:"restaurant_id"`
	BookingTime  string `json:"booking_time"`
}
