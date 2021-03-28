package model

import (
	"github.com/SmartDuck9000/travelly-api/services/feed_service/config"
	"github.com/SmartDuck9000/travelly-api/services/feed_service/db"
)

type FeedModelInterface interface {
	Run() error

	GetHotels(filter db.HotelFilterParameters, authHeader string) ([]db.Hotel, error)
	GetEvents(filter db.EventsFilterParameters, authHeader string) ([]db.Event, error)
	GetRestaurants(filter db.RestaurantFilterParameters, authHeader string) ([]db.Restaurant, error)
}

type FeedModel struct {
	db db.FeedDB
}

func CreateFeedModel(config config.FeedModelConfig) FeedModelInterface {
	return &FeedModel{
		db: db.CreateFeedServiceDB(*config.DbConfig),
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

func (model FeedModel) validateToken(authHeader string) error {
	return nil
}
