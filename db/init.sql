CREATE DATABASE travelly_db;
CREATE USER smart_duck WITH ENCRYPTED PASSWORD --*password*;
GRANT ALL PRIVILEGES ON DATABASE travelly_db TO smart_duck;

CREATE ROLE travelly_dba WITH SUPERUSER NOINHERIT;
GRANT travelly_dba TO smart_duck;