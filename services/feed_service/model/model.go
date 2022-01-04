package model

import (
	db2 "github.com/SmartDuck9000/travelly-api/server/db"
	"github.com/SmartDuck9000/travelly-api/services/feed_service/config"
	"github.com/SmartDuck9000/travelly-api/services/feed_service/db"
	"github.com/SmartDuck9000/travelly-api/token_manager"
)

type FeedModelInterface interface {
	Run() error

	GetHotels(filter db2.HotelFilterParameters, authHeader string) ([]db.Hotel, error)
	GetEvents(filter db2.EventsFilterParameters, authHeader string) ([]db.Event, error)
	GetRestaurants(filter db2.RestaurantFilterParameters, authHeader string) ([]db.Restaurant, error)
	GetTickets(filter db2.TicketFilterParameters, authHeader string) ([]db.Ticket, error)
}

type FeedModel struct {
	db           db.FeedDB
	tokenManager token_manager.TokenManager
}

func CreateFeedModel(config config.FeedModelConfig) FeedModelInterface {
	return &FeedModel{
		db:           db.CreateFeedServiceDB(*config.DbConfig),
		tokenManager: token_manager.CreateJWTManager(*config.TokenConfig),
	}
}

func (model FeedModel) Run() error {
	return model.db.Open()
}

func (model FeedModel) GetHotels(filter db.HotelFilterParameters, authHeader string) ([]db.Hotel, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}
	return model.db.GetHotels(filter)
}

func (model FeedModel) GetEvents(filter db.EventsFilterParameters, authHeader string) ([]db.Event, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}
	return model.db.GetEvents(filter)
}

func (model FeedModel) GetRestaurants(filter db.RestaurantFilterParameters, authHeader string) ([]db.Restaurant, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}
	return model.db.GetRestaurants(filter)
}

func (model FeedModel) GetTickets(filter db.TicketFilterParameters, authHeader string) ([]db.Ticket, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}
	return model.db.GetTickets(filter)
}

func (model FeedModel) validateToken(authHeader string) error {
	return model.tokenManager.ValidateToken(authHeader)
}
