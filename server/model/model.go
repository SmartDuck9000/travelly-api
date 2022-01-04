package model

import (
	"fmt"
	"github.com/SmartDuck9000/travelly-api/server/config"
	"github.com/SmartDuck9000/travelly-api/server/db"
	"github.com/SmartDuck9000/travelly-api/token_manager"
	"strings"
)

type ModelInterface interface {
	Run() error

	RefreshToken(httpHeader string) (*AuthData, error)
	Register(user db.User) (*AuthData, error)
	Login(user db.User) (*AuthData, error)

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

	GetHotels(filter db.HotelFilterParameters, authHeader string) ([]db.HotelFeedItem, error)
	GetEvents(filter db.EventsFilterParameters, authHeader string) ([]db.EventFeedItem, error)
	GetRestaurants(filter db.RestaurantFilterParameters, authHeader string) ([]db.RestaurantFeedItem, error)
	GetTickets(filter db.TicketFilterParameters, authHeader string) ([]db.TicketFeedItem, error)

	GetCities(authHeader string) ([]db.City, error)
	GetHotel(id int, authHeader string) (*db.Hotel, error)
	GetEvent(id int, authHeader string) (*db.Event, error)
	GetRestaurant(id int, authHeader string) (*db.Restaurant, error)
	GetTicket(id int, authHeader string) (*db.Ticket, error)
}

type Model struct {
	db           db.Repository
	tokenManager token_manager.TokenManager
}

type AuthData struct {
	UserId       int    `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
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

func CreateModel(config config.ModelConfig) *Model {
	return &Model{
		db:           db.CreateRepository(*config.DbConfig),
		tokenManager: token_manager.CreateJWTManager(*config.TokenConfig),
	}
}

func (model Model) Run() error {
	return model.db.Open()
}

func (model Model) RefreshToken(authHeader string) (*AuthData, error) {
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return nil, fmt.Errorf("wrong length of header")
	}

	if headerParts[0] != "Bearer" {
		return nil, fmt.Errorf("wrong header")
	}

	claims, err := model.tokenManager.ParseRefreshToken(headerParts[1])
	if err != nil {
		return nil, err
	}

	return model.getAuthData(claims.ID)
}

func (model Model) Register(user db.User) (*AuthData, error) {
	err := model.db.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return model.getAuthData(user.Id)
}

func (model Model) Login(user db.User) (*AuthData, error) {
	userData := model.db.GetUserByEmail(user.Email)
	if userData == nil {
		return nil, fmt.Errorf("user with this email doesn't exist")
	}

	if user.Password != userData.Password {
		return nil, fmt.Errorf("wrong password")
	}

	return model.getAuthData(userData.Id)
}

func (model Model) getAuthData(userId int) (*AuthData, error) {
	accessToken, accessTokenErr := model.tokenManager.CreateAccessToken(userId)
	if accessTokenErr != nil {
		return nil, accessTokenErr
	}

	refreshToken, refreshTokenErr := model.tokenManager.CreateRefreshToken(userId)
	if refreshTokenErr != nil {
		return nil, refreshTokenErr
	}

	return &AuthData{
		UserId:       userId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (model Model) GetUser(userId int, authHeader string) (*db.UserData, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	user := model.db.GetUser(userId)
	if user == nil {
		return user, fmt.Errorf("user with this id doesn't exist")
	}

	return user, nil
}

func (model Model) GetTours(userId int, authHeader string) ([]db.TourData, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	tours := model.db.GetTours(userId)
	if tours == nil {
		return tours, fmt.Errorf("user with this id doesn't exist")
	}

	return tours, nil
}

func (model Model) GetCityTours(tourId int, authHeader string) ([]db.CityTourData, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	cityTours := model.db.GetCityTours(tourId)
	if cityTours == nil {
		return cityTours, fmt.Errorf("tour with this id doesn't exist")
	}

	return cityTours, nil
}

func (model Model) GetCityTourEvents(cityTourId int, authHeader string) ([]db.EventData, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	events := model.db.GetUserEvents(cityTourId)
	if events == nil {
		return events, fmt.Errorf("city tour with this id doesn't exist")
	}

	return events, nil
}

func (model Model) GetCityTourRestaurantBookings(cityTourId int, authHeader string) ([]db.RestaurantBookingData, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	rb := model.db.GetRestaurantBookings(cityTourId)
	if rb == nil {
		return rb, fmt.Errorf("city tour with this id doesn't exist")
	}

	return rb, nil
}

func (model Model) GetCityTourTickets(cityTourId int, authHeader string) (*db.CityTourTicketData, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	tickets := model.db.GetUserTickets(cityTourId)
	if tickets == nil {
		return tickets, fmt.Errorf("city tour with this id doesn't exist")
	}

	return tickets, nil
}

func (model Model) GetCityTourHotel(cityTourId int, authHeader string) (*db.HotelData, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	hotel := model.db.GetUserHotel(cityTourId)
	if hotel == nil {
		return hotel, fmt.Errorf("city tour with this id doesn't exist")
	}

	return hotel, nil
}

func (model Model) CreateTour(tour *db.Tour, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.CreateTour(tour)
}

func (model Model) CreateCityTour(cityTour *db.CityTour, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.CreateCityTour(cityTour)
}

func (model Model) CreateCityTourEvent(cityTourEvent *db.CityToursEvent, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.CreateCityTourEvent(cityTourEvent)
}

func (model Model) CreateRestaurantBooking(restaurantBooking *db.RestaurantBookingDTO, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.CreateRestaurantBooking(restaurantBooking)
}

func (model Model) UpdateUser(user *UpdateUserData, authHeader string) error {
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

func (model Model) UpdateTour(tour *db.Tour, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.UpdateTour(tour)
}

func (model Model) UpdateCityTour(cityTour *db.CityTour, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.UpdateCityTour(cityTour)
}

func (model Model) UpdateRestaurantBooking(restaurantBooking *db.RestaurantBooking, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.UpdateRestaurantBooking(restaurantBooking)
}

func (model Model) DeleteUser(userId int, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.DeleteUser(userId)
}

func (model Model) DeleteTour(tourId int, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.DeleteTour(tourId)
}

func (model Model) DeleteCityTour(cityTourId int, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.DeleteCityTour(cityTourId)
}

func (model Model) DeleteRestaurantBooking(restaurantBookingId int, authHeader string) error {
	if err := model.validateToken(authHeader); err != nil {
		return err
	}
	return model.db.DeleteRestaurantBooking(restaurantBookingId)
}

func (model Model) validateToken(authHeader string) error {
	return model.tokenManager.ValidateToken(authHeader)
}

func (model Model) GetHotels(filter db.HotelFilterParameters, authHeader string) ([]db.HotelFeedItem, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}
	return model.db.GetHotels(filter)
}

func (model Model) GetEvents(filter db.EventsFilterParameters, authHeader string) ([]db.EventFeedItem, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}
	return model.db.GetEvents(filter)
}

func (model Model) GetRestaurants(filter db.RestaurantFilterParameters, authHeader string) ([]db.RestaurantFeedItem, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}
	return model.db.GetRestaurants(filter)
}

func (model Model) GetTickets(filter db.TicketFilterParameters, authHeader string) ([]db.TicketFeedItem, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}
	return model.db.GetTickets(filter)
}

func (model Model) GetCities(authHeader string) ([]db.City, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	return model.db.GetCities()
}

func (model Model) GetHotel(id int, authHeader string) (*db.Hotel, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	return model.db.GetHotel(id)
}

func (model Model) GetEvent(id int, authHeader string) (*db.Event, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	return model.db.GetEvent(id)
}

func (model Model) GetRestaurant(id int, authHeader string) (*db.Restaurant, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	return model.db.GetRestaurant(id)
}

func (model Model) GetTicket(id int, authHeader string) (*db.Ticket, error) {
	if err := model.validateToken(authHeader); err != nil {
		return nil, err
	}

	return model.db.GetTicket(id)
}
