package db

import (
	"github.com/SmartDuck9000/travelly-api/services/full_info_service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type FullInfoDb interface {
	Open() error
	configureConnectionPools() error

	GetHotel(id int) (*Hotel, error)
	GetEvent(id int) (*Event, error)
	GetRestaurant(id int) (*Restaurant, error)
}

type FullInfoPostgres struct {
	url             string
	maxIdleConn     int
	maxOpenConn     int
	connMaxLifetime time.Duration
	conn            *gorm.DB
}

func (db FullInfoPostgres) Open() error {
	var err error
	db.conn, err = gorm.Open(postgres.Open(db.url), &gorm.Config{})
	if err == nil {
		err = db.configureConnectionPools()
	}
	return err
}

func (db FullInfoPostgres) configureConnectionPools() error {
	sqlDB, err := db.conn.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(db.maxIdleConn)
	sqlDB.SetMaxOpenConns(db.maxOpenConn)
	sqlDB.SetConnMaxLifetime(db.connMaxLifetime)

	return nil
}

func CreateFullInfoDB(conf config.FullInfoDbConfig) FullInfoDb {
	return &FullInfoPostgres{
		url:             conf.URL,
		maxIdleConn:     conf.MaxIdleConn,     // maximum number of connections in the idle connection pool
		maxOpenConn:     conf.MaxOpenConn,     // maximum number of open connections to the database
		connMaxLifetime: conf.ConnMaxLifetime, // maximum amount of time a connection may be reused
		conn:            nil,
	}
}

func (db FullInfoPostgres) GetHotel(id int) (*Hotel, error) {
	var hotel Hotel

	res := db.conn.
		Table("hotels").
		Select("hotels.id, hotel_name, hotel_description, hotel_addr, stars, hotel_rating, avg_price, near_sea, country_name, city_name").
		Joins("JOIN cities ON hotels.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("hotels.id = ?", id).Scan(&hotel)

	return &hotel, res.Error
}

func (db FullInfoPostgres) GetEvent(id int) (*Event, error) {
	var event Event

	res := db.conn.
		Table("events").
		Select("events.id, event_name, event_description, event_addr, country_name, city_name, event_start, event_end, event_price, event_rating, max_persons, cur_persons, languages").
		Joins("JOIN cities ON hotels.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("events.id = ?", id).Scan(&event)

	return &event, res.Error
}

func (db FullInfoPostgres) GetRestaurant(id int) (*Restaurant, error) {
	var restaurant Restaurant

	res := db.conn.
		Table("restaurants").
		Select("restaurants.id, rest_name, rest_description, rest_addr, avg_price, rest_rating, child_menu, smoking_room, country_name, city_name").
		Joins("JOIN cities ON hotels.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("restaurants.id = ?", id).Scan(&restaurant)

	return &restaurant, res.Error
}
