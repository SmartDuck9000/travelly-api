package db

type TravellyDb interface {
	Open()
	Close()

	GetUser(userId int) *User
	GetTours(userId int) []Tour
	GetCityTours(tourId int) []CityTour

	GetEvents(cityTourId int) []Event
	GetRestaurantBookings(cityTourId int) []RestaurantBooking
	GetTickets(cityTourId int) *Ticket
	GetHotel(cityTourId int) *Hotel
}

type TravellyPostgres struct {
}

func (db TravellyPostgres) Open() {

}

func (db TravellyPostgres) Close() {

}

func (db TravellyPostgres) GetUser(userId int) *User {
	return nil
}

func (db TravellyPostgres) GetTour(tourId int) *Tour {
	return nil
}

func (db TravellyPostgres) GetCityTours(tourId int) []CityTour {
	return nil
}

func (db TravellyPostgres) GetEvents(cityTourId int) []Event {
	return nil
}

func (db TravellyPostgres) GetRestaurantBookings(cityTourId int) []RestaurantBooking {
	return nil
}

func (db TravellyPostgres) GetTicket(ticketId int) *Ticket {
	return nil
}

func (db TravellyPostgres) GetHotel(hotelId int) *Hotel {
	return nil
}
