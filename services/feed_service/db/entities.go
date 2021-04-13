package db

type Hotel struct {
	HotelId     int     `json:"hotel_id"`
	HotelName   string  `json:"hotel_name"`
	Stars       int     `json:"stars"`
	HotelRating float64 `json:"hotel_rating"`

	CountryName string `json:"country_name"`
	CityName    string `json:"city_name"`
}

type Event struct {
	EventId   int    `json:"event_id"`
	EventName string `json:"event_name"`

	EventStart string  `json:"event_start"`
	EventEnd   string  `json:"event_end"`
	Rating     float64 `json:"rating"`

	MaxPersons int `json:"max_persons"`
	CurPersons int `json:"cur_persons"`

	CountryName string `json:"country_name"`
	CityName    string `json:"city_name"`
}

type Restaurant struct {
	RestaurantId   int     `json:"restaurant_id"`
	RestaurantName string  `json:"restaurant_name"`
	Rating         float64 `json:"rating"`

	CountryName string `json:"country_name"`
	CityName    string `json:"city_name"`
}

type Ticket struct {
	TicketId      int     `json:"ticket_id"`
	TransportType string  `json:"transport_type"`
	Price         float64 `json:"price"`
	Date          string  `json:"date"`

	OrigCountryName string `json:"orig_country_name"`
	OrigCityName    string `json:"orig_city_name"`
	DestCountryName string `json:"dest_country_name"`
	DestCityName    string `json:"dest_city_name"`
}
