DROP TABLE IF EXISTS skus;
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS discounts;

CREATE TABLE skus (
  id serial PRIMARY KEY,
  sku_id character varying NOT NULL,
  name character varying NOT NULL,
  price NUMERIC(18,2),
  quantity INTEGER,
  type character varying NOT NULL,
  description character varying NOT NULL,
  image_url character varying NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE,
  updated_at TIMESTAMP WITHOUT TIME ZONE,
  deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE orders (
  id serial PRIMARY KEY,
  order_id character varying NOT NULL,
  sku_id character varying NOT NULL,
  customer_id character varying NOT NULL,
  discount_code character varying NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE,
  updated_at TIMESTAMP WITHOUT TIME ZONE,
  deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE customers (
  id serial PRIMARY KEY,
  customer_id character varying NOT NULL,
  name character varying NOT NULL,
  email character varying NOT NULL,
  phone_number character varying NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE,
  updated_at TIMESTAMP WITHOUT TIME ZONE,
  deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE carts (
  id serial PRIMARY KEY,
  cart_id character varying NOT NULL,
  sku_id character varying NOT NULL,
  customer_id character varying NOT NULL,
  discount_code character varying NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE,
  updated_at TIMESTAMP WITHOUT TIME ZONE,
  deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE discounts (
  id serial PRIMARY KEY,
  discount_id character varying NOT NULL,
  amount decimal,
  discount_code character varying NOT NULL,
  description character varying NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE,
  updated_at TIMESTAMP WITHOUT TIME ZONE,
  deleted_at TIMESTAMP WITHOUT TIME ZONE
);

-- FOR TYPE, I COULD USE PROTO ENUM AND MAP IT TO THE ITEM

INSERT INTO skus(sku_id, name, price, type, description, image_url, quantity, created_at, updated_at) VALUES ('11111', 'Egg Tart', 6.8, 'tart', 'Savoury egg tart', 'egg', 3, NOW(), NOW());
INSERT INTO skus (sku_id, name, price, type, description, image_url, quantity, created_at, updated_at) VALUES ('11222', 'Lemon Cake', 11.5, 'cake', 'Special day Lemon Cake', 'lemon', 3, NOW(), NOW());
INSERT INTO skus (sku_id, name, price, type, description, image_url, quantity, created_at, updated_at) VALUES ('11221', 'Cheese Tart', 7.2, 'tart', 'Creamy Cheese tart', 'cheese', 3, NOW(), NOW());

INSERT INTO customers (customer_id, name, email, phone_number, created_at, updated_at) VALUES ('00000', 'Zaffere', 'zaf@mail.com', '97898788', NOW(), NOW());
INSERT INTO customers (customer_id, name, email, phone_number, created_at, updated_at) VALUES ('00011', 'Timothy', 'timo@mail.com', '88735546', NOW(), NOW());
INSERT INTO customers (customer_id, name, email, phone_number, created_at, updated_at) VALUES ('00022', 'Mary', 'mary@mail.com', '87934321', NOW(), NOW());

INSERT INTO orders (order_id, sku_id, customer_id, discount_code, created_at, updated_at) VALUES ('33333', '11111', '00000', 'discount1', NOW(), NOW());
INSERT INTO orders (order_id, sku_id, customer_id, discount_code, created_at, updated_at) VALUES ('33444', '11222', '00011', 'discount1', NOW(), NOW());
INSERT INTO orders (order_id, sku_id, customer_id, discount_code, created_at, updated_at) VALUES ('34444', '11221', '00000', 'discount2', NOW(), NOW());

INSERT INTO carts (cart_id, sku_id, customer_id, discount_code, created_at, updated_at) VALUES ('55555', '11221', '00022', 'discount2', NOW(), NOW());

INSERT INTO discounts (discount_id, amount, discount_code, description, created_at, updated_at) VALUES ('99999', 1, 'discount1', 'Happy new year!', NOW(), NOW());
INSERT INTO discounts (discount_id, amount, discount_code, description, created_at, updated_at) VALUES ('99900', .5, 'discount2', 'Christmas Eve', NOW(), NOW());