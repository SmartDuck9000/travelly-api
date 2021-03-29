# travelly-api
API for travelly app(app, that helps people to organize their travel plan)

# Authorization

`api/auth/` returns two tokens `access_token` and `refresh_token`

Two other methods return user id and two tokens

JSON with tokens example:
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIzMzM0NTc3ODkwIiwibmFtZSI6IkdvZ2kiLCJpYXQiOjM1MTYyMzkwMjJ9.pZm2pmR7FoyI0hwfSF_OMuE7tD3MVqeN6-D2UuVSYnQ.eyJzdWIiOiIxMzM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9yaWFuIiwiaWF0IjoxNTE2MjM5MDIyfQ.UCSQHuC44ByGLwA7F5gcYea2rruRlbH6_kXuVv7_6Rg"
}
```

JSON with user id and tokens example:
```json
{
  "user_id": 1,
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIzMzM0NTc3ODkwIiwibmFtZSI6IkdvZ2kiLCJpYXQiOjM1MTYyMzkwMjJ9.pZm2pmR7FoyI0hwfSF_OMuE7tD3MVqeN6-D2UuVSYnQ.eyJzdWIiOiIxMzM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9yaWFuIiwiaWF0IjoxNTE2MjM5MDIyfQ.UCSQHuC44ByGLwA7F5gcYea2rruRlbH6_kXuVv7_6Rg"
}
```

### GET 
- `api/auth/` - use it to refresh access token 
  
input: http header "Authorization" with two header parts: Bearer and refresh token

### POST
- `api/auth/email_register` - registration via email method
  
input: JSON with email, password, first name, last name and optionally photo url

Example with photos url:
```json
{
  "email": "qwerty@gmail.com",
  "password": "awesome_password",
  "first_name": "John",
  "last_name": "Dorian",
  "photo_url": "http://www.google.com/url?sa=i&url=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FHTML%2FElement%2Fimg&psig=AOvVaw2Sx5WUbOxNOXkF4Px38IOk&ust=1616444812929000&source=images&cd=vfe&ved=0CAIQjRxqFwoTCKilgqmcwu8CFQAAAAAdAAAAABAI"
}
```

Example without photos url:
```json
{
  "email": "qwerty@gmail.com",
  "password": "awesome_password",
  "first_name": "John",
  "last_name": "Dorian"
}
```

- `api/auth/login` - login existing user method

input: JSON with email and password

Example:
```json
{
  "email": "qwerty@gmail.com",
  "password": "awesome_password"
}
```

# User profile

### GET
- `/api/users?user_id=` - returns JSON with data about user profile
  
returning JSON example:
```json
{
  "user_id": 1,
  "first_name": "John",
  "last_name":  "Dorian",
  "photo_url": null
}
```

- `/api/users/tours?user_id=` - returns JSON with an array of user's tour data

returning JSON example:
```json
[
  {

    "tour_id": 1,
    "tour_name": "German",
    "tour_price": 1000.0,
    "tour_date_from": "2021-06-01",
    "tour_date_to": "2021-06-15"
  }
]
```

- `/api/users/tours/city_tours?tour_id=` - returns JSON with an array of data about tours in city

returning JSON example:
```json
[
  {
    "city_tour_id": 1,
    "country_name": "German",
    "city_name": "Berlin",
    "city_tour_price": 1000.0,
    "date_from": "2021-06-01",
    "date_to": "2021-06-15",
    "ticket_arrival_id": 25,
    "ticket_departure_id": 67,
    "hotel_name": "Radisson Blu Hotel"
  }
]
```

- `/api/users/tours/city_tours/events?city_tour_id=` - returns JSON with an array of events in city tour

returning JSON example:
```json
[
  {
    "event_id": 1,
    "event_name": "Art festival",
    "event_start": "2021-06-05",
    "event_end": "2021-06-09",
    "price": 50.0,
    "rating": 4.1,
    "max_persons": 600,
    "cur_persons": 230
  }
]
```

- `/api/users/tours/city_tours/restaurant_bookings?city_tour_id=` - returns JSON with an array of restaurant bookings in city tour

returning JSON example:
```json
[
  {
    "restaurant_booking_id": 1,
    "restaurant_id": 1,
    "booking_time": "2021-06-03", 
    "restaurant_name": "Die Eselin von A.",
    "average_price": 65.0,
    "rating": 4.3
  }
]
```

- `/api/users/tours/city_tours/tickets?city_tour_id=` - returns JSON with two structure with tickets

returning JSON example:
```json
{
  "arrival_ticket": {
    "ticket_id": 1,
    "transport_type": "airplane",
    "price": 100.0,
    "date": "2020-06-01",
    "orig_station_name": "Шереметьево",
    "orig_station_addr": "Московская область, Химки, Международное шоссе, 1А",
    "dst_station_name": "Flughafen Berlin Brandenburg",
    "dst_station_addr": "Willy-Brandt-Platz, 12529 Schönefeld",
    "company_name": "S7",
    "company_rating": 4.6
  },
  "departure_ticket": {
    "ticket_id": 2,
    "transport_type": "airplane",
    "price": 100.0,
    "date": "2020-06-15",
    "orig_station_name": "Flughafen Berlin Brandenburg",
    "orig_station_addr": "Willy-Brandt-Platz, 12529 Schönefeld",
    "dst_station_name": "Шереметьево",
    "dst_station_addr": "Московская область, Химки, Международное шоссе, 1А",
    "company_name": "S7",
    "company_rating": 4.6
  }
}
```
- `/api/users/tours/city_tours/hotels?city_tour_id=` - returns JSON with data about hotel in city tour

```json
{
  "hotel_id": 1,
  "hotel_name": "Radisson Blu Hotel",
  "stars": 4,
  "hotel_rating": 4.8
}
```

### POST
- `/api/users/tours` - creates new tour
  
posted JSON example:

```json
{
  "id": 0,
  "user_id": 1,
  "tour_name": "German",
  "tour_price": 0.0,
  "tour_date_from": "2021-06-01",
  "tour_date_to": "2021-06-15"
}
```

- `/api/users/city_tours` - creates new tour in city

posted JSON example:
```json
{
  "id": 0,
  "user_id": 1,
  "city_id": 1,
  "city_tour_price": 1000.0,
  "date_from": "2021-06-01",
  "date_to": "2021-06-15",
  "ticket_arrival_id": 1,
  "ticket_departure_id": 2,
  "hotel_id": 1
}
```

- `/api/users/restaurant_bookings` - creates new booking in restaurant

posted JSON example:
```json
{
  "id": 0,
  "restaurant_id": 1,
  "booking_time": "2021-06-03"
}
```

### PUT
- `/api/users` - updates user info

input JSON example:
```json
{
  "id": 1,
  "email": "qwerty@gmail.com",
  "password": "qwerty12345",
  "first_name": "John",
  "last_name": "Dorian",
  "photo_url": null
}
```

- `/api/users/tours` - updates tour info

input JSON example:
```json
{
  "id": 0,
  "user_id": 1,
  "tour_name": "German",
  "tour_price": 0.0,
  "tour_date_from": "2021-06-01",
  "tour_date_to": "2021-06-15"
}
```

- `/api/users/city_tours` - updates city tour info

input JSON example:
```json
{
  "id": 0,
  "user_id": 1,
  "city_id": 1,
  "city_tour_price": 1000.0,
  "date_from": "2021-06-01",
  "date_to": "2021-06-15",
  "ticket_arrival_id": 1,
  "ticket_departure_id": 2,
  "hotel_id": 1
}
```

- `/api/users/restaurant_bookings` - updates info about booking in restaurant

posted JSON example:
```json
{
  "id": 0,
  "restaurant_id": 1,
  "booking_time": "2021-06-03"
}
```

### DELETE
- `/api/users?user_id=` - deletes user
- `/api/users/tours?tour_id=` - deletes tour
- `/api/users/city_tours?city_tour_id=` - deletes city tour
- `/api/users/restaurant_bookings?restaurant_booking_id=` - deletes booking in restaurant

# Info

### GET
- `/api/info/hotels?id=` - returns JSON with full info about a hotel with given `id`

result JSON example:
```json
{
  "hotel_id": 1,
  "hotel_name": "Radisson Blu Hotel",
  "hotel_description": "Luxury Berlin hotel",
  "hotel_addr": "",
  "stars": "4",
  "hotel_rating": 4.6,
  "average_price": 50.0,
  "near_sea": true,
  "country_name": "German",
  "city_name": "Berlin"
}
```

- `/api/info/events?id=` - returns JSON with full info about an event with given `id`
  
result JSON example:
```json
{
  "event_id": 1,
  "event_name": "Art festival",
  "event_description": "Festival of moder arts",
  "event_addr": "",
  "country_name": "German",
  "city_name": "Berlin",
  "event_start": "2021-06-05",
  "event_end": "2021-06-09",
  "price": 50.0,
  "rating": 4.1,
  "max_persons": 10,
  "cur_persons": 3,
  "languages": ["German", "English"]
}
```

- `/api/info/restaurants?id=` - returns JSON with full info about a restaurant with given `id`

result JSON example:
```json
{
  "restaurant_id": 1,
  "restaurant_name": "Die Eselin von A.",
  "restaurant_description": "", 
  "restaurant_addr": "",
  "average_price": 65.0, 
  "rating": 4.3,
  "child_menu": true,
  "smoking_room": false,
  "country_name": "German",
  "city_name": "Berlin"
}
```

# Feed

order_type should be "inc" for increase sorting or "dec" for decrease sorting

### GET
- `/api/feed/hotels` - returns array of hotel data, gets JSON with filter data

if near_sea is false, method will return hotels with both true and false values of near_sea field

filter data JSON example:
```json
{
  "limit": 20,
  "offset": 10,
  "order_by": "hotel_name",
  "order_type": "inc",
  "hotel_name": "Radisson Blu Hotel",
  "stars_from": 3,
  "stars_to": 4,
  "rating_from": 4.0,
  "rating_to": 5.0,
  "price_from": 0.0,
  "price_to": 50.0,
  "near_sea": false,
  "city_name": "Berlin"
}
```

result JSON example:
```json
[
  {
    "hotel_id": 1,
    "hotel_name": "Radisson Blu Hotel",
    "stars": 4,
    "hotel_rating": 4.6,
    "country_name": "German",
    "city_name": "Berlin"
  },
  {
    "hotel_id": 2,
    "hotel_name": "Berlin Hotel",
    "stars": 3,
    "hotel_rating": 3.2,
    "country_name": "German",
    "city_name": "Berlin"
  }
]
```

- `/api/feed/events` - returns array of event data, gets JSON with filter data

filter data JSON example:
```json
{
  "limit": 20,
  "offset": 10,
  "order_by": "event_name",
  "order_type": "inc",
  "event_name": "Art",
  "from": "2021-06-01",
  "to": "2021-06-01",
  "rating_from": 4.0,
  "rating_to": 5.0,
  "price_from": 0.0,
  "price_to": 50.0,
  "city_name": "Berlin"
}
```

result JSON example:

```json
[
  {
    "event_id": 1,
    "event_name": "Art festival",
    "event_start": "2021-06-05",
    "event_end": "2021-06-06",
    "rating": 4.0,
    "max_persons": 10,
    "cur_persons": 4,
    "country_name": "German",
    "city_name": "Berlin"
  },
  {
    "event_id": 2,
    "event_name": "Food festival",
    "event_start": "2021-06-03",
    "event_end": "2021-06-10",
    "rating": 4.2,
    "max_persons": 20,
    "cur_persons": 12,
    "country_name": "German",
    "city_name": "Berlin"
  }
]
```

- `/api/feed/restaurants` - returns array of restaurant data, gets JSON with filter data

if child_menu is false, method will return hotels with both true and false values of child_menu field

if smoking_room is false, method will return hotels with both true and false values of smoking_room field

filter data JSON example:
```json
{
  "limit": 20,
  "offset": 10,
  "order_by": "restaurant_name",
  "order_type": "inc",
  "restaurant_name": "",
  "rating_from": 4.0,
  "rating_to": 5.0,
  "price_from": 0.0,
  "price_to": 50.0,
  "child_menu": true,
  "smoking_room": false,
  "city_name": "Berlin"
}
```

result JSON example:

```json
[
  {
    "restaurant_id": 1,
    "restaurantName": "Die Eselin von A.",
    "rating": 4.2,
    "country_name": "German",
    "city_name": "Berlin"
  },
  {
    "restaurant_id": 2,
    "restaurantName": "Berlin central restaurant",
    "rating": 3.8,
    "country_name": "German",
    "city_name": "Berlin"
  }
]
```