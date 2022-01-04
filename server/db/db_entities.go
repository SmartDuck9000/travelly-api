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

type HotelFeedItem struct {
	HotelId     int     `json:"hotel_id"`
	HotelName   string  `json:"hotel_name"`
	Stars       int     `json:"stars"`
	HotelRating float64 `json:"hotel_rating"`

	CountryName string `json:"country_name"`
	CityName    string `json:"city_name"`
}

type EventFeedItem struct {
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

type RestaurantFeedItem struct {
	RestaurantId   int     `json:"restaurant_id"`
	RestaurantName string  `json:"restaurant_name"`
	Rating         float64 `json:"rating"`

	CountryName string `json:"country_name"`
	CityName    string `json:"city_name"`
}

type TicketFeedItem struct {
	TicketId      int    `json:"ticket_id"`
	TransportType string `json:"transport_type"`
	Price         string `json:"price"`
	Date          string `json:"date"`

	OrigCountryName string `json:"orig_country_name"`
	OrigCityName    string `json:"orig_city_name"`
	DestCountryName string `json:"dest_country_name"`
	DestCityName    string `json:"dest_city_name"`
}

type City struct {
	CityId   int    `json:"city_id"`
	CityName string `json:"city_name"`
}

type Hotel struct {
	HotelId          int    `json:"hotel_id"`
	HotelName        string `json:"hotel_name"`
	HotelDescription string `json:"hotel_description"`
	HotelAddr        string `json:"hotel_addr"`

	Stars       int     `json:"stars"`
	HotelRating float64 `json:"hotel_rating"`
	AvgPrice    string  `json:"average_price"`

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
	Price      string  `json:"price"`
	Rating     float64 `json:"rating"`

	MaxPersons int      `json:"max_persons"`
	CurPersons int      `json:"cur_persons"`
	Languages  []string `json:"languages"`
}

type Restaurant struct {
	RestaurantId int `json:"restaurant_id"`

	RestName        string `json:"restaurant_name"`
	RestDescription string `json:"restaurant_description"`
	RestAddr        string `json:"restaurant_addr"`

	AvgPrice   string  `json:"average_price"`
	RestRating float64 `json:"rating"`

	ChildMenu   bool `json:"child_menu"`
	SmokingRoom bool `json:"smoking_room"`

	CountryName string `json:"country_name"`
	CityName    string `json:"city_name"`
}

type Ticket struct {
	TicketId int `json:"ticket_id"`

	CompanyName   string  `json:"company_name"`
	CompanyRating float64 `json:"company_rating"`

	OrigStationName string `json:"orig_station_name"`
	OrigStationAddr string `json:"orig_station_addr"`
	OrigCountryName string `json:"orig_country_name"`
	OrigCityName    string `json:"orig_city_name"`

	DestStationName string `json:"dest_station_name"`
	DestStationAddr string `json:"dest_station_addr"`
	DestCityName    string `json:"dest_city_name"`
	DestCountryName string `json:"dest_country_name"`

	TransportType string `json:"transport_type"`
	Price         string `json:"price"`
	TicketDate    string `json:"ticket_date"`
}
