package db

import (
	"github.com/SmartDuck9000/travelly-api/services/feed_service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type FeedDB interface {
	Open() error
	configureConnectionPools() error

	GetHotels(orderedBy string) []Hotel
	GetEvents(orderedBy string) []Event
	GetRestaurants(orderedBy string) []Restaurant
}

type FeedPostgres struct {
	url             string
	maxIdleConn     int
	maxOpenConn     int
	connMaxLifetime time.Duration
	conn            *gorm.DB
}

func CreateFeedServiceDB(conf config.FeedDBConfig) *FeedPostgres {
	return &FeedPostgres{
		url:             conf.URL,
		maxIdleConn:     conf.MaxIdleConn,     // maximum number of connections in the idle connection pool
		maxOpenConn:     conf.MaxOpenConn,     // maximum number of open connections to the database
		connMaxLifetime: conf.ConnMaxLifetime, // maximum amount of time a connection may be reused
		conn:            nil,
	}
}

func (db FeedPostgres) Open() error {
	var err error
	db.conn, err = gorm.Open(postgres.Open(db.url), &gorm.Config{})
	if err == nil {
		err = db.configureConnectionPools()
	}
	return err
}

func (db FeedPostgres) configureConnectionPools() error {
	sqlDB, err := db.conn.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(db.maxIdleConn)
	sqlDB.SetMaxOpenConns(db.maxOpenConn)
	sqlDB.SetConnMaxLifetime(db.connMaxLifetime)

	return nil
}

func (db FeedPostgres) GetHotels(orderedBy string) []Hotel {
	var hotels []Hotel

	db.conn.
		Table("hotels").
		Select("hotels.id, hotel_name, hotel_description, hotel_addr, stars, hotel_rating, avg_price, near_sea, country_name, city_name").
		Joins("JOIN cities ON hotels.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Order(orderedBy).Scan(&hotels)

	return hotels
}

func (db FeedPostgres) GetEvents(orderedBy string) []Event {
	var events []Event

	db.conn.
		Table("events").
		Select("events.id, event_name, event_description, event_addr, country_name, city_name, event_start, event_end, event_price, event_rating, max_persons, cur_persons, languages").
		Joins("JOIN cities ON hotels.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Order(orderedBy).Scan(&events)

	return events
}

func (db FeedPostgres) GetRestaurants(orderedBy string) []Restaurant {
	var restaurants []Restaurant

	db.conn.
		Table("restaurants").
		Select("restaurants.id, rest_name, rest_description, rest_addr, avg_price, rest_rating, child_menu, smoking_room, country_name, city_name").
		Joins("JOIN cities ON hotels.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Order(orderedBy).Scan(&restaurants)

	return restaurants
}
