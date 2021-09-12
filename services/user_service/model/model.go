package model

import (
	"fmt"
	"github.com/SmartDuck9000/travelly-api/services/user_service/config"
	"github.com/SmartDuck9000/travelly-api/services/user_service/db"
	"github.com/SmartDuck9000/travelly-api/token_manager"
)

type UserModelInterface interface {
	Run() error

	GetUser(userId int, authHeader string) (*db.UserData, error)
	GetTours(userId int, authHeader string) ([]db.TourData, error)
	GetCityTours(tourId int, authHeader string) ([]db.CityTourData, error)
	GetCityTourEvents(cityTourId int, authHeader string) ([]db.EventData, error)
	GetCityTourRestaurantBookings(cityTourId int, authHeader string) ([]db.RestaurantBookingData, error)
	GetCityTourTickets(cityTourId int, authHeader string) (*db.CityTourTicketData, error)
	GetCityTourHotel(cityTourId int, authHeader string) (*db.HotelData, error)

	CreateTour(tour *db.Tour, authHeader string) error
	CreateCityTour(cityTour *db.CityTour, authHeader string) error
	CreateCityTourEvent(cityTourEvent *db.CityToursEvent, authHeader string) error
	CreateRestaurantBooking(restaurantBooking *db.RestaurantBookingDTO, authHeader string) error

	UpdateUser(user *UpdateUserData, authHeader string) error
	UpdateTour(tour *db.Tour, authHeader string) error
	UpdateCityTour(cityTour *db.CityTour, authHeader string) error
	UpdateRestaurantBooking(restaurantBooking *db.RestaurantBooking, authHeader string) error

	DeleteUser(userId int, authHeader string) error
	DeleteTour(tourId int, authHeader string) error
	DeleteCityTour(cityTourId int, authHeader string) error
	DeleteRestaurantBooking(restaurantBookingId int, authHeader string) error
}

type UserModel struct {
	db           db.UserProfileDb
	tokenManager token_manager.TokenManager
}

type UpdateUserData struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`

	PhotoUrl string `json:"photo_url"`
}

func CreateUserModel(config config.UserModelConfig) *UserModel {
	return &UserModel{
		db:           db.CreateUserServiceDb(*config.DbConfig),
		tokenManager: token_manager.CreateJWTManager(*config.TokenConfig),
	}
}

func (model UserModel) Run() error {
	return model.db.Open()
}

func (model UserModel) GetUser(userId int, authHeader string) (*db.UserData, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	user := model.db.GetUser(userId)
	if user == nil {
		return user, fmt.Errorf("user with this id doesn't exist")
	}

	return user, nil
}

func (model UserModel) GetTours(userId int, authHeader string) ([]db.TourData, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	tours := model.db.GetTours(userId)
	if tours == nil {
		return tours, fmt.Errorf("user with this id doesn't exist")
	}

	return tours, nil
}

func (model UserModel) GetCityTours(tourId int, authHeader string) ([]db.CityTourData, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	cityTours := model.db.GetCityTours(tourId)
	if cityTours == nil {
		return cityTours, fmt.Errorf("tour with this id doesn't exist")
	}

	return cityTours, nil
}

func (model UserModel) GetCityTourEvents(cityTourId int, authHeader string) ([]db.EventData, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	events := model.db.GetEvents(cityTourId)
	if events == nil {
		return events, fmt.Errorf("city tour with this id doesn't exist")
	}

	return events, nil
}

func (model UserModel) GetCityTourRestaurantBookings(cityTourId int, authHeader string) ([]db.RestaurantBookingData, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	rb := model.db.GetRestaurantBookings(cityTourId)
	if rb == nil {
		return rb, fmt.Errorf("city tour with this id doesn't exist")
	}

	return rb, nil
}

func (model UserModel) GetCityTourTickets(cityTourId int, authHeader string) (*db.CityTourTicketData, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	tickets := model.db.GetTickets(cityTourId)
	if tickets == nil {
		return tickets, fmt.Errorf("city tour with this id doesn't exist")
	}

	return tickets, nil
}

func (model UserModel) GetCityTourHotel(cityTourId int, authHeader string) (*db.HotelData, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	hotel := model.db.GetHotel(cityTourId)
	if hotel == nil {
		return hotel, fmt.Errorf("city tour with this id doesn't exist")
	}

	return hotel, nil
}

func (model UserModel) CreateTour(tour *db.Tour, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.CreateTour(tour)
}

func (model UserModel) CreateCityTour(cityTour *db.CityTour, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.CreateCityTour(cityTour)
}

func (model UserModel) CreateCityTourEvent(cityTourEvent *db.CityToursEvent, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.CreateCityTourEvent(cityTourEvent)
}

func (model UserModel) CreateRestaurantBooking(restaurantBooking *db.RestaurantBookingDTO, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.CreateRestaurantBooking(restaurantBooking)
}

func (model UserModel) UpdateUser(user *UpdateUserData, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}

	curPassword, err := model.db.GetUserPassword(user.Id)
	if err != nil {
		return err
	}

	if curPassword != user.OldPassword {
		return fmt.Errorf("can't update user, wrong password")
	}

	var password = user.OldPassword
	if user.NewPassword != "" {
		password = user.NewPassword
	}

	var userEntity = db.User{
		Id:        user.Id,
		Email:     user.Email,
		Password:  password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		PhotoUrl:  user.PhotoUrl,
	}

	return model.db.UpdateUser(&userEntity)
}

func (model UserModel) UpdateTour(tour *db.Tour, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.UpdateTour(tour)
}

func (model UserModel) UpdateCityTour(cityTour *db.CityTour, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.UpdateCityTour(cityTour)
}

func (model UserModel) UpdateRestaurantBooking(restaurantBooking *db.RestaurantBooking, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.UpdateRestaurantBooking(restaurantBooking)
}

func (model UserModel) DeleteUser(userId int, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.DeleteUser(userId)
}

func (model UserModel) DeleteTour(tourId int, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.DeleteTour(tourId)
}

func (model UserModel) DeleteCityTour(cityTourId int, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.DeleteCityTour(cityTourId)
}

func (model UserModel) DeleteRestaurantBooking(restaurantBookingId int, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.DeleteRestaurantBooking(restaurantBookingId)
}

func (model UserModel) validateToken(authHeader string) error {
	return model.tokenManager.ValidateToken(authHeader)
}
