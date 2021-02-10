import os
from dotenv import load_dotenv

from termcolor import colored

from db import Postgres


def fill_countries(pg: Postgres):
    pass


def fill_cities(pg: Postgres):
    pass


def fill_transport_stations(pg: Postgres):
    pass


def fill_transport_companies(pg: Postgres):
    pass


def fill_tickets(pg: Postgres):
    pass


def fill_hotels(pg: Postgres):
    pass


def fill_events(pg: Postgres):
    pass


def fill_restaurants(pg: Postgres):
    pass


def fill_rest_bookings(pg: Postgres):
    pass


def fill_users(pg: Postgres):
    pass


def fill_ct_events(pg: Postgres):
    pass


def fill_ct_rb(pg: Postgres):
    pass


def fill_tours(pg: Postgres):
    pass


def fill_city_tours(pg: Postgres):
    pass


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
    fill_countries(db_pg)
    fill_cities(db_pg)
    fill_transport_stations(db_pg)
    fill_transport_companies(db_pg)
    fill_tickets(db_pg)
    fill_hotels(db_pg)
    fill_events(db_pg)
    fill_restaurants(db_pg)
    fill_rest_bookings(db_pg)
    fill_users(db_pg)
    fill_ct_events(db_pg)
    fill_ct_rb(db_pg)
    fill_tours(db_pg)
    fill_city_tours(db_pg)
