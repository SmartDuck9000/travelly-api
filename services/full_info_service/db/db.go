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

	GetCities() ([]City, error)
	GetHotel(id int) (*Hotel, error)
	GetEvent(id int) (*Event, error)
	GetRestaurant(id int) (*Restaurant, error)
	GetTicket(id int) (*Ticket, error)
}

type FullInfoPostgres struct {
	url             string
	maxIdleConn     int
	maxOpenConn     int
	connMaxLifetime time.Duration
	conn            *gorm.DB
}

func (db *FullInfoPostgres) Open() error {
	var err error
	db.conn, err = gorm.Open(postgres.Open(db.url), &gorm.Config{})
	if err == nil {
		err = db.configureConnectionPools()
	}
	return err
}

func (db *FullInfoPostgres) configureConnectionPools() error {
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

func (db FullInfoPostgres) GetCities() ([]City, error) {
	var cities []City

	res := db.conn.
		Table("cities").
		Select("id AS city_id, city_name").
		Order("city_name").
		Scan(&cities)

	return cities, res.Error
}

func (db FullInfoPostgres) GetHotel(id int) (*Hotel, error) {
	var hotel Hotel

	res := db.conn.
		Table("hotels").
		Select("hotels.id AS hotel_id, hotel_name, hotel_description, hotel_addr, stars, hotel_rating, avg_price, near_sea, country_name, city_name").
		Joins("JOIN cities ON hotels.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("hotels.id = ?", id).Scan(&hotel)

	return &hotel, res.Error
}

func (db FullInfoPostgres) GetEvent(id int) (*Event, error) {
	var event Event

	res := db.conn.
		Table("events").
		Select("events.id AS event_id, event_name, event_description, event_addr, country_name, city_name, event_start, event_end, event_price AS price, event_rating AS rating, max_persons, cur_persons, languages").
		Joins("JOIN cities ON events.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("events.id = ?", id).Scan(&event)

	return &event, res.Error
}

func (db FullInfoPostgres) GetRestaurant(id int) (*Restaurant, error) {
	var restaurant Restaurant

	res := db.conn.
		Table("restaurants").
		Select("restaurants.id AS restaurant_id, rest_name, rest_description, rest_addr, avg_price, rest_rating, child_menu, smoking_room, country_name, city_name").
		Joins("JOIN cities ON restaurants.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("restaurants.id = ?", id).Scan(&restaurant)

	return &restaurant, res.Error
}

func (db FullInfoPostgres) GetTicket(id int) (*Ticket, error) {
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
