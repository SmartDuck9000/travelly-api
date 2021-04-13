package db

type HotelFilterParameters struct {
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	OrderBy   string `json:"order_by"`
	OrderType string `json:"order_type"`

	HotelName  string  `json:"hotel_name"`
	StarsFrom  int     `json:"stars_from"`
	StarsTo    int     `json:"stars_to"`
	RatingFrom float64 `json:"rating_from"`
	RatingTo   float64 `json:"rating_to"`
	PriceFrom  float64 `json:"price_from"`
	PriceTo    float64 `json:"price_to"`

	NearSea  bool   `json:"near_sea"`
	CityName string `json:"city_name"`
}

type EventsFilterParameters struct {
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	OrderBy   string `json:"order_by"`
	OrderType string `json:"order_type"`

	EventName  string  `json:"event_name"`
	From       string  `json:"from"`
	To         string  `json:"to"`
	RatingFrom float64 `json:"rating_from"`
	RatingTo   float64 `json:"rating_to"`
	PriceFrom  float64 `json:"price_from"`
	PriceTo    float64 `json:"price_to"`

	CityName string `json:"city_name"`
}

type RestaurantFilterParameters struct {
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	OrderBy   string `json:"order_by"`
	OrderType string `json:"order_type"`

	RestaurantName string  `json:"restaurant_name"`
	RatingFrom     float64 `json:"rating_from"`
	RatingTo       float64 `json:"rating_to"`
	PriceFrom      float64 `json:"price_from"`
	PriceTo        float64 `json:"price_to"`

	ChildMenu   bool `json:"child_menu"`
	SmokingRoom bool `json:"smoking_room"`

	CityName string `json:"city_name"`
}

type TicketFilterParameters struct {
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	OrderBy   string `json:"order_by"`
	OrderType string `json:"order_type"`

	TransportType string  `json:"transport_type"`
	DateFrom      string  `json:"date_from"`
	DateTo        string  `json:"date_to"`
	PriceFrom     float64 `json:"price_from"`
	PriceTo       float64 `json:"price_to"`

	OrigCityName string `json:"orig_city_name"`
	DestCityName string `json:"dest_city_name"`
}
