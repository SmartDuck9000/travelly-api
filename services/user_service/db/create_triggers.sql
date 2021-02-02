CREATE EXTENSION plpython3u;

CREATE OR REPLACE FUNCTION update_event_ct_price() RETURNS TRIGGER AS $$

$$ language plpython3u;

CREATE OR REPLACE FUNCTION update_restaurant_ct_price() RETURNS TRIGGER AS $$

$$ language plpython3u;

CREATE OR REPLACE FUNCTION update_hotel_ct_price() RETURNS TRIGGER AS $$

$$ language plpython3u;

CREATE OR REPLACE FUNCTION update_tickets_ct_price() RETURNS TRIGGER AS $$

$$ language plpython3u;

CREATE OR REPLACE FUNCTION update_ct_tout_price() RETURNS TRIGGER AS $$

$$ language plpython3u;

CREATE OR REPLACE FUNCTION insert_ct_tout_price() RETURNS TRIGGER AS $$

$$ language plpython3u;

CREATE OR REPLACE FUNCTION insert_tout_tour_price() RETURNS TRIGGER AS $$

$$ language plpython3u;