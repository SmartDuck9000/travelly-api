package db

import (
	"github.com/SmartDuck9000/travelly-api/services/user_service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type UserProfileDb interface {
	Open() error
	configureConnectionPools() error

	GetUser(userId int) *UserData
	GetTours(userId int) []TourData
	GetCityTours(tourId int) []CityTourData

	GetEvents(cityTourId int) []EventData
	GetRestaurantBookings(cityTourId int) []RestaurantBookingData
	GetTickets(cityTourId int) *CityTourTicketData
	GetHotel(cityTourId int) *HotelData

	CreateTour(tour *TourEntity) error
	CreateCityTour(cityTour *CityTourEntity) error
	CreateRestaurantBooking(restaurantBooking *RestaurantBookingEntity) error

	UpdateUser(user *UserEntity) error
	UpdateTour(tour *TourEntity) error
	UpdateCityTour(cityTour *CityTourEntity) error
	UpdateRestaurantBooking(restaurantBooking *RestaurantBookingEntity) error

	DeleteUser(userId int) error
	DeleteTour(tourId int) error
	DeleteCityTour(cityTourId int) error
	DeleteRestaurantBooking(restaurantBookingId int) error
}

type UserProfilePostgres struct {
	url             string
	maxIdleConn     int
	maxOpenConn     int
	connMaxLifetime time.Duration
	conn            *gorm.DB
}

func CreateUserServiceDb(conf config.UserDbConfig) UserProfileDb {
	return &UserProfilePostgres{
		url:             conf.URL,
		maxIdleConn:     conf.MaxIdleConn,     // maximum number of connections in the idle connection pool
		maxOpenConn:     conf.MaxOpenConn,     // maximum number of open connections to the database
		connMaxLifetime: conf.ConnMaxLifetime, // maximum amount of time a connection may be reused
		conn:            nil,
	}
}

func (db *UserProfilePostgres) Open() error {
	var err error
	db.conn, err = gorm.Open(postgres.Open(db.url), &gorm.Config{})
	if err == nil {
		err = db.configureConnectionPools()
	}
	return err
}

func (db UserProfilePostgres) configureConnectionPools() error {
	sqlDB, err := db.conn.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(db.maxIdleConn)
	sqlDB.SetMaxOpenConns(db.maxOpenConn)
	sqlDB.SetConnMaxLifetime(db.connMaxLifetime)

	return nil
}

func (db UserProfilePostgres) GetUser(userId int) *UserData {
	var user UserData
	db.conn.Table("users").Select("id, first_name, last_name, photo_url").Where("id = ?", userId).Scan(&user)
	return &user
}

func (db UserProfilePostgres) GetTours(userId int) []TourData {
	var tours []TourData
	db.conn.
		Table("tours").
		Select("id, tour_name, tour_price, tour_date_from, tour_date_to").
		Where("user_id = ?", userId).Scan(&tours)
	return tours
}

func (db UserProfilePostgres) GetCityTours(tourId int) []CityTourData {
	var cityTours []CityTourData
	db.conn.
		Table("city_tours").
		Select("city_tours.id, country_name, city_name, city_tour_price, date_from, date_to, ticket_arrival_id, ticket_departure_id, hotel_name").
		Joins("JOIN hotels ON city_tours.hotel_id = hotels.id").
		Joins("JOIN cities ON city_tours.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("tour_id = ?", tourId).Scan(&cityTours)
	return cityTours
}

func (db UserProfilePostgres) GetEvents(cityTourId int) []EventData {
	var events []EventData
	db.conn.
		Table("city_tours").
		Select("events.id, event_name, event_start, event_end, event_price, event_rating, max_persons, cur_persons").
		Joins("JOIN city_tours_events ON city_tours.id = city_tours_events.ct_id").
		Joins("JOIN events ON city_tours_events.event_id = events.id").
		Where("city_tours.id = ?", cityTourId).Scan(&events)
	return events
}

func (db UserProfilePostgres) GetRestaurantBookings(cityTourId int) []RestaurantBookingData {
	var restaurantBookings []RestaurantBookingData
	db.conn.
		Table("city_tours").
		Select("restaurant_bookings.id, restaurant_id, booking_time, rest_name, avg_price, rest_rating").
		Joins("JOIN city_tours_rest_bookings ON city_tours.id = city_tours_rest_bookings.ct_id").
		Joins("JOIN restaurant_bookings ON city_tours_rest_bookings.rb_id = restaurant_bookings.id").
		Joins("JOIN restaurants ON restaurant_bookings.restaurant_id = restaurants.id").
		Where("city_tours.id = ?", cityTourId).Scan(&restaurantBookings)
	return restaurantBookings
}

func (db UserProfilePostgres) getTicket(ticketId int) *TicketData {
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

func (db UserProfilePostgres) GetTickets(cityTourId int) *CityTourTicketData {
	var ticketIds CityTourTicketIdData
	var tickets CityTourTicketData
	var ticketData *TicketData

	db.conn.
		Table("city_tours").
		Select("ticket_arrival_id, ticket_departure_id").
		Where("id = ?", cityTourId).Scan(&ticketIds)

	ticketData = db.getTicket(ticketIds.ticketArrivalId)
	if ticketData == nil {
		return nil
	}
	tickets.arrivalTicket = *ticketData

	ticketData = db.getTicket(ticketIds.ticketDepartureId)
	if ticketData == nil {
		return nil
	}
	tickets.departureTicket = *ticketData

	return &tickets
}

func (db UserProfilePostgres) GetHotel(cityTourId int) *HotelData {
	var hotel HotelData
	db.conn.
		Table("city_tours").
		Select("hotels.id, hotel_name, stars, hotel_rating").
		Joins("JOIN hotels ON city_tours.hotel_id = hotels.id").
		Where("city_tours.id = ?", cityTourId).Scan(&hotel)
	return &hotel
}

func (db UserProfilePostgres) CreateTour(tour *TourEntity) error {
	res := db.conn.Select("userId", "tourName", "tourPrice", "tourDateFrom", "tourDateTo").Create(tour)
	return res.Error
}

func (db UserProfilePostgres) CreateCityTour(cityTour *CityTourEntity) error {
	res := db.conn.
		Select("tourId", "cityId", "cityTourPrice", "dateFrom", "dateTo", "ticketArrivalId", "ticketDepartureId", "hotelId").
		Create(cityTour)
	return res.Error
}

func (db UserProfilePostgres) CreateRestaurantBooking(restaurantBooking *RestaurantBookingEntity) error {
	res := db.conn.Select("restaurantId", "bookingTime").Create(restaurantBooking)
	return res.Error
}

func (db UserProfilePostgres) UpdateUser(user *UserEntity) error {
	res := db.conn.Save(user)
	return res.Error
}

func (db UserProfilePostgres) UpdateTour(tour *TourEntity) error {
	res := db.conn.Save(tour)
	return res.Error
}

func (db UserProfilePostgres) UpdateCityTour(cityTour *CityTourEntity) error {
	res := db.conn.Save(cityTour)
	return res.Error
}

func (db UserProfilePostgres) UpdateRestaurantBooking(restaurantBooking *RestaurantBookingEntity) error {
	res := db.conn.Save(restaurantBooking)
	return res.Error
}

func (db UserProfilePostgres) DeleteUser(userId int) error {
	res := db.conn.Delete(&UserEntity{}, userId)
	return res.Error
}

func (db UserProfilePostgres) DeleteTour(tourId int) error {
	res := db.conn.Delete(&TourEntity{}, tourId)
	return res.Error
}

func (db UserProfilePostgres) DeleteCityTour(cityTourId int) error {
	res := db.conn.Delete(&CityTourEntity{}, cityTourId)
	return res.Error
}

func (db UserProfilePostgres) DeleteRestaurantBooking(restaurantBookingId int) error {
	res := db.conn.Delete(&RestaurantBookingEntity{}, restaurantBookingId)
	return res.Error
}
