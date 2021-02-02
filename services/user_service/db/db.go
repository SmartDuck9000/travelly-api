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
	GetTickets(cityTourId int) [2]Ticket
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
		Select("events.id, event_name, event_description, event_addr, country_name, city_name, event_start, event_end, event_price, event_rating, max_persons, cur_persons, languages").
		Joins("JOIN city_tours_events ON city_tours.id = city_tours_events.ct_id").
		Joins("JOIN events ON city_tours_events.event_id = events.id").
		Joins("JOIN cities ON city_tours.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("city_tours.id = ?", cityTourId).Scan(&events)
	return events
}

func (db TravellyPostgres) GetRestaurantBookings(cityTourId int) []RestaurantBooking {
	var restaurantBookings []RestaurantBooking
	db.conn.
		Table("city_tours").
		Select("restaurant_bookings.id, restaurant_id, booking_time, rest_name, rest_description, rest_addr, avg_price, rest_rating, child_menu, smoking_room, country_name, city_name").
		Joins("JOIN city_tours_rest_bookings ON city_tours.id = city_tours_rest_bookings.ct_id").
		Joins("JOIN restaurant_bookings ON city_tours_rest_bookings.rb_id = restaurant_bookings.id").
		Joins("JOIN restaurants ON restaurant_bookings.restaurant_id = restaurants.id").
		Joins("JOIN cities ON city_tours.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("city_tours.id = ?", cityTourId).Scan(&restaurantBookings)
	return restaurantBookings
}

func (db TravellyPostgres) getTicket(ticketId int) *Ticket {
	var ticket Ticket
	db.conn.
		Table("tickets").
		Select("tickets.id, transport_type, price, ticket_date, orig_ts.station_name, orig_ts.station_addr, dest_ts.station_name, dest_ts.station_addr, company_name, company_rating").
		Joins("JOIN transport_companies on tickets.company_id = transport_companies.id").
		Joins("JOIN transport_stations orig_ts ON tickets.orig_station_id = orig_ts.id").
		Joins("JOIN transport_stations dest_ts ON tickets.dest_station_id = dest_ts.id").
		Where("id = ?", ticketId).Scan(&ticket)

	return &ticket
}

func (db TravellyPostgres) GetTickets(cityTourId int) [2]Ticket {
	var ticketIds CityTourTicketID
	var tickets [2]Ticket

	db.conn.
		Table("city_tours").
		Select("ticket_arrival_id, ticket_departure_id").
		Where("id = ?", cityTourId).Scan(&ticketIds)

	tickets[0] = *db.getTicket(ticketIds.ticketArrivalId)
	tickets[1] = *db.getTicket(ticketIds.ticketDepartureId)

	return tickets
}

func (db TravellyPostgres) GetHotel(cityTourId int) *Hotel {
	var hotel Hotel
	db.conn.
		Table("city_tours").
		Select("hotels.id, hotel_name, hotel_description, hotel_addr, stars, hotel_rating, avg_price, near_sea, country_name, city_name").
		Joins("JOIN hotels ON city_tours.hotel_id = hotels.id").
		Joins("JOIN cities ON city_tours.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("city_tours.id = ?", cityTourId).Scan(&hotel)
	return &hotel
}
