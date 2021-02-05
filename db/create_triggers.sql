CREATE EXTENSION plpython3u;

CREATE OR REPLACE FUNCTION update_event_ct_price() RETURNS TRIGGER AS $$
old = TD['old']
new = TD['new']

diff_price = new['event_price'] - old['event_price']
update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])

request = plpy.prepare('SELECT city_tours.id, city_tour_price FROM city_tours ct JOIN city_tours_events cte on ct.id = cte.ct_id JOIN events e ON cte.event_id = e.id WHERE e.id = $1', ['int'])
city_tours = plpy.execute(request, old['id'])

for ct in city_tours:
    plpy.execute(update_request, ct['city_tour_price'] + diff_price, ct['id'])

return TD['new']
$$ language plpython3u;

CREATE OR REPLACE FUNCTION update_restaurant_ct_price() RETURNS TRIGGER AS $$
old = TD['old']
new = TD['new']

diff_price = new['avg_price'] - old['avg_price']
update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])

request = plpy.prepare('SELECT city_tours.id, city_tour_price FROM city_tours ct JOIN city_tours_rest_bookings ctr on ct.id = ctr.ct_id JOIN restaurant_bookings rb ON ctr.rb_id = rb.id WHERE restaurant_id = $1', ['int'])
city_tours = plpy.execute(request, old['id'])

for ct in city_tours:
    plpy.execute(update_request, ct['city_tour_price'] + diff_price, ct['id'])

return TD['new']
$$ language plpython3u;

CREATE OR REPLACE FUNCTION update_hotel_ct_price() RETURNS TRIGGER AS $$
old = TD['old']
new = TD['new']

diff_price = new['avg_price'] - old['avg_price']
update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])

request = plpy.prepare('SELECT city_tours.id, city_tour_price FROM city_tours WHERE hotel_id = $1', ['int'])
city_tours = plpy.execute(request, old['id'])

for ct in city_tours:
    plpy.execute(update_request, ct['city_tour_price'] + diff_price, ct['id'])

return TD['new']
$$ language plpython3u;

CREATE OR REPLACE FUNCTION update_tickets_ct_price() RETURNS TRIGGER AS $$
old = TD['old']
new = TD['new']

diff_price = new['price'] - old['price']
update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])

request = plpy.prepare('SELECT city_tours.id, city_tour_price FROM city_tours WHERE ticket_arrival_id = $1 OR ticket_departure_id = $1', ['int'])
city_tours = plpy.execute(request, old['id'])

for ct in city_tours:
    plpy.execute(update_request, ct['city_tour_price'] + diff_price, ct['id'])

return TD['new']
$$ language plpython3u;

CREATE OR REPLACE FUNCTION update_ct_tour_price() RETURNS TRIGGER AS $$
old = TD['old']
new = TD['new']

diff_price = new['city_tour_price'] - old['city_tour_price']
update_request = plpy.prepare('UPDATE tours SET tour_price = $1 WHERE id = $2', ['int', 'int'])

request = plpy.prepare('SELECT tours.id, tour_price FROM tours JOIN city_tours ct ON tours.id = ct.tour_id WHERE ct.id = $1', ['int'])
tours = plpy.execute(request, old['id'])

for tour in tours:
    plpy.execute(update_request, tour['tour_price'] + diff_price, tour['id'])

return TD['new']
$$ language plpython3u;

CREATE OR REPLACE FUNCTION insert_ct_price() RETURNS TRIGGER AS $$
new = TD['new']
sum_price = 0

ticket_request = plpy.prepare('SELECT price FROM tickets WHERE tickets.id = $1', ['int'])
sum_price += plpy.execute(ticket_request, new['ticket_arrival_id'])[0]['price']
sum_price += plpy.execute(ticket_request, new['ticket_departure_id'])[0]['price']

hotel_request = plpy.prepare('SELECT avg_price FROM hotels WHERE hotel.id = $1', ['int'])
sum_price += plpy.execute(hotel_request, new['hotel_id'])[0]['avg_price']

event_request = plpy.prepare('SELECT event_price FROM city_tours ct JOIN city_tours_events cte on ct.id = cte.ct_id JOIN events e ON cte.event_id = e.id WHERE ct.id = $1', ['int'])
events = plpy.execute(event_request, new['id'])
for event in events:
    sum_price += event['event_price']

rest_request = plpy.prepare('SELECT avg_price FROM city_tours ct JOIN city_tours_rest_bookings ctr on ct.id = ctr.ct_id JOIN restaurant_bookings rb ON ctr.rb_id = rb.id WHERE ct.id = $1', ['int'])
restaurants = plpy.execute(rest_request, new['id'])
for rest in restaurants:
    sum_price += rest['avg_price']

TD['new']['city_tour_price'] = sum_price
return TD['new']
$$ language plpython3u;

CREATE OR REPLACE FUNCTION insert_ct_tour_price() RETURNS TRIGGER AS $$
new = TD['new']
update_request = plpy.prepare('UPDATE tours SET tour_price = $1 WHERE id = $2', ['int', 'int'])

request = plpy.prepare('SELECT tour_price FROM tours WHERE tours.id = $1', ['int'])
tour = plpy.execute(request, new['tour_id'])[0]
plpy.execute(update_request, tour['tour_price'] + new['city_tour_price'], new['tour_id'])

return TD['new']
$$ language plpython3u;