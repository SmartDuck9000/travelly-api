import csv
import random
import os

from dotenv import load_dotenv
from termcolor import colored
from faker import Faker
from password_generator import PasswordGenerator

from db import Postgres


def fill_countries_cities(pg: Postgres):
    with open('data/country_cities.csv') as csv_file:
        reader = csv.DictReader(csv_file, delimiter=',')
        countries = dict()

        for row in reader:
            country_name = row['country']
            city_name = row['city_ascii']

            if country_name not in countries.keys():
                countries[country_name] = pg.insert('countries', {
                    'country_name': country_name
                })[0]['id']

            pg.insert('cities', {
                'country_id': countries[country_name],
                'city_name': city_name
            })


def fill_transport_stations(pg: Postgres):
    with open('data/airports.csv') as csv_file:
        reader = csv.DictReader(csv_file, delimiter=',')
        fake = Faker()

        for row in reader:
            country_name = row['country']
            city_name = row['city']
            station_name = row['name']

            country_ids = pg.get_country_id(country_name)
            if len(country_ids) == 0:
                country_id = pg.insert('countries', {
                    'country_name': country_name
                })[0]['id']
            else:
                country_id = country_ids[0]['id']

            city_ids = pg.get_city_id(city_name)
            if len(city_ids) == 0:
                city_id = pg.insert('cities', {
                    'country_id': country_id,
                    'city_name': city_name
                })[0]['id']
            else:
                city_id = city_ids[0]['id']

            pg.insert('transport_stations', {
                'city_id': city_id,
                'station_name': station_name,
                'station_addr': fake.address().split('\n')[0]
            })


def fill_transport_companies(pg: Postgres):
    with open('data/airlines.csv') as csv_file:
        reader = csv.DictReader(csv_file, delimiter=',')

        for row in reader:
            if row['Active'] == 'Y':
                pg.insert('transport_companies', {
                    'company_name': row['Name'],
                    'company_rating': round(random.uniform(2, 5), 1)
                })


def fill_tickets(pg: Postgres):
    fake = Faker()
    for i in range(1000):
        date = fake.date_between(start_date='+90d', end_date='+1y')
        date_str = str(date.year) + "-" + str(date.month) + "-" + str(date.day)
        orig_station_id = random.randint(1, 6341)
        dest_station_id = random.randint(1, 6341)

        while dest_station_id == orig_station_id:
            dest_station_id = random.randint(1, 6341)

        pg.insert('tickets', {
            'company_id': random.randint(1, 1254),
            'orig_station_id': orig_station_id,
            'dest_station_id': dest_station_id,
            'transport_type': "aircraft",
            'price': random.randrange(200, 500),
            'ticket_date': date_str
        })


def fill_hotels(pg: Postgres):
    pass


def fill_events(pg: Postgres):
    pass


def fill_restaurants(pg: Postgres):
    with open('data/restaurants.csv') as csv_file:
        reader = csv.DictReader(csv_file, delimiter=',')
        fake = Faker()

        for row in reader:
            country = row['region']
            city = row['city']
            rest = row['name']

            country_ids = pg.get_country_id(country)
            if len(country_ids) == 0:
                country_id = pg.insert('countries', {
                    'country_name': country
                })[0]['id']
            else:
                country_id = country_ids[0]['id']

            city_ids = pg.get_city_id(city)
            if len(city_ids) == 0:
                city_id = pg.insert('cities', {
                    'country_id': country_id,
                    'city_name': city
                })[0]['id']
            else:
                city_id = city_ids[0]['id']

            pg.insert('restaurants', {
                'city_id': city_id,
                'rest_name': rest,
                'rest_addr': fake.address().split('\n')[0],
                'avg_price': len(row['price']) * 50,
                'rest_rating': round(random.uniform(3, 5), 1),
                'child_menu': bool(random.randint(0, 1)),
                'smoking_room': bool(random.randint(0, 1))
            })


def fill_users(pg: Postgres):
    fake = Faker()
    pw_generator = PasswordGenerator()

    for i in range(1000):
        pg.insert('users', {
            'email': fake.ascii_free_email(),
            'password': pw_generator.shuffle_password('qwertyuioplkjhgfdsazxcvbnm0123456789', 20),
            'first_name': fake.first_name(),
            'last_name': fake.last_name(),
            'photo_url': fake.image_url()
        })


def init_db(config_file):
    try:
        dotenv_path = os.path.join(os.path.dirname(__file__), config_file)
        if os.path.exists(dotenv_path):
            load_dotenv(dotenv_path)

        db = Postgres(
            db_name=os.environ.get('DB_NAME'),
            db_username=os.environ.get('DB_USER'),
            db_password=os.environ.get('DB_PASSWORD'),
            db_host=os.environ.get('DB_HOST'),
            db_port=os.environ.get('DB_PORT')
        )
    except Exception as e:
        print(colored(e, 'red'))
        return None

    return db


if __name__ == '__main__':
    db_pg = init_db('.env')
    # fill_countries_cities(db_pg)
    # fill_transport_stations(db_pg)
    # fill_transport_companies(db_pg)
    # fill_tickets(db_pg)
    fill_hotels(db_pg)
    fill_events(db_pg)
    fill_restaurants(db_pg)
    # fill_users(db_pg)
