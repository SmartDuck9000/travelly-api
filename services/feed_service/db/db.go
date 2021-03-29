package db

import (
	"fmt"
	"github.com/SmartDuck9000/travelly-api/services/feed_service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type FeedDB interface {
	Open() error
	configureConnectionPools() error

	GetHotels(filter HotelFilterParameters) ([]Hotel, error)
	GetEvents(filter EventsFilterParameters) ([]Event, error)
	GetRestaurants(filter RestaurantFilterParameters) ([]Restaurant, error)
	GetTickets(filter TicketFilterParameters) ([]Ticket, error)
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

func (db FeedPostgres) GetHotels(filter HotelFilterParameters) ([]Hotel, error) {
	var hotels []Hotel
	var order string

	if filter.orderType == "inc" {
		order = filter.orderBy + " " + "ASC"
	} else if filter.orderType == "dec" {
		order = filter.orderBy + " " + "DESC"
	} else {
		return nil, fmt.Errorf("wrong order type")
	}

	res := db.conn.
		Table("hotels").
		Select("hotels.id, hotel_name, stars, hotel_rating, country_name, city_name").
		Joins("JOIN cities ON hotels.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("stars BETWEEN ? AND ?", filter.starsFrom, filter.starsTo).
		Where("hotel_rating BETWEEN ? AND ?", filter.ratingFrom, filter.ratingTo).
		Where("avg_price BETWEEN ? AND ?", filter.priceFrom, filter.priceTo)

	if filter.hotelName != "" {
		res = res.Where("hotel_name LIKE ?", filter.hotelName)
	}

	if filter.cityName != "" {
		res = res.Where("city_name LIKE ?", filter.cityName)
	}

	if filter.nearSea {
		res = res.Where("near_sea = ?", filter.nearSea)
	}

	res = res.Order(order).Offset(filter.offset).Limit(filter.limit).Scan(&hotels)

	return hotels, res.Error
}

func (db FeedPostgres) GetEvents(filter EventsFilterParameters) ([]Event, error) {
	var events []Event
	var order string

	if filter.orderType == "inc" {
		order = filter.orderBy + " " + "ASC"
	} else if filter.orderType == "dec" {
		order = filter.orderBy + " " + "DESC"
	} else {
		return nil, fmt.Errorf("wrong order type")
	}

	res := db.conn.
		Table("events").
		Select("events.id, event_name, event_start, event_end, event_rating, max_persons, cur_persons, country_name, city_name").
		Joins("JOIN cities ON hotels.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("event_start <= ? AND event_end <= ?", filter.from, filter.to).
		Where("event_rating BETWEEN ? AND ?", filter.ratingFrom, filter.ratingTo).
		Where("event_price BETWEEN ? AND ?", filter.priceFrom, filter.priceTo)

	if filter.eventName != "" {
		res = res.Where("event_name LIKE ?", filter.eventName)
	}

	if filter.cityName != "" {
		res = res.Where("city_name LIKE ?", filter.cityName)
	}

	res = res.Order(order).Offset(filter.offset).Limit(filter.limit).Scan(&events)

	return events, res.Error
}

func (db FeedPostgres) GetRestaurants(filter RestaurantFilterParameters) ([]Restaurant, error) {
	var restaurants []Restaurant
	var order string

	if filter.orderType == "inc" {
		order = filter.orderBy + " " + "ASC"
	} else if filter.orderType == "dec" {
		order = filter.orderBy + " " + "DESC"
	} else {
		return nil, fmt.Errorf("wrong order type")
	}

	res := db.conn.
		Table("restaurants").
		Select("restaurants.id, rest_name, rest_rating, country_name, city_name").
		Joins("JOIN cities ON hotels.city_id = cities.id").
		Joins("JOIN countries ON cities.country_id = countries.id").
		Where("rest_rating BETWEEN ? AND ?", filter.ratingFrom, filter.ratingTo).
		Where("avg_price BETWEEN ? AND ?", filter.priceFrom, filter.priceTo)

	if filter.restaurantName != "" {
		res = res.Where("rest_name LIKE ?", filter.restaurantName)
	}

	if filter.cityName != "" {
		res = res.Where("city_name LIKE ?", filter.cityName)
	}

	if filter.childMenu {
		res = res.Where("child_menu = ?", filter.childMenu)
	}

	if filter.smokingRoom {
		res = res.Where("smoking_room = ?", filter.smokingRoom)
	}

	res = res.Order(order).Offset(filter.offset).Limit(filter.limit).Scan(&restaurants)

	return restaurants, res.Error
}

func (db FeedPostgres) GetTickets(filter TicketFilterParameters) ([]Ticket, error) {
	var tickets []Ticket
	var order string

	if filter.orderType == "inc" {
		order = filter.orderBy + " " + "ASC"
	} else if filter.orderType == "dec" {
		order = filter.orderBy + " " + "DESC"
	} else {
		return nil, fmt.Errorf("wrong order type")
	}

	res := db.conn.
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
		Where("ticket_date BETWEEN ? AND ?", filter.dateFrom, filter.dateTo).
		Where("price BETWEEN ? AND ?", filter.priceFrom, filter.priceTo)

	if filter.transportType != "" {
		res = res.Where("transport_type LIKE ?", filter.transportType)
	}

	if filter.origCityName != "" {
		res = res.Where("orig_city.city_name LIKE ?", filter.origCityName)
	}

	if filter.destCityName != "" {
		res = res.Where("dest_city.city_name LIKE ?", filter.destCityName)
	}

	res = res.Order(order).Offset(filter.offset).Limit(filter.limit).Scan(&tickets)

	return tickets, res.Error
}
