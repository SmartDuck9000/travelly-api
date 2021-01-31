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
	return nil
}

func (db TravellyPostgres) GetEvents(cityTourId int) []Event {
	return nil
}

func (db TravellyPostgres) GetRestaurantBookings(cityTourId int) []RestaurantBooking {
	return nil
}

func (db TravellyPostgres) GetTickets(cityTourId int) []Ticket {
	return nil
}

func (db TravellyPostgres) GetHotel(cityTourId int) *Hotel {
	return nil
}
