package db

type Hotel struct {
	HotelId          int    `json:"hotel_id"`
	HotelName        string `json:"hotel_name"`
	HotelDescription string `json:"hotel_description"`
	HotelAddr        string `json:"hotel_addr"`

	Stars       int     `json:"stars"`
	HotelRating float64 `json:"hotel_rating"`
	AvgPrice    float64 `json:"average_price"`

	NearSea bool `json:"near_sea"`

	CountryName string `json:"country_name"`
	CityName    string `json:"city_name"`
}

type Event struct {
	EventId          int    `json:"event_id"`
	EventName        string `json:"event_name"`
	EventDescription string `json:"event_description"`
	EventAddr        string `json:"event_addr"`

	CountryName string `json:"country_name"`
	CityName    string `json:"city_name"`

	EventStart string  `json:"event_start"`
	EventEnd   string  `json:"event_end"`
	Price      float64 `json:"price"`
	Rating     float64 `json:"rating"`

	MaxPersons int      `json:"max_persons"`
	CurPersons int      `json:"cur_persons"`
	Languages  []string `json:"languages"`
}

type Restaurant struct {
	RestaurantId int `json:"restaurant_id"`

	RestaurantName        string `json:"restaurant_name"`
	RestaurantDescription string `json:"restaurant_description"`
	RestaurantAddr        string `json:"restaurant_addr"`

	AvgPrice float64 `json:"average_price"`
	Rating   float64 `json:"rating"`

	ChildMenu   bool `json:"child_menu"`
	SmokingRoom bool `json:"smoking_room"`

	CountryName string `json:"country_name"`
	CityName    string `json:"city_name"`
}

type Ticket struct {
	TicketId int `json:"ticket_id"`

	CompanyName   string `json:"company_name"`
	CompanyRating string `json:"company_rating"`

	OrigStationName string `json:"orig_station_name"`
	OrigStationAddr string `json:"orig_station_addr"`
	OrigCountryName string `json:"orig_country_name"`
	OrigCityName    string `json:"orig_city_name"`

	DestStationName string `json:"dest_station_name"`
	DestStationAddr string `json:"dest_station_addr"`
	DestCityName    string `json:"dest_city_name"`
	DestCountryName string `json:"dest_country_name"`

	TransportType string  `json:"transport_type"`
	Price         float64 `json:"price"`
	TicketDate    string  `json:"ticket_date"`
}
