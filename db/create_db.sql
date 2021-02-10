SET ROLE travelly_dba;

CREATE TABLE countries(
    id SERIAL PRIMARY KEY,
    country_name text UNIQUE
);

CREATE TABLE cities(
    id SERIAL PRIMARY KEY,
    country_id int NOT NULL REFERENCES countries(id),
    city_name text
);

CREATE TABLE transport_companies(
    id SERIAL PRIMARY KEY,
    company_name text,
    company_rating numeric(1, 1) CONSTRAINT tc_rating_validation CHECK (1.0 <= company_rating AND company_rating <= 5.0)
);

CREATE TABLE transport_stations(
    id SERIAL PRIMARY KEY,
    city_id int NOT NULL REFERENCES cities(id),
    station_name text,
    station_addr text
);

CREATE TABLE tickets(
    id SERIAL PRIMARY KEY,
    company_id int NOT NULL REFERENCES transport_companies(id),
    orig_station_id int NOT NULL REFERENCES transport_stations(id),
    dest_station_id int NOT NULL REFERENCES transport_stations(id),
    transport_type text,
    price money CONSTRAINT ticket_price_validation CHECK (price::numeric > 0),
    ticket_date date
);

CREATE TABLE hotels(
    id SERIAL PRIMARY KEY,
    city_id int NOT NULL REFERENCES cities(id),
    hotel_name text,
    hotel_description text,
    hotel_addr text,
    stars int CONSTRAINT hotel_stars_validation CHECK (1 <= stars AND stars <= 5),
    hotel_rating numeric(1, 1) CONSTRAINT hotel_rating_validation CHECK (1.0 <= hotel_rating AND hotel_rating <= 5.0),
    avg_price money CONSTRAINT hotel_price_validation CHECK (avg_price::numeric > 0),
    near_sea bool
);

CREATE TABLE events(
    id SERIAL PRIMARY KEY,
    city_id int NOT NULL REFERENCES cities(id),
    event_name text,
    event_description text,
    event_addr text,
    event_start date,
    event_end date,
    event_price money CONSTRAINT event_price_validation CHECK (event_price::numeric > 0),
    max_persons int,
    cur_persons int CONSTRAINT event_person_validation CHECK (cur_persons <= events.max_persons),
    languages text[],
    event_rating numeric(1, 1) CONSTRAINT event_rating_validation CHECK (1.0 <= event_rating AND event_rating <= 5.0)
);

CREATE TABLE restaurants(
    id SERIAL PRIMARY KEY,
    city_id int NOT NULL REFERENCES cities(id),
    rest_name text,
    rest_description text,
    rest_addr text,
    avg_price money CONSTRAINT restaurant_price_validation CHECK (avg_price::numeric > 0),
    rest_rating numeric(1, 1) CONSTRAINT rest_rating_validation CHECK (1.0 <= rest_rating AND rest_rating <= 5.0),
    child_menu bool,
    smoking_room bool
);

CREATE TABLE restaurant_bookings(
    id SERIAL PRIMARY KEY,
    restaurant_id int NOT NULL REFERENCES restaurants(id),
    booking_time date
);

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    email text NOT NULL UNIQUE CONSTRAINT email_validation CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    password text NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    photo_url text
);

CREATE TABLE tours(
    id SERIAL PRIMARY KEY,
    user_id int NOT NULL REFERENCES users(id),
    tour_name text,
    tour_price money CONSTRAINT tour_price_validation CHECK (tour_price::numeric >= 0),
    tour_date_from date,
    tour_date_to date
);

CREATE TABLE city_tours(
    id SERIAL PRIMARY KEY,
    tour_id int NOT NULL REFERENCES tours(id),
    city_id int NOT NULL REFERENCES cities(id),
    city_tour_price money CONSTRAINT ct_price_validation CHECK (city_tour_price::numeric >= 0),
    date_from date,
    date_to date,
    ticket_arrival_id int NOT NULL REFERENCES tickets(id),
    ticket_departure_id int NOT NULL REFERENCES tickets(id),
    hotel_id int NOT NULL REFERENCES hotels(id)
);

CREATE TABLE city_tours_events(
    ct_id int NOT NULL REFERENCES city_tours(id),
    event_id int NOT NULL REFERENCES events(id)
);

CREATE TABLE city_tours_rest_bookings(
    ct_id int NOT NULL REFERENCES city_tours(id),
    rb_id int NOT NULL REFERENCES restaurant_bookings(id)
)