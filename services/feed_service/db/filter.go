package db

type HotelFilterParameters struct {
	limit     int
	offset    int
	orderBy   string
	orderType string

	hotelName  string
	starsFrom  int
	starsTo    int
	ratingFrom float64
	ratingTo   float64
	priceFrom  float64
	priceTo    float64

	nearSea  bool
	cityName string
}

type EventsFilterParameters struct {
	limit     int
	offset    int
	orderBy   string
	orderType string

	eventName  string
	from       string
	to         string
	ratingFrom float64
	ratingTo   float64
	priceFrom  float64
	priceTo    float64

	cityName string
}

type RestaurantFilterParameters struct {
	limit     int
	offset    int
	orderBy   string
	orderType string

	restaurantName string
	ratingFrom     float64
	ratingTo       float64
	priceFrom      float64
	priceTo        float64

	childMenu   bool
	smokingRoom bool

	cityName string
}

type TicketFilterParameters struct {
	limit     int
	offset    int
	orderBy   string
	orderType string

	transportType string
	dateFrom      string
	dateTo        string
	priceFrom     float64
	priceTo       float64

	origCityName string
	destCityName string
}
