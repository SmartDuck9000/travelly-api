package db

import (
	"github.com/SmartDuck9000/travelly-api/services/auth_service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type AuthDB interface {
	Open() error
	configureConnectionPools() error

	CreateUser(user User) *User
	GetUser(email string) *User
}

type AuthPostgres struct {
	url             string
	maxIdleConn     int
	maxOpenConn     int
	connMaxLifetime time.Duration
	conn            *gorm.DB
}

type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
}

func CreateAuthDB(conf config.AuthDBConfig) *AuthPostgres {
	return &AuthPostgres{
		url:             conf.URL,
		maxIdleConn:     conf.MaxIdleConn,     // maximum number of connections in the idle connection pool
		maxOpenConn:     conf.MaxOpenConn,     // maximum number of open connections to the database
		connMaxLifetime: conf.ConnMaxLifetime, // maximum amount of time a connection may be reused
		conn:            nil,
	}
}

func (db AuthPostgres) Open() error {
	var err error
	db.conn, err = gorm.Open(postgres.Open(db.url), &gorm.Config{})
	if err == nil {
		err = db.configureConnectionPools()
	}
	return err
}

func (db AuthPostgres) configureConnectionPools() error {
	sqlDB, err := db.conn.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(db.maxIdleConn)
	sqlDB.SetMaxOpenConns(db.maxOpenConn)
	sqlDB.SetConnMaxLifetime(db.connMaxLifetime)

	return nil
}

func (db AuthPostgres) CreateUser(user User) *User {
	var userData User
	return &userData
}

func (db AuthPostgres) GetUser(email string) *User {
	var user User
	return &user
}
