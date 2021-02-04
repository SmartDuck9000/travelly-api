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
