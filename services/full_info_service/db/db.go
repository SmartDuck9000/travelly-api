package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type FullInfoDb interface {
	Open() error
	configureConnectionPools() error

	GetHotel(id int) *Hotel
	GetEvent(id int) *Event
	GetRestaurant(id int) *Restaurant
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
