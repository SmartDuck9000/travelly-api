package model

import (
	"github.com/SmartDuck9000/travelly-api/services/full_info_service/config"
	"github.com/SmartDuck9000/travelly-api/services/full_info_service/db"
	"github.com/SmartDuck9000/travelly-api/token_manager"
)

type FullInfoModelInterface interface {
	Run() error

	GetHotel(id int, authHeader string) (*db.Hotel, error)
	GetEvent(id int, authHeader string) (*db.Event, error)
	GetRestaurant(id int, authHeader string) (*db.Restaurant, error)
	GetTicket(id int, authHeader string) (*db.Ticket, error)
}

type FullInfoModel struct {
	db           db.FullInfoDb
	tokenManager token_manager.TokenManager
}

func CreateFullInfoModel(config config.FullInfoModelConfig) FullInfoModelInterface {
	return &FullInfoModel{
		db:           db.CreateFullInfoDB(*config.Db),
		tokenManager: token_manager.CreateJWTManager(*config.TokenConfig),
	}
}

func (model FullInfoModel) Run() error {
	return model.db.Open()
}

func (model FullInfoModel) GetHotel(id int, authHeader string) (*db.Hotel, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	return model.db.GetHotel(id)
}

func (model FullInfoModel) GetEvent(id int, authHeader string) (*db.Event, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	return model.db.GetEvent(id)
}

func (model FullInfoModel) GetRestaurant(id int, authHeader string) (*db.Restaurant, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	return model.db.GetRestaurant(id)
}

func (model FullInfoModel) GetTicket(id int, authHeader string) (*db.Ticket, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	return model.db.GetTicket(id)
}

func (model FullInfoModel) validateToken(authHeader string) error {
	return model.tokenManager.ValidateToken(authHeader)
}
