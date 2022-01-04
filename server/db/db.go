package db

import (
	"fmt"
	"github.com/SmartDuck9000/travelly-api/server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Repository interface {
	Open() error
	configureConnectionPools() error

	CreateUser(user *User) error
	GetUserByEmail(email string) *User

	GetUser(userId int) *UserData
	GetUserPassword(userId int) (string, error)
	GetTours(userId int) []TourData
	GetCityTours(tourId int) []CityTourData

	GetUserEvents(cityTourId int) []EventData
	GetRestaurantBookings(cityTourId int) []RestaurantBookingData
	GetUserTickets(cityTourId int) *CityTourTicketData
	GetUserHotel(cityTourId int) *HotelData

	CreateTour(tour *Tour) error
	CreateCityTour(cityTour *CityTour) error
	CreateCityTourEvent(cityTourEvent *CityToursEvent) error
	CreateRestaurantBooking(restaurantBooking *RestaurantBookingDTO) error

	UpdateUser(user *User) error
	UpdateTour(tour *Tour) error
	UpdateCityTour(cityTour *CityTour) error
	UpdateRestaurantBooking(restaurantBooking *RestaurantBooking) error

	DeleteUser(userId int) error
	DeleteTour(tourId int) error
	DeleteCityTour(cityTourId int) error
	DeleteRestaurantBooking(restaurantBookingId int) error

	GetCities() ([]City, error)
	GetHotel(id int) (*Hotel, error)
	GetEvent(id int) (*Event, error)
	GetRestaurant(id int) (*Restaurant, error)
	GetTicket(id int) (*Ticket, error)

	GetHotels(filter HotelFilterParameters) ([]HotelFeedItem, error)
	GetEvents(filter EventsFilterParameters) ([]EventFeedItem, error)
	GetRestaurants(filter RestaurantFilterParameters) ([]RestaurantFeedItem, error)
	GetTickets(filter TicketFilterParameters) ([]TicketFeedItem, error)
}

type PostgresRepository struct {
	url             string
	maxIdleConn     int
	maxOpenConn     int
	connMaxLifetime time.Duration
	conn            *gorm.DB
}

func CreateRepository(conf config.DbConfig) Repository {
	return &PostgresRepository{
		url:             conf.URL,
		maxIdleConn:     conf.MaxIdleConn,     // maximum number of connections in the idle connection pool
		maxOpenConn:     conf.MaxOpenConn,     // maximum number of open connections to the database
		connMaxLifetime: conf.ConnMaxLifetime, // maximum amount of time a connection may be reused
		conn:            nil,
	}
}

func (db *PostgresRepository) Open() error {
	var err error
	db.conn, err = gorm.Open(postgres.Open(db.url), &gorm.Config{})
	if err == nil {
		err = db.configureConnectionPools()
	}
	return err
}

func (db PostgresRepository) configureConnectionPools() error {
	sqlDB, err := db.conn.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(db.maxIdleConn)
	sqlDB.SetMaxOpenConns(db.maxOpenConn)
	sqlDB.SetConnMaxLifetime(db.connMaxLifetime)

	return nil
}

func (db PostgresRepository) CreateUser(user *User) error {
	res := db.conn.Select("Email", "Password", "FirstName", "LastName").Create(user)
	return res.Error
}

func (db PostgresRepository) GetUserByEmail(email string) *User {
	var user User

	db.conn.
		Table("users").
		Select("id, email, password, first_name, last_name, photo_url").
		Where("email = ?", email).Scan(&user)

	return &user
}

func (db PostgresRepository) GetUser(userId int) *UserData {
	var user UserData
	db.conn.Table("users").Select("id, first_name, last_name, photo_url").Where("id = ?", userId).Scan(&user)
	return &user
}

func (db PostgresRepository) GetUserPassword(userId int) (string, error) {
	var password string
	res := db.conn.Table("users").Select("password").Where("id = ?", userId).Scan(&password)
	return password, res.Error
}

func (db PostgresRepository) GetTours(userId int) []TourData {
	var tours []TourData
	db.conn.
		Table("tours").
		Select("id, tour_name, tour_price, tour_date_from, tour_date_to").
		Where("user_id = ?", userId).Scan(&tours)
	if tours == nil {
		return []TourData{}
	}
	return tours
}

func (db PostgresRepository) GetCityTours(tourId int) []CityTourData {
	var cityTours []CityTourData
	db.conn.
		Table("city_tours").
		Select("city_tours.id, country_name, cities.id AS city_id, city_name, city_tour_price, date_from, date_to, ticket_arrival_id, ticket_departure_id, hotels.id AS hotel_id").
		Joins("JOIN hotels ON city_tours.hotel_id = hotels.id").
		Joins("JOIN cities ON city_tours.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("tour_id = ?", tourId).Scan(&cityTours)
	if cityTours == nil {
		return []CityTourData{}
	}
	return cityTours
}

func (db PostgresRepository) GetUserEvents(cityTourId int) []EventData {
	var events []EventData
	db.conn.
		Table("city_tours").
		Select("events.id AS id, event_name, event_start, event_end, event_price AS price, event_rating AS rating, max_persons, cur_persons").
		Joins("JOIN city_tours_events ON city_tours.id = city_tours_events.ct_id").
		Joins("JOIN events ON city_tours_events.event_id = events.id").
		Where("city_tours.id = ?", cityTourId).Scan(&events)
	if events == nil {
		return []EventData{}
	}
	return events
}

func (db PostgresRepository) GetRestaurantBookings(cityTourId int) []RestaurantBookingData {
	var restaurantBookings []RestaurantBookingData
	db.conn.
		Table("city_tours").
		Select("restaurant_bookings.id AS id, restaurant_id, booking_time, rest_name AS restaurant_name, avg_price AS average_price, rest_rating AS rating").
		Joins("JOIN city_tours_rest_bookings ON city_tours.id = city_tours_rest_bookings.ct_id").
		Joins("JOIN restaurant_bookings ON city_tours_rest_bookings.rb_id = restaurant_bookings.id").
		Joins("JOIN restaurants ON restaurant_bookings.restaurant_id = restaurants.id").
		Where("city_tours.id = ?", cityTourId).Scan(&restaurantBookings)
	if restaurantBookings == nil {
		return []RestaurantBookingData{}
	}
	return restaurantBookings
}

func (db PostgresRepository) getUserTicket(ticketId int) *TicketData {
	var ticket TicketData
	db.conn.
		Table("tickets").
		Select("tickets.id, transport_type, price, ticket_date, "+
			"orig_c.country_name, orig_city.city_name, "+
			"dest_c.country_name, dest_city.city_name, "+
			"company_name, company_rating").
		Joins("JOIN transport_companies on tickets.company_id = transport_companies.id").
		Joins("JOIN transport_stations orig_ts ON tickets.orig_station_id = orig_ts.id").
		Joins("JOIN cities orig_city ON orig_ts.city_id = orig_city.id").
		Joins("JOIN countries orig_c ON orig_city.country_id = orig_c.id").
		Joins("JOIN transport_stations dest_ts ON tickets.dest_station_id = dest_ts.id").
		Joins("JOIN cities dest_city ON dest_ts.city_id = dest_city.id").
		Joins("JOIN countries dest_c ON dest_city.country_id = dest_c.id").
		Where("id = ?", ticketId).Scan(&ticket)

	return &ticket
}

func (db PostgresRepository) GetUserTickets(cityTourId int) *CityTourTicketData {
	var ticketIds CityTourTicketIdData
	var tickets CityTourTicketData
	var ticketData *TicketData

	db.conn.
		Table("city_tours").
		Select("ticket_arrival_id, ticket_departure_id").
		Where("id = ?", cityTourId).Scan(&ticketIds)

	ticketData = db.getUserTicket(ticketIds.TicketArrivalId)
	if ticketData == nil {
		return nil
	}
	tickets.ArrivalTicket = *ticketData

	ticketData = db.getUserTicket(ticketIds.TicketDepartureId)
	if ticketData == nil {
		return nil
	}
	tickets.DepartureTicket = *ticketData

	return &tickets
}

func (db PostgresRepository) GetUserHotel(cityTourId int) *HotelData {
	var hotel HotelData
	db.conn.
		Table("city_tours").
		Select("hotels.id, hotel_name, stars, hotel_rating").
		Joins("JOIN hotels ON city_tours.hotel_id = hotels.id").
		Where("city_tours.id = ?", cityTourId).Scan(&hotel)
	return &hotel
}

func (db PostgresRepository) CreateTour(tour *Tour) error {
	res := db.conn.Select("UserId", "TourName", "TourPrice", "TourDateFrom", "TourDateTo").Create(tour)
	return res.Error
}

func (db PostgresRepository) CreateCityTour(cityTour *CityTour) error {
	res := db.conn.
		Select("TourId", "CityId", "CityTourPrice", "DateFrom", "DateTo", "TicketArrivalId", "TicketDepartureId", "HotelId").
		Create(cityTour)
	return res.Error
}

func (db PostgresRepository) CreateCityTourEvent(cityTourEvent *CityToursEvent) error {
	res := db.conn.Select("CtId", "EventId").Create(cityTourEvent)
	return res.Error
}

func (db PostgresRepository) CreateRestaurantBooking(restaurantBooking *RestaurantBookingDTO) error {
	dao := RestaurantBooking{
		Id:           0,
		RestaurantId: restaurantBooking.RestaurantId,
		BookingTime:  restaurantBooking.BookingTime,
	}
	res := db.conn.Select("RestaurantId", "BookingTime").Create(&dao)

	ctRb := CityToursRestBooking{
		CtId: restaurantBooking.CtId,
		RbId: dao.Id,
	}
	res = db.conn.Select("CtId", "RbId").Create(ctRb)
	return res.Error
}

func (db PostgresRepository) UpdateUser(user *User) error {
	res := db.conn.Save(user)
	return res.Error
}

func (db PostgresRepository) UpdateTour(tour *Tour) error {
	res := db.conn.Save(tour)
	return res.Error
}

func (db PostgresRepository) UpdateCityTour(cityTour *CityTour) error {
	db.conn.Exec(
		"CALL public.update_ct_price(?, ?, ?, ?)",
		cityTour.TicketArrivalId,
		cityTour.TicketDepartureId,
		cityTour.HotelId,
		cityTour.Id,
	)
	var cityTourPrice string
	db.conn.
		Table("city_tours").
		Select("city_tour_price").
		Where("id = ?", cityTour.Id).
		Scan(&cityTourPrice)
	newCityTour := cityTour
	newCityTour.CityTourPrice, _ = strconv.ParseFloat(cityTourPrice[1:], 64)

	res := db.conn.Save(&newCityTour)
	return res.Error
}

func (db PostgresRepository) UpdateRestaurantBooking(restaurantBooking *RestaurantBooking) error {
	res := db.conn.Save(restaurantBooking)
	return res.Error
}

func (db PostgresRepository) DeleteUser(userId int) error {
	res := db.conn.Delete(&User{}, userId)
	return res.Error
}

func (db PostgresRepository) DeleteTour(tourId int) error {
	res := db.conn.Delete(&Tour{}, tourId)
	return res.Error
}

func (db PostgresRepository) DeleteCityTour(cityTourId int) error {
	res := db.conn.Delete(&CityTour{}, cityTourId)
	return res.Error
}

func (db PostgresRepository) DeleteRestaurantBooking(restaurantBookingId int) error {
	res := db.conn.Delete(&RestaurantBooking{}, restaurantBookingId)
	return res.Error
}

func (db PostgresRepository) GetCities() ([]City, error) {
	var cities []City

	res := db.conn.
		Table("cities").
		Select("id AS city_id, city_name").
		Order("city_name").
		Scan(&cities)

	return cities, res.Error
}

func (db PostgresRepository) GetHotel(id int) (*Hotel, error) {
	var hotel Hotel

	res := db.conn.
		Table("hotels").
		Select("hotels.id AS hotel_id, hotel_name, hotel_description, hotel_addr, stars, hotel_rating, avg_price, near_sea, country_name, city_name").
		Joins("JOIN cities ON hotels.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("hotels.id = ?", id).Scan(&hotel)

	return &hotel, res.Error
}

func (db PostgresRepository) GetEvent(id int) (*Event, error) {
	var event Event

	res := db.conn.
		Table("events").
		Select("events.id AS event_id, event_name, event_description, event_addr, country_name, city_name, event_start, event_end, event_price AS price, event_rating AS rating, max_persons, cur_persons, languages").
		Joins("JOIN cities ON events.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("events.id = ?", id).Scan(&event)

	return &event, res.Error
}

func (db PostgresRepository) GetRestaurant(id int) (*Restaurant, error) {
	var restaurant Restaurant

	res := db.conn.
		Table("restaurants").
		Select("restaurants.id AS restaurant_id, rest_name, rest_description, rest_addr, avg_price, rest_rating, child_menu, smoking_room, country_name, city_name").
		Joins("JOIN cities ON restaurants.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("restaurants.id = ?", id).Scan(&restaurant)

	return &restaurant, res.Error
}

func (db PostgresRepository) GetTicket(id int) (*Ticket, error) {
	var ticket Ticket

	res := db.conn.
		Table("tickets").
		Select(
			"tickets.id AS ticket_id, company_name, company_rating, "+
				"orig_ts.station_name AS orig_station_name, orig_ts.station_addr AS orig_station_addr, "+
				"orig_c.country_name AS orig_country_name, orig_city.city_name AS orig_city_name, "+
				"dest_ts.station_name AS dest_station_name, dest_ts.station_addr AS dest_station_addr, "+
				"dest_c.country_name AS dest_country_name, dest_city.city_name AS dest_city_name, "+
				"transport_type, price, ticket_date").
		Joins("JOIN transport_companies on tickets.company_id = transport_companies.id").
		Joins("JOIN transport_stations orig_ts ON tickets.orig_station_id = orig_ts.id").
		Joins("JOIN cities orig_city ON orig_ts.city_id = orig_city.id").
		Joins("JOIN countries orig_c ON orig_city.country_id = orig_c.id").
		Joins("JOIN transport_stations dest_ts ON tickets.dest_station_id = dest_ts.id").
		Joins("JOIN cities dest_city ON dest_ts.city_id = dest_city.id").
		Joins("JOIN countries dest_c ON dest_city.country_id = dest_c.id").
		Where("tickets.id = ?", id).Scan(&ticket)

	return &ticket, res.Error
}

func (db PostgresRepository) GetHotels(filter HotelFilterParameters) ([]HotelFeedItem, error) {
	var hotels []HotelFeedItem
	var order string

	if filter.OrderType == "inc" {
		order = filter.OrderBy + " " + "ASC"
	} else if filter.OrderType == "dec" {
		order = filter.OrderBy + " " + "DESC"
	} else {
		return nil, fmt.Errorf("wrong order type")
	}

	res := db.conn.
		Table("hotels").
		Select("hotels.id AS hotel_id, hotel_name, stars, hotel_rating, country_name, city_name").
		Joins("JOIN cities ON hotels.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("hotels.id <> 0").
		Where("stars BETWEEN ? AND ?", filter.StarsFrom, filter.StarsTo).
		Where("hotel_rating BETWEEN ? AND ?", filter.RatingFrom, filter.RatingTo).
		Where("avg_price BETWEEN ? AND ?", filter.PriceFrom, filter.PriceTo)

	if filter.HotelName != "" {
		res = res.Where("hotel_name LIKE ?", filter.HotelName)
	}

	if filter.CityName != "" {
		res = res.Where("city_name LIKE ?", filter.CityName)
	}

	if filter.NearSea {
		res = res.Where("near_sea = ?", filter.NearSea)
	}

	res = res.Order(order).Offset(filter.Offset).Limit(filter.Limit).Scan(&hotels)

	return hotels, res.Error
}

func (db PostgresRepository) GetEvents(filter EventsFilterParameters) ([]EventFeedItem, error) {
	var events []EventFeedItem
	var order string

	if filter.OrderType == "inc" {
		order = filter.OrderBy + " " + "ASC"
	} else if filter.OrderType == "dec" {
		order = filter.OrderBy + " " + "DESC"
	} else {
		return nil, fmt.Errorf("wrong order type")
	}

	res := db.conn.
		Table("events").
		Select("events.id AS event_id, event_name, event_start, event_end, event_rating AS rating, max_persons, cur_persons, country_name, city_name").
		Joins("JOIN cities ON events.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("event_rating BETWEEN ? AND ?", filter.RatingFrom, filter.RatingTo).
		Where("event_price BETWEEN ? AND ?", filter.PriceFrom, filter.PriceTo)

	if filter.From != "" {
		res = res.Where("event_start <= ?", filter.From)
	}

	if filter.To != "" {
		res = res.Where("event_end <= ?", filter.To)
	}

	if filter.EventName != "" {
		res = res.Where("event_name LIKE ?", filter.EventName)
	}

	if filter.CityName != "" {
		res = res.Where("city_name LIKE ?", filter.CityName)
	}

	res = res.Order(order).Offset(filter.Offset).Limit(filter.Limit).Scan(&events)

	return events, res.Error
}

func (db PostgresRepository) GetRestaurants(filter RestaurantFilterParameters) ([]RestaurantFeedItem, error) {
	var restaurants []RestaurantFeedItem
	var order string

	if filter.OrderType == "inc" {
		order = filter.OrderBy + " " + "ASC"
	} else if filter.OrderType == "dec" {
		order = filter.OrderBy + " " + "DESC"
	} else {
		return nil, fmt.Errorf("wrong order type")
	}

	res := db.conn.
		Table("restaurants").
		Select("restaurants.id AS restaurant_id, rest_name AS restaurant_name, rest_rating AS rating, country_name, city_name").
		Joins("JOIN cities ON restaurants.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("rest_rating BETWEEN ? AND ?", filter.RatingFrom, filter.RatingTo).
		Where("avg_price BETWEEN ? AND ?", filter.PriceFrom, filter.PriceTo)

	if filter.RestaurantName != "" {
		res = res.Where("rest_name LIKE ?", filter.RestaurantName)
	}

	if filter.CityName != "" {
		res = res.Where("city_name LIKE ?", filter.CityName)
	}

	if filter.ChildMenu {
		res = res.Where("child_menu = ?", filter.ChildMenu)
	}

	if filter.SmokingRoom {
		res = res.Where("smoking_room = ?", filter.SmokingRoom)
	}

	res = res.Order(order).Offset(filter.Offset).Limit(filter.Limit).Scan(&restaurants)

	return restaurants, res.Error
}

func (db PostgresRepository) GetTickets(filter TicketFilterParameters) ([]TicketFeedItem, error) {
	var tickets []TicketFeedItem
	var order string

	if filter.OrderType == "inc" {
		order = filter.OrderBy + " " + "ASC"
	} else if filter.OrderType == "dec" {
		order = filter.OrderBy + " " + "DESC"
	} else {
		return nil, fmt.Errorf("wrong order type")
	}

	res := db.conn.
		Table("tickets").
		Select("tickets.id AS ticket_id, transport_type, price, ticket_date AS date, "+
			"orig_c.country_name AS orig_country_name, orig_city.city_name AS orig_city_name, "+
			"dest_c.country_name AS dest_country_name, dest_city.city_name AS dest_city_name, "+
			"company_name, company_rating").
		Joins("JOIN transport_companies on tickets.company_id = transport_companies.id").
		Joins("JOIN transport_stations orig_ts ON tickets.orig_station_id = orig_ts.id").
		Joins("JOIN cities orig_city ON orig_ts.city_id = orig_city.id").
		Joins("JOIN countries orig_c ON orig_city.country_id = orig_c.id").
		Joins("JOIN transport_stations dest_ts ON tickets.dest_station_id = dest_ts.id").
		Joins("JOIN cities dest_city ON dest_ts.city_id = dest_city.id").
		Joins("JOIN countries dest_c ON dest_city.country_id = dest_c.id").
		Where("tickets.id <> 0").
		Where("price BETWEEN ? AND ?", filter.PriceFrom, filter.PriceTo)

	if filter.DateFrom != "" {
		res = res.Where("ticket_date >= ?", filter.DateFrom)
	}

	if filter.DateTo != "" {
		res = res.Where("ticket_date <= ?", filter.DateTo)
	}

	if filter.TransportType != "" {
		res = res.Where("transport_type LIKE ?", filter.TransportType)
	}

	if filter.OrigCityName != "" {
		res = res.Where("orig_city.city_name LIKE ?", filter.OrigCityName)
	}

	if filter.DestCityName != "" {
		res = res.Where("dest_city.city_name LIKE ?", filter.DestCityName)
	}

	res = res.Order(order).Offset(filter.Offset).Limit(filter.Limit).Scan(&tickets)

	return tickets, res.Error
}
