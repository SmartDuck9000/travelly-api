package db

import (
	"fmt"
	"github.com/SmartDuck9000/travelly-api/server/db"
	"github.com/SmartDuck9000/travelly-api/services/feed_service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type FeedDB interface {
	Open() error
	configureConnectionPools() error

	GetHotels(filter db.HotelFilterParameters) ([]Hotel, error)
	GetEvents(filter db.EventsFilterParameters) ([]Event, error)
	GetRestaurants(filter db.RestaurantFilterParameters) ([]Restaurant, error)
	GetTickets(filter db.TicketFilterParameters) ([]Ticket, error)
}

type FeedPostgres struct {
	url             string
	maxIdleConn     int
	maxOpenConn     int
	connMaxLifetime time.Duration
	conn            *gorm.DB
}

func CreateFeedServiceDB(conf config.FeedDBConfig) FeedDB {
	return &FeedPostgres{
		url:             conf.URL,
		maxIdleConn:     conf.MaxIdleConn,     // maximum number of connections in the idle connection pool
		maxOpenConn:     conf.MaxOpenConn,     // maximum number of open connections to the database
		connMaxLifetime: conf.ConnMaxLifetime, // maximum amount of time a connection may be reused
		conn:            nil,
	}
}

func (db *FeedPostgres) Open() error {
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

func (db FeedPostgres) GetHotels(filter db.HotelFilterParameters) ([]Hotel, error) {
	var hotels []Hotel
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

func (db FeedPostgres) GetEvents(filter db.EventsFilterParameters) ([]Event, error) {
	var events []Event
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

func (db FeedPostgres) GetRestaurants(filter db.RestaurantFilterParameters) ([]Restaurant, error) {
	var restaurants []Restaurant
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

func (db FeedPostgres) GetTickets(filter db.TicketFilterParameters) ([]Ticket, error) {
	var tickets []Ticket
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
