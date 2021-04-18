package db

type HotelFilterParameters struct {
	Limit     int    `form:"limit"`
	Offset    int    `form:"offset"`
	OrderBy   string `form:"order_by"`
	OrderType string `form:"order_type"`

	HotelName  string  `form:"hotel_name"`
	StarsFrom  int     `form:"stars_from"`
	StarsTo    int     `form:"stars_to"`
	RatingFrom float64 `form:"rating_from"`
	RatingTo   float64 `form:"rating_to"`
	PriceFrom  float64 `form:"price_from"`
	PriceTo    float64 `form:"price_to"`

	NearSea  bool   `form:"near_sea"`
	CityName string `form:"city_name"`
}

type EventsFilterParameters struct {
	Limit     int    `form:"limit"`
	Offset    int    `form:"offset"`
	OrderBy   string `form:"order_by"`
	OrderType string `form:"order_type"`

	EventName  string  `form:"event_name"`
	From       string  `form:"from"`
	To         string  `form:"to"`
	RatingFrom float64 `form:"rating_from"`
	RatingTo   float64 `form:"rating_to"`
	PriceFrom  float64 `form:"price_from"`
	PriceTo    float64 `form:"price_to"`

	CityName string `form:"city_name"`
}

type RestaurantFilterParameters struct {
	Limit     int    `form:"limit"`
	Offset    int    `form:"offset"`
	OrderBy   string `form:"order_by"`
	OrderType string `form:"order_type"`

	RestaurantName string  `form:"restaurant_name"`
	RatingFrom     float64 `form:"rating_from"`
	RatingTo       float64 `form:"rating_to"`
	PriceFrom      float64 `form:"price_from"`
	PriceTo        float64 `form:"price_to"`

	ChildMenu   bool `form:"child_menu"`
	SmokingRoom bool `form:"smoking_room"`

	CityName string `form:"city_name"`
}

type TicketFilterParameters struct {
	Limit     int    `form:"limit"`
	Offset    int    `form:"offset"`
	OrderBy   string `form:"order_by"`
	OrderType string `form:"order_type"`

	TransportType string  `form:"transport_type"`
	DateFrom      string  `form:"date_from"`
	DateTo        string  `form:"date_to"`
	PriceFrom     float64 `form:"price_from"`
	PriceTo       float64 `form:"price_to"`

	OrigCityName string `form:"orig_city_name"`
	DestCityName string `form:"dest_city_name"`
}
