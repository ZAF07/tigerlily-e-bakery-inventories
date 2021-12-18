DROP TABLE IF EXISTS skus;
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS discounts;

CREATE TABLE skus (
  id serial,
  sku_id text,
  name text,
  price NUMERIC(18,2),
  type text,
  description text,
  image_url text
);

CREATE TABLE orders (
  id serial,
  order_id text,
  sku_id text,
  customer_id text,
  discount_code text
);

CREATE TABLE customers (
  id serial,
  customer_id text,
  name text,
  email text,
  phone_number text
);

CREATE TABLE carts (
  id serial,
  cart_id text,
  sku_id text,
  customer_id text,
  discount_code text
);

CREATE TABLE discounts (
  id serial,
  discount_id text,
  amount decimal,
  discount_code text,
  description text
);

-- FOR TYPE, I COULD USE PROTO ENUM AND MAP IT TO THE ITEM

INSERT INTO skus(sku_id, name, price, type, description, image_url) VALUES ('11111', 'Egg Tart', 6.8, 'tart', 'Savoury egg tart', 'https://ec2-aws-s3bucket.com');
INSERT INTO skus (sku_id, name, price, type, description, image_url) VALUES ('11222', 'Lemon Cake', 11.5, 'cake', 'Special day Lemon Cake', 'https://ec2-aws-s3bucket.com');
INSERT INTO skus (sku_id, name, price, type, description, image_url) VALUES ('11221', 'Cheese Tart', 7.2, 'tart', 'Creamy Cheese tart', 'https://ec2-aws-s3bucket.com');

INSERT INTO customers (customer_id, name, email, phone_number) VALUES ('00000', 'Zaffere', 'zaf@mail.com', '97898788');
INSERT INTO customers (customer_id, name, email, phone_number) VALUES ('00011', 'Timothy', 'timo@mail.com', '88735546');
INSERT INTO customers (customer_id, name, email, phone_number) VALUES ('00022', 'Mary', 'mary@mail.com', '87934321');

INSERT INTO orders (order_id, sku_id, customer_id, discount_code) VALUES ('33333', '11111', '00000', 'discount1');
INSERT INTO orders (order_id, sku_id, customer_id, discount_code) VALUES ('33444', '11222', '00011', 'discount1');
INSERT INTO orders (order_id, sku_id, customer_id, discount_code) VALUES ('34444', '11221', '00000', 'discount2');

INSERT INTO carts (cart_id, sku_id, customer_id, discount_code) VALUES ('55555', '11221', '00022', 'discount2');

INSERT INTO discounts (discount_id, amount, discount_code, description) VALUES ('99999', 1, 'discount1', 'Happy new year!');
INSERT INTO discounts (discount_id, amount, discount_code, description) VALUES ('99900', .5, 'discount2', 'Christmas Eve');