package model

import (
	"github.com/SmartDuck9000/travelly-api/services/feed_service/config"
	"github.com/SmartDuck9000/travelly-api/services/feed_service/db"
)

type FeedModelInterface interface {
	Run() error

	GetHotels(parameters db.FeedParameters, authHeader string) ([]db.Hotel, error)
	GetEvents(parameters db.FeedParameters, authHeader string) ([]db.Event, error)
	GetRestaurants(parameters db.FeedParameters, authHeader string) ([]db.Restaurant, error)
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

func (model FeedModel) GetHotels(parameters db.FeedParameters, authHeader string) ([]db.Hotel, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	return nil, nil
}

func (model FeedModel) GetEvents(parameters db.FeedParameters, authHeader string) ([]db.Event, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	return nil, nil
}

func (model FeedModel) GetRestaurants(parameters db.FeedParameters, authHeader string) ([]db.Restaurant, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	return nil, nil
}

func (model FeedModel) validateToken(authHeader string) error {
	return nil
}
