--
-- PostgreSQL database dump
--

-- Dumped from database version 10.1
-- Dumped by pg_dump version 11.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: mk_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mk_user (
    id integer NOT NULL,
    username text,
    pass text
);


ALTER TABLE public.mk_user OWNER TO postgres;

--
-- Name: mk_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mk_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mk_user_id_seq OWNER TO postgres;

--
-- Name: mk_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mk_user_id_seq OWNED BY public.mk_user.id;


--
-- Name: mk_user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mk_user ALTER COLUMN id SET DEFAULT nextval('public.mk_user_id_seq'::regclass);


--
-- Data for Name: mk_user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mk_user (id, username, pass) FROM stdin;
\.


--
-- Name: mk_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mk_user_id_seq', 1, false);


--
-- Name: mk_user mk_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mk_user
    ADD CONSTRAINT mk_user_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

