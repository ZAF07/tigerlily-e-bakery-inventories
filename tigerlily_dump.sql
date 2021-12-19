--
-- PostgreSQL database dump
--

-- Dumped from database version 13.4
-- Dumped by pg_dump version 13.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: carts; Type: TABLE; Schema: public; Owner: zaffere
--

CREATE TABLE public.carts (
    id integer NOT NULL,
    cart_id text,
    sku_id text,
    customer_id text,
    discount_code text
);


ALTER TABLE public.carts OWNER TO zaffere;

--
-- Name: carts_id_seq; Type: SEQUENCE; Schema: public; Owner: zaffere
--

CREATE SEQUENCE public.carts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.carts_id_seq OWNER TO zaffere;

--
-- Name: carts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zaffere
--

ALTER SEQUENCE public.carts_id_seq OWNED BY public.carts.id;


--
-- Name: customers; Type: TABLE; Schema: public; Owner: zaffere
--

CREATE TABLE public.customers (
    id integer NOT NULL,
    customer_id text,
    name text,
    email text,
    phone_number text
);


ALTER TABLE public.customers OWNER TO zaffere;

--
-- Name: customers_id_seq; Type: SEQUENCE; Schema: public; Owner: zaffere
--

CREATE SEQUENCE public.customers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.customers_id_seq OWNER TO zaffere;

--
-- Name: customers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zaffere
--

ALTER SEQUENCE public.customers_id_seq OWNED BY public.customers.id;


--
-- Name: discounts; Type: TABLE; Schema: public; Owner: zaffere
--

CREATE TABLE public.discounts (
    id integer NOT NULL,
    discount_id text,
    amount numeric,
    discount_code text,
    description text
);


ALTER TABLE public.discounts OWNER TO zaffere;

--
-- Name: discounts_id_seq; Type: SEQUENCE; Schema: public; Owner: zaffere
--

CREATE SEQUENCE public.discounts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.discounts_id_seq OWNER TO zaffere;

--
-- Name: discounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zaffere
--

ALTER SEQUENCE public.discounts_id_seq OWNED BY public.discounts.id;


--
-- Name: orders; Type: TABLE; Schema: public; Owner: zaffere
--

CREATE TABLE public.orders (
    id integer NOT NULL,
    order_id text,
    sku_id text,
    customer_id text,
    discount_code text
);


ALTER TABLE public.orders OWNER TO zaffere;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: zaffere
--

CREATE SEQUENCE public.orders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.orders_id_seq OWNER TO zaffere;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zaffere
--

ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;


--
-- Name: skus; Type: TABLE; Schema: public; Owner: zaffere
--

CREATE TABLE public.skus (
    id integer NOT NULL,
    sku_id text,
    name text,
    price numeric(18,2),
    type text,
    description text,
    image_url text
);


ALTER TABLE public.skus OWNER TO zaffere;

--
-- Name: skus_id_seq; Type: SEQUENCE; Schema: public; Owner: zaffere
--

CREATE SEQUENCE public.skus_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.skus_id_seq OWNER TO zaffere;

--
-- Name: skus_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zaffere
--

ALTER SEQUENCE public.skus_id_seq OWNED BY public.skus.id;


--
-- Name: carts id; Type: DEFAULT; Schema: public; Owner: zaffere
--

ALTER TABLE ONLY public.carts ALTER COLUMN id SET DEFAULT nextval('public.carts_id_seq'::regclass);


--
-- Name: customers id; Type: DEFAULT; Schema: public; Owner: zaffere
--

ALTER TABLE ONLY public.customers ALTER COLUMN id SET DEFAULT nextval('public.customers_id_seq'::regclass);


--
-- Name: discounts id; Type: DEFAULT; Schema: public; Owner: zaffere
--

ALTER TABLE ONLY public.discounts ALTER COLUMN id SET DEFAULT nextval('public.discounts_id_seq'::regclass);


--
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: zaffere
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: skus id; Type: DEFAULT; Schema: public; Owner: zaffere
--

ALTER TABLE ONLY public.skus ALTER COLUMN id SET DEFAULT nextval('public.skus_id_seq'::regclass);


--
-- Data for Name: carts; Type: TABLE DATA; Schema: public; Owner: zaffere
--

COPY public.carts (id, cart_id, sku_id, customer_id, discount_code) FROM stdin;
1	55555	11221	00022	discount2
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: zaffere
--

COPY public.customers (id, customer_id, name, email, phone_number) FROM stdin;
1	00000	Zaffere	zaf@mail.com	97898788
2	00011	Timothy	timo@mail.com	88735546
3	00022	Mary	mary@mail.com	87934321
\.


--
-- Data for Name: discounts; Type: TABLE DATA; Schema: public; Owner: zaffere
--

COPY public.discounts (id, discount_id, amount, discount_code, description) FROM stdin;
1	99999	1	discount1	Happy new year!
2	99900	0.5	discount2	Christmas Eve
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: zaffere
--

COPY public.orders (id, order_id, sku_id, customer_id, discount_code) FROM stdin;
1	33333	11111	00000	discount1
2	33444	11222	00011	discount1
3	34444	11221	00000	discount2
\.


--
-- Data for Name: skus; Type: TABLE DATA; Schema: public; Owner: zaffere
--

COPY public.skus (id, sku_id, name, price, type, description, image_url) FROM stdin;
1	11111	Egg Tart	6.80	tart	Savoury egg tart	https://ec2-aws-s3bucket.com
2	11222	Lemon Cake	11.50	cake	Special day Lemon Cake	https://ec2-aws-s3bucket.com
3	11221	Cheese Tart	7.20	tart	Creamy Cheese tart	https://ec2-aws-s3bucket.com
\.


--
-- Name: carts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zaffere
--

SELECT pg_catalog.setval('public.carts_id_seq', 1, true);


--
-- Name: customers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zaffere
--

SELECT pg_catalog.setval('public.customers_id_seq', 3, true);


--
-- Name: discounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zaffere
--

SELECT pg_catalog.setval('public.discounts_id_seq', 2, true);


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zaffere
--

SELECT pg_catalog.setval('public.orders_id_seq', 3, true);


--
-- Name: skus_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zaffere
--

SELECT pg_catalog.setval('public.skus_id_seq', 3, true);


--
-- PostgreSQL database dump complete
--

