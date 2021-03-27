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

	CreateUser(user *User) error
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
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PhotoURL  string `json:"photo_url"`
}

func CreateAuthDB(conf config.AuthDbConfig) AuthDB {
	return &AuthPostgres{
		url:             conf.URL,
		maxIdleConn:     conf.MaxIdleConn,     // maximum number of connections in the idle connection pool
		maxOpenConn:     conf.MaxOpenConn,     // maximum number of open connections to the database
		connMaxLifetime: conf.ConnMaxLifetime, // maximum amount of time a connection may be reused
		conn:            nil,
	}
}

func (db *AuthPostgres) Open() error {
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

func (db AuthPostgres) CreateUser(user *User) error {
	res := db.conn.Select("Email", "Password", "FirstName", "LastName").Create(user)
	return res.Error
}

func (db AuthPostgres) GetUser(email string) *User {
	var user User

	db.conn.
		Table("users").
		Select("id, email, password, first_name, last_name, photo_url").
		Where("email = ?", email).Scan(&user)

	return &user
}
