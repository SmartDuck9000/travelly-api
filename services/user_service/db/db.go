package db

import (
	"github.com/SmartDuck9000/travelly-api/services/user_service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type TravellyDb interface {
	Open() error
	configureConnectionPools() error

	GetUser(userId int) *User
	GetTours(userId int) []Tour
	GetCityTours(tourId int) []CityTour

	GetEvents(cityTourId int) []Event
	GetRestaurantBookings(cityTourId int) []RestaurantBooking
	GetTickets(cityTourId int) []Ticket
	GetHotel(cityTourId int) *Hotel
}

type TravellyPostgres struct {
	url             string
	maxIdleConn     int
	maxOpenConn     int
	connMaxLifetime time.Duration
	conn            *gorm.DB
}

func CreateUserServiceDb(conf config.UserServiceDbConfig) *TravellyPostgres {
	return &TravellyPostgres{
		url:             conf.URL,
		maxIdleConn:     conf.MaxIdleConn,     // maximum number of connections in the idle connection pool
		maxOpenConn:     conf.MaxOpenConn,     // maximum number of open connections to the database
		connMaxLifetime: conf.ConnMaxLifetime, // maximum amount of time a connection may be reused
		conn:            nil,
	}
}

func (db TravellyPostgres) Open() error {
	var err error
	db.conn, err = gorm.Open(postgres.Open(db.url), &gorm.Config{})
	if err == nil {
		err = db.configureConnectionPools()
	}
	return err
}

func (db TravellyPostgres) configureConnectionPools() error {
	sqlDB, err := db.conn.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(db.maxIdleConn)
	sqlDB.SetMaxOpenConns(db.maxOpenConn)
	sqlDB.SetConnMaxLifetime(db.connMaxLifetime)

	return nil
}

func (db TravellyPostgres) GetUser(userId int) *User {
	var user User
	db.conn.Table("users").Select("id, first_name, last_name, photo_url").Where("id = ?", userId).Scan(&user)
	return &user
}

func (db TravellyPostgres) GetTours(userId int) []Tour {
	var tours []Tour
	db.conn.
		Table("tours").
		Select("id, user_id, tour_name, tour_price, tour_date_from, tour_date_to").
		Where("user_id = ?", userId).Scan(&tours)
	return tours
}

func (db TravellyPostgres) GetCityTours(tourId int) []CityTour {
	var cityTours []CityTour
	db.conn.
		Table("city_tours").
		Select("city_tours.id, tour_id, country_name, city_name, city_tour_price, date_from, date_to, ticket_arrival_id, ticket_departure_id, hotel_name").
		Joins("JOIN hotels ON city_tours.hotel_id = hotels.id").
		Joins("JOIN cities ON city_tours.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("tour_id = ?", tourId).Scan(&cityTours)
	return cityTours
}

func (db TravellyPostgres) GetEvents(cityTourId int) []Event {
	var events []Event
	db.conn.
		Table("city_tours").
		Select("events.id, event_name, event_description, event_addr, country_name, city_name, event_start, event_end, price, rating, max_persons, cur_persons, languages").
		Joins("JOIN city_tours_events ON city_tours.id = city_tours_events.ct_id").
		Joins("JOIN events ON city_tours_events.event_id = events.id").
		Joins("JOIN cities ON city_tours.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("city_tour_id = ?", cityTourId).Scan(&events)
	return events
}

func (db TravellyPostgres) GetRestaurantBookings(cityTourId int) []RestaurantBooking {
	var restaurantBookings []RestaurantBooking
	db.conn.
		Table("city_tours").
		Select("restaurant_bookings.id, restaurant_id, booking_time, restaurant_name, restaurant_description, restaurant_addr, average_price, rating, child_menu, smoking_room, country_name, city_name").
		Joins("JOIN city_tours_rest_bookings ON city_tours.id = city_tours_rest_bookings.ct_id").
		Joins("JOIN restaurant_bookings ON city_tours_rest_bookings.rb_id = restaurant_bookings.id").
		Joins("JOIN restaurants ON restaurant_bookings.restaurant_id = restaurants.id").
		Joins("JOIN cities ON city_tours.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("city_tour_id = ?", cityTourId).Scan(&restaurantBookings)
	return restaurantBookings
}

func (db TravellyPostgres) GetTickets(cityTourId int) []Ticket {
	return nil
}

func (db TravellyPostgres) GetHotel(cityTourId int) *Hotel {
	return nil
}
