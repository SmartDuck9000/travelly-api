package db

type UserData struct {
	Id        int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PhotoUrl  string `json:"photo_url"`
}

type TourData struct {
	Id int `json:"tour_id"`

	TourName  string `json:"tour_name"`
	TourPrice string `json:"tour_price"`

	TourDateFrom string `json:"tour_date_from"`
	TourDateTo   string `json:"tour_date_to"`
}

type CityTourData struct {
	Id int `json:"city_tour_id"`

	CountryName   string `json:"country_name"`
	CityName      string `json:"city_name"`
	CityTourPrice string `json:"city_tour_price"`

	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`

	TicketArrivalId   int `json:"ticket_arrival_id"`
	TicketDepartureId int `json:"ticket_departure_id"`

	HotelId int `json:"hotel_id"`
}

type CityTourTicketIdData struct {
	TicketArrivalId   int `json:"ticket_arrival_id"`
	TicketDepartureId int `json:"ticket_departure_id"`
}

type TicketData struct {
	Id            int    `json:"ticket_id"`
	TransportType string `json:"transport_type"`
	Price         string `json:"price"`
	Date          string `json:"date"`

	OrigCountryName string `json:"orig_country_name"`
	OrigCityName    string `json:"orig_city_name"`
	DestCountryName string `json:"dest_country_name"`
	DestCityName    string `json:"dest_city_name"`

	CompanyName   string  `json:"company_name"`
	CompanyRating float64 `json:"company_rating"`
}

type CityTourTicketData struct {
	ArrivalTicket   TicketData `json:"arrival_ticket"`
	DepartureTicket TicketData `json:"departure_ticket"`
}

type HotelData struct {
	Id        int    `json:"hotel_id"`
	HotelName string `json:"hotel_name"`

	Stars       int     `json:"stars"`
	HotelRating float64 `json:"hotel_rating"`
}

type EventData struct {
	Id        int    `json:"event_id"`
	EventName string `json:"event_name"`

	EventStart string  `json:"event_start"`
	EventEnd   string  `json:"event_end"`
	Price      string  `json:"price"`
	Rating     float64 `json:"rating"`

	MaxPersons int `json:"max_persons"`
	CurPersons int `json:"cur_persons"`
}

type RestaurantBookingData struct {
	Id           int `json:"restaurant_booking_id"`
	RestaurantId int `json:"restaurant_id"`

	BookingTime    string `json:"booking_time"`
	RestaurantName string `json:"restaurant_name"`

	AveragePrice string  `json:"average_price"`
	Rating       float64 `json:"rating"`
}
