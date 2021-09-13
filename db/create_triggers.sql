CREATE OR REPLACE PROCEDURE update_ct_price(ticket_arrival_id int, ticket_departure_id int, hotel_id int, city_tour_id int) AS $$
sum_price = 0
old_price = 0

if ticket_arrival_id != 0:
    ticket_request = plpy.prepare('SELECT price FROM tickets WHERE tickets.id = $1', ['int'])
    ticket_arrival_price = plpy.execute(ticket_request, [ticket_arrival_id])[0]['price']
    sum_price += float(ticket_arrival_price[1:])

if ticket_departure_id != 0:
    ticket_request = plpy.prepare('SELECT price FROM tickets WHERE tickets.id = $1', ['int'])
    ticket_departure_price =  plpy.execute(ticket_request, [ticket_departure_id])[0]['price']
    sum_price += float(ticket_departure_price[1:])

if hotel_id != 0:
    hotel_request = plpy.prepare('SELECT avg_price FROM hotels WHERE hotels.id = $1', ['int'])
    hotel_price = plpy.execute(hotel_request, [hotel_id])[0]['avg_price']
    sum_price += float(hotel_price[1:])

old_price_request = plpy.prepare('SELECT city_tour_price, avg_price, ta.price AS ta_price, td.price AS td_price FROM city_tours ct JOIN hotels h ON ct.hotel_id = h.id JOIN tickets ta ON ta.id = ct.ticket_arrival_id JOIN tickets td ON td.id = ct.ticket_departure_id WHERE ct.id = $1', ['int'])
old_price_res = plpy.execute(old_price_request, [city_tour_id])[0]

new_price = old_price_res['city_tour_price'] - old_price_res['avg_price'] - old_price_res['ta_price'] - old_price_res['td_price'] + sum_price

update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])
plpy.execute(update_request, [int(sum_price), city_tour_id])
plpy.error("update_ct_price.sum_price: " + str(sum_price))

$$ language plpython3u;


CREATE EXTENSION plpython3u;

CREATE OR REPLACE FUNCTION update_event_ct_price() RETURNS TRIGGER AS $$
old = TD['old']
new = TD['new']

diff_price = new['event_price'] - old['event_price']
update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])

request = plpy.prepare('SELECT city_tours.id, city_tour_price FROM city_tours ct JOIN city_tours_events cte on ct.id = cte.ct_id JOIN events e ON cte.event_id = e.id WHERE e.id = $1', ['int'])
city_tours = plpy.execute(request, [old['id']])

for ct in city_tours:
    plpy.execute(update_request, [ct['city_tour_price'] + diff_price, ct['id']])

$$ language plpython3u;

CREATE TRIGGER ct_price_after_event_update
AFTER UPDATE OF event_price ON events
FOR EACH ROW EXECUTE FUNCTION update_event_ct_price();

CREATE OR REPLACE FUNCTION update_restaurant_ct_price() RETURNS TRIGGER AS $$
old = TD['old']
new = TD['new']

diff_price = new['avg_price'] - old['avg_price']
update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])

request = plpy.prepare('SELECT city_tours.id, city_tour_price FROM city_tours ct JOIN city_tours_rest_bookings ctr on ct.id = ctr.ct_id JOIN restaurant_bookings rb ON ctr.rb_id = rb.id WHERE restaurant_id = $1', ['int'])
city_tours = plpy.execute(request, [old['id']])

for ct in city_tours:
    plpy.execute(update_request, [ct['city_tour_price'] + diff_price, ct['id']])

$$ language plpython3u;

CREATE TRIGGER ct_price_after_rest_update
AFTER UPDATE OF avg_price ON restaurants
FOR EACH ROW EXECUTE FUNCTION update_restaurant_ct_price();

CREATE OR REPLACE FUNCTION update_hotel_ct_price() RETURNS TRIGGER AS $$
old = TD['old']
new = TD['new']

if old['id'] != 0:
    diff_price = new['avg_price'] - old['avg_price']
else:
    diff_price = new['avg_price']
update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])

request = plpy.prepare('SELECT city_tours.id, city_tour_price FROM city_tours WHERE hotel_id = $1', ['int'])
city_tours = plpy.execute(request, [old['id']])

for ct in city_tours:
    plpy.execute(update_request, [ct['city_tour_price'] + diff_price, ct['id']])

$$ language plpython3u;

CREATE TRIGGER ct_price_after_hotel_update
AFTER UPDATE OF avg_price ON hotels
FOR EACH ROW EXECUTE FUNCTION update_hotel_ct_price();

CREATE OR REPLACE FUNCTION update_tickets_ct_price() RETURNS TRIGGER AS $$
old = TD['old']
new = TD['new']

if old['id'] != 0:
    diff_price = new['price'] - old['price']
else:
    diff_price = new['price']
update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])

request = plpy.prepare('SELECT city_tours.id, city_tour_price FROM city_tours WHERE ticket_arrival_id = $1 OR ticket_departure_id = $1', ['int'])
city_tours = plpy.execute(request, [old['id']])

for ct in city_tours:
    plpy.execute(update_request, [ct['city_tour_price'] + diff_price, ct['id']])

$$ language plpython3u;

CREATE TRIGGER ct_price_after_ticket_update
AFTER UPDATE OF price ON tickets
FOR EACH ROW EXECUTE FUNCTION update_tickets_ct_price();

CREATE OR REPLACE FUNCTION update_ct_tour_price() RETURNS TRIGGER AS $$
old = TD['old']
new = TD['new']

diff_price = float(new['city_tour_price'][1:]) - float(old['city_tour_price'][1:])
update_request = plpy.prepare('UPDATE tours SET tour_price = $1 WHERE id = $2', ['int', 'int'])

request = plpy.prepare('SELECT tours.id, tour_price FROM tours JOIN city_tours ct ON tours.id = ct.tour_id WHERE ct.id = $1', ['int'])
tours = plpy.execute(request, [old['id']])

for tour in tours:
    plpy.execute(update_request, [int(float(tour['tour_price'][1:]) + diff_price), tour['id']])

$$ language plpython3u;

CREATE TRIGGER tour_price_after_ct_update
AFTER UPDATE OF city_tour_price ON city_tours
FOR EACH ROW EXECUTE FUNCTION update_ct_tour_price();

CREATE OR REPLACE FUNCTION insert_ct_price() RETURNS TRIGGER AS $$
new = TD['new']
sum_price = 0

if new['ticket_arrival_id'] != 0:
    ticket_request = plpy.prepare('SELECT price FROM tickets WHERE tickets.id = $1', ['int'])
    sum_price += plpy.execute(ticket_request, [new['ticket_arrival_id']])[0]['price']
if new['ticket_departure_id'] != 0:
    ticket_request = plpy.prepare('SELECT price FROM tickets WHERE tickets.id = $1', ['int'])
    sum_price += plpy.execute(ticket_request, [new['ticket_departure_id']])[0]['price']

if new['hotel_id'] != 0:
    hotel_request = plpy.prepare('SELECT avg_price FROM hotels WHERE hotels.id = $1', ['int'])
    sum_price += plpy.execute(hotel_request, [new['hotel_id']])[0]['avg_price']

TD['new']['city_tour_price'] = sum_price
$$ language plpython3u;

CREATE TRIGGER ct_price_after_ct_insert
AFTER INSERT ON city_tours
FOR EACH ROW EXECUTE FUNCTION insert_ct_price();

CREATE OR REPLACE FUNCTION insert_ct_events_price() RETURNS TRIGGER AS $$
new = TD['new']
update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])
event_request = plpy.prepare('SELECT city_tour_price, event_price FROM city_tours ct JOIN city_tours_events cte on ct.id = cte.ct_id JOIN events e ON cte.event_id = e.id WHERE ct.id = $1 and e.id = $2', ['int', 'int'])

cte = plpy.execute(event_request, [new['ct_id'], new['event_id']])[0]
price = int(float(cte['city_tour_price'][1:]) + float(cte['event_price'][1:]))
plpy.execute(update_request, [price, new['ct_id']])

$$ language plpython3u;

CREATE TRIGGER ct_price_after_cte_insert
AFTER INSERT ON city_tours_events
FOR EACH ROW EXECUTE FUNCTION insert_ct_events_price();

CREATE OR REPLACE FUNCTION insert_ct_rb_price() RETURNS TRIGGER AS $$
new = TD['new']
update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])
rest_request = plpy.prepare('SELECT city_tour_price, avg_price FROM city_tours ct JOIN city_tours_rest_bookings ctr on ct.id = ctr.ct_id JOIN restaurant_bookings rb ON ctr.rb_id = rb.id JOIN restaurants r ON rb.restaurant_id = r.id WHERE ct.id = $1 AND rb.id = $2', ['int', 'int'])

ct_rb = plpy.execute(rest_request, [new['ct_id'], new['rb_id']])[0]
price = int(float(ct_rb['city_tour_price'][1:]) + float(ct_rb['avg_price'][1:]))
plpy.execute(update_request, [price, new['ct_id']])

$$ language plpython3u;

CREATE TRIGGER ct_price_after_ct_rb_insert
AFTER INSERT ON city_tours_rest_bookings
FOR EACH ROW EXECUTE FUNCTION insert_ct_rb_price();

CREATE OR REPLACE FUNCTION insert_ct_tour_price() RETURNS TRIGGER AS $$
new = TD['new']
update_request = plpy.prepare('UPDATE tours SET tour_price = $1 WHERE id = $2', ['int', 'int'])

request = plpy.prepare('SELECT tour_price FROM tours WHERE tours.id = $1', ['int'])
tour = plpy.execute(request, [new['tour_id']])
tour_price = 0.0
if len(tour) > 0:
    tour_price = float(tour[0]['tour_price'][1:])
new_city_tour_price = float(new['city_tour_price'][1:])
plpy.execute(update_request, [int(tour_price + new_city_tour_price), new['tour_id']])

$$ language plpython3u;

CREATE TRIGGER tour_price_after_ct_insert
AFTER INSERT ON city_tours
FOR EACH ROW EXECUTE FUNCTION insert_ct_tour_price();

CREATE OR REPLACE FUNCTION delete_cte_ct_price() RETURNS TRIGGER AS $$
old = TD['old']

update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])
request = plpy.prepare('SELECT ct.id, city_tour_price FROM city_tours ct JOIN city_tours_events cte on ct.id = cte.ct_id JOIN events e ON cte.event_id = e.id WHERE e.id = $1', ['int'])
city_tours = plpy.execute(request, [old['event_id']])

event_price_request = plpy.prepare('SELECT event_price FROM events WHERE id = $1', ['int'])
event_price = float(plpy.execute(event_price_request, [old['event_id']])[0]['event_price'][1:])

for ct in city_tours:
    price = int(float(ct['city_tour_price'][1:]) - event_price)
    plpy.execute(update_request, [price, ct['id']])

$$ language plpython3u;

CREATE TRIGGER ct_price_before_cte_delete
BEFORE DELETE ON city_tours_events
FOR EACH ROW EXECUTE FUNCTION delete_cte_ct_price();

CREATE OR REPLACE FUNCTION delete_ct_rb_ct_price() RETURNS TRIGGER AS $$
old = TD['old']

update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])
request = plpy.prepare('SELECT ct.id, city_tour_price FROM city_tours ct JOIN city_tours_rest_bookings ctr on ct.id = ctr.ct_id JOIN restaurant_bookings rb ON ctr.rb_id = rb.id WHERE rb.id = $1', ['int'])
city_tours = plpy.execute(request, [old['rb_id']])

rb_price_request = plpy.prepare('SELECT avg_price FROM restaurant_bookings rb JOIN restaurants r ON rb.restaurant_id = r.id WHERE rb.id = $1', ['int'])
rb_price = float(plpy.execute(rb_price_request, [old['rb_id']])[0]['avg_price'][1:])

for ct in city_tours:
    plpy.execute(update_request, [int(float(ct['city_tour_price'][1:]) - rb_price), ct['id']])

$$ language plpython3u;

CREATE TRIGGER ct_price_before_ct_rb_delete
BEFORE DELETE ON city_tours_rest_bookings
FOR EACH ROW EXECUTE FUNCTION delete_ct_rb_ct_price();

CREATE OR REPLACE FUNCTION delete_hotel_ct_price() RETURNS TRIGGER AS $$
old = TD['old']

update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])
request = plpy.prepare('SELECT city_tours.id, city_tour_price FROM city_tours WHERE hotel_id = $1', ['int'])
city_tours = plpy.execute(request, old['id'])

if old['id'] != 0:
    for ct in city_tours:
        plpy.execute(update_request, [ct['city_tour_price'] - old['avg_price'], ct['id']])

$$ language plpython3u;

CREATE TRIGGER ct_price_before_hotel_delete
BEFORE DELETE ON hotels
FOR EACH ROW EXECUTE FUNCTION delete_hotel_ct_price();

CREATE OR REPLACE FUNCTION delete_tickets_ct_price() RETURNS TRIGGER AS $$
old = TD['old']

update_request = plpy.prepare('UPDATE city_tours SET city_tour_price = $1 WHERE id = $2', ['int', 'int'])
request = plpy.prepare('SELECT city_tours.id, city_tour_price FROM city_tours WHERE ticket_arrival_id = $1 OR ticket_departure_id = $1', ['int'])
city_tours = plpy.execute(request, [old['id']])

if old['id'] != 0:
    for ct in city_tours:
        plpy.execute(update_request, [ct['city_tour_price'] - old['price'], ct['id']])

$$ language plpython3u;

CREATE TRIGGER ct_price_before_ticket_delete
BEFORE DELETE ON tickets
FOR EACH ROW EXECUTE FUNCTION delete_tickets_ct_price();