package db

import (
	"github.com/SmartDuck9000/travelly-api/services/user_service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserProfileDb interface {
	Open() error
	configureConnectionPools() error

	GetUser(userId int) *UserData
	GetUserPassword(userId int) (string, error)
	GetTours(userId int) []TourData
	GetCityTours(tourId int) []CityTourData

	GetEvents(cityTourId int) []EventData
	GetRestaurantBookings(cityTourId int) []RestaurantBookingData
	GetTickets(cityTourId int) *CityTourTicketData
	GetHotel(cityTourId int) *HotelData

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

func (db UserProfilePostgres) GetUserPassword(userId int) (string, error) {
	var password string
	res := db.conn.Table("users").Select("password").Where("id = ?", userId).Scan(&password)
	return password, res.Error
}

func (db UserProfilePostgres) GetTours(userId int) []TourData {
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

func (db UserProfilePostgres) GetCityTours(tourId int) []CityTourData {
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

func (db UserProfilePostgres) GetEvents(cityTourId int) []EventData {
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

func (db UserProfilePostgres) GetRestaurantBookings(cityTourId int) []RestaurantBookingData {
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

	ticketData = db.getTicket(ticketIds.TicketArrivalId)
	if ticketData == nil {
		return nil
	}
	tickets.ArrivalTicket = *ticketData

	ticketData = db.getTicket(ticketIds.TicketDepartureId)
	if ticketData == nil {
		return nil
	}
	tickets.DepartureTicket = *ticketData

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

func (db UserProfilePostgres) CreateTour(tour *Tour) error {
	res := db.conn.Select("UserId", "TourName", "TourPrice", "TourDateFrom", "TourDateTo").Create(tour)
	return res.Error
}

func (db UserProfilePostgres) CreateCityTour(cityTour *CityTour) error {
	res := db.conn.
		Select("TourId", "CityId", "CityTourPrice", "DateFrom", "DateTo", "TicketArrivalId", "TicketDepartureId", "HotelId").
		Create(cityTour)
	return res.Error
}

func (db UserProfilePostgres) CreateCityTourEvent(cityTourEvent *CityToursEvent) error {
	res := db.conn.Select("CtId", "EventId").Create(cityTourEvent)
	return res.Error
}

func (db UserProfilePostgres) CreateRestaurantBooking(restaurantBooking *RestaurantBookingDTO) error {
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

func (db UserProfilePostgres) UpdateUser(user *User) error {
	res := db.conn.Save(user)
	return res.Error
}

func (db UserProfilePostgres) UpdateTour(tour *Tour) error {
	res := db.conn.Save(tour)
	return res.Error
}

func (db UserProfilePostgres) UpdateCityTour(cityTour *CityTour) error {
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

func (db UserProfilePostgres) UpdateRestaurantBooking(restaurantBooking *RestaurantBooking) error {
	res := db.conn.Save(restaurantBooking)
	return res.Error
}

func (db UserProfilePostgres) DeleteUser(userId int) error {
	res := db.conn.Delete(&User{}, userId)
	return res.Error
}

func (db UserProfilePostgres) DeleteTour(tourId int) error {
	res := db.conn.Delete(&Tour{}, tourId)
	return res.Error
}

func (db UserProfilePostgres) DeleteCityTour(cityTourId int) error {
	res := db.conn.Delete(&CityTour{}, cityTourId)
	return res.Error
}

func (db UserProfilePostgres) DeleteRestaurantBooking(restaurantBookingId int) error {
	res := db.conn.Delete(&RestaurantBooking{}, restaurantBookingId)
	return res.Error
}
