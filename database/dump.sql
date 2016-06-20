--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.3
-- Dumped by pg_dump version 9.5.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: accounts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE accounts (
    accountid integer NOT NULL,
    clientid integer NOT NULL,
    balance double precision
);


ALTER TABLE accounts OWNER TO postgres;

--
-- Name: accounts_accountid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE accounts_accountid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE accounts_accountid_seq OWNER TO postgres;

--
-- Name: accounts_accountid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE accounts_accountid_seq OWNED BY accounts.accountid;


--
-- Name: accounts_clientid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE accounts_clientid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE accounts_clientid_seq OWNER TO postgres;

--
-- Name: accounts_clientid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE accounts_clientid_seq OWNED BY accounts.clientid;


--
-- Name: cards; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE cards (
    cardid integer NOT NULL,
    accountid integer NOT NULL,
    active boolean DEFAULT false
);


ALTER TABLE cards OWNER TO postgres;

--
-- Name: cards_accountid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE cards_accountid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE cards_accountid_seq OWNER TO postgres;

--
-- Name: cards_accountid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE cards_accountid_seq OWNED BY cards.accountid;


--
-- Name: cards_cardid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE cards_cardid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE cards_cardid_seq OWNER TO postgres;

--
-- Name: cards_cardid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE cards_cardid_seq OWNED BY cards.cardid;


--
-- Name: clientlogin; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE clientlogin (
    id integer NOT NULL,
    username character(20) NOT NULL,
    password character(67) NOT NULL,
    active boolean DEFAULT false NOT NULL
);


ALTER TABLE clientlogin OWNER TO postgres;

--
-- Name: clientlogin_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE clientlogin_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE clientlogin_id_seq OWNER TO postgres;

--
-- Name: clientlogin_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE clientlogin_id_seq OWNED BY clientlogin.id;


--
-- Name: clients; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE clients (
    id integer NOT NULL,
    firstname character(50) NOT NULL,
    lastname character(50) NOT NULL,
    dateofbirth date
);


ALTER TABLE clients OWNER TO postgres;

--
-- Name: clients_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE clients_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE clients_id_seq OWNER TO postgres;

--
-- Name: clients_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE clients_id_seq OWNED BY clients.id;


--
-- Name: loans; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE loans (
    loanid integer NOT NULL,
    clientid integer NOT NULL,
    amount double precision NOT NULL,
    paidamount double precision,
    interest double precision
);


ALTER TABLE loans OWNER TO postgres;

--
-- Name: loans_clientid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE loans_clientid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE loans_clientid_seq OWNER TO postgres;

--
-- Name: loans_clientid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE loans_clientid_seq OWNED BY loans.clientid;


--
-- Name: loans_loanid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE loans_loanid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE loans_loanid_seq OWNER TO postgres;

--
-- Name: loans_loanid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE loans_loanid_seq OWNED BY loans.loanid;


--
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE transactions (
    transactionid integer NOT NULL,
    personid integer NOT NULL,
    clientrequest boolean,
    accountid integer NOT NULL,
    transdate date NOT NULL,
    value double precision
);


ALTER TABLE transactions OWNER TO postgres;

--
-- Name: transactions_accountid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE transactions_accountid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE transactions_accountid_seq OWNER TO postgres;

--
-- Name: transactions_accountid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE transactions_accountid_seq OWNED BY transactions.accountid;


--
-- Name: transactions_transactionid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE transactions_transactionid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE transactions_transactionid_seq OWNER TO postgres;

--
-- Name: transactions_transactionid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE transactions_transactionid_seq OWNED BY transactions.transactionid;


--
-- Name: accountid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY accounts ALTER COLUMN accountid SET DEFAULT nextval('accounts_accountid_seq'::regclass);


--
-- Name: clientid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY accounts ALTER COLUMN clientid SET DEFAULT nextval('accounts_clientid_seq'::regclass);


--
-- Name: cardid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY cards ALTER COLUMN cardid SET DEFAULT nextval('cards_cardid_seq'::regclass);


--
-- Name: accountid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY cards ALTER COLUMN accountid SET DEFAULT nextval('cards_accountid_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY clientlogin ALTER COLUMN id SET DEFAULT nextval('clientlogin_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY clients ALTER COLUMN id SET DEFAULT nextval('clients_id_seq'::regclass);


--
-- Name: loanid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY loans ALTER COLUMN loanid SET DEFAULT nextval('loans_loanid_seq'::regclass);


--
-- Name: clientid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY loans ALTER COLUMN clientid SET DEFAULT nextval('loans_clientid_seq'::regclass);


--
-- Name: transactionid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transactions ALTER COLUMN transactionid SET DEFAULT nextval('transactions_transactionid_seq'::regclass);


--
-- Name: accountid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transactions ALTER COLUMN accountid SET DEFAULT nextval('transactions_accountid_seq'::regclass);


--
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY accounts (accountid, clientid, balance) FROM stdin;
1	1	653.299999999999955
\.


--
-- Name: accounts_accountid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('accounts_accountid_seq', 1, true);


--
-- Name: accounts_clientid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('accounts_clientid_seq', 1, false);


--
-- Data for Name: cards; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY cards (cardid, accountid, active) FROM stdin;
1	1	t
\.


--
-- Name: cards_accountid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('cards_accountid_seq', 1, false);


--
-- Name: cards_cardid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('cards_cardid_seq', 1, true);


--
-- Data for Name: clientlogin; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY clientlogin (id, username, password, active) FROM stdin;
1	ma.kacmar@gmail.com 	$2a$10$o9LdHBvqu0dSVskp.uQCBOBtGXWbu/wU.GHz6IE96X4U/J93ZdKUa       	t
\.


--
-- Name: clientlogin_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('clientlogin_id_seq', 1, true);


--
-- Data for Name: clients; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY clients (id, firstname, lastname, dateofbirth) FROM stdin;
1	Matus                                             	Kacmar                                            	1995-08-23
\.


--
-- Name: clients_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('clients_id_seq', 1, true);


--
-- Data for Name: loans; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY loans (loanid, clientid, amount, paidamount, interest) FROM stdin;
1	1	300	100	10
2	1	900	200	5
\.


--
-- Name: loans_clientid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('loans_clientid_seq', 1, false);


--
-- Name: loans_loanid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('loans_loanid_seq', 2, true);


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY transactions (transactionid, personid, clientrequest, accountid, transdate, value) FROM stdin;
1	3	t	1	2016-06-20	100
2	1	t	1	2016-06-20	100
\.


--
-- Name: transactions_accountid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('transactions_accountid_seq', 1, false);


--
-- Name: transactions_transactionid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('transactions_transactionid_seq', 2, true);


--
-- Name: accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (accountid);


--
-- Name: cards_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY cards
    ADD CONSTRAINT cards_pkey PRIMARY KEY (cardid);


--
-- Name: clientlogin_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY clientlogin
    ADD CONSTRAINT clientlogin_pkey PRIMARY KEY (id);


--
-- Name: clients_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY clients
    ADD CONSTRAINT clients_pkey PRIMARY KEY (id);


--
-- Name: loans_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY loans
    ADD CONSTRAINT loans_pkey PRIMARY KEY (loanid);


--
-- Name: transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (transactionid);


--
-- Name: accounts_clientid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY accounts
    ADD CONSTRAINT accounts_clientid_fkey FOREIGN KEY (clientid) REFERENCES clients(id);


--
-- Name: cards_accountid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY cards
    ADD CONSTRAINT cards_accountid_fkey FOREIGN KEY (accountid) REFERENCES accounts(accountid);


--
-- Name: clientlogin_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY clientlogin
    ADD CONSTRAINT clientlogin_id_fkey FOREIGN KEY (id) REFERENCES clients(id);


--
-- Name: loans_clientid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY loans
    ADD CONSTRAINT loans_clientid_fkey FOREIGN KEY (clientid) REFERENCES clients(id);


--
-- Name: transactions_accountid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transactions
    ADD CONSTRAINT transactions_accountid_fkey FOREIGN KEY (accountid) REFERENCES accounts(accountid);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

