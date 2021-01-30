package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type TravellyDb interface {
	Open() error
	Close()

	GetUser(userId int) *User
	GetTours(userId int) []Tour
	GetCityTours(tourId int) []CityTour

	GetEvents(cityTourId int) []Event
	GetRestaurantBookings(cityTourId int) []RestaurantBooking
	GetTickets(cityTourId int) *Ticket
	GetHotel(cityTourId int) *Hotel
}

type TravellyPostgres struct {
	url  string
	conn *gorm.DB
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

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}

func (db TravellyPostgres) GetUser(userId int) *User {
	return nil
}

func (db TravellyPostgres) GetTour(tourId int) *Tour {
	return nil
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

func (db TravellyPostgres) GetTicket(ticketId int) *Ticket {
	return nil
}

func (db TravellyPostgres) GetHotel(hotelId int) *Hotel {
	return nil
}
