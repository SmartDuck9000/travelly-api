package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type FeedDb interface {
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
