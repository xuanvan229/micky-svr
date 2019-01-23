--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1 (Debian 11.1-1.pgdg90+1)
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
-- Name: mk_post; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mk_post (
    id integer NOT NULL,
    title text,
    description text
);


ALTER TABLE public.mk_post OWNER TO postgres;

--
-- Name: mk_session; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mk_session (
    id integer NOT NULL,
    content text,
    post_id integer
);


ALTER TABLE public.mk_session OWNER TO postgres;

--
-- Name: mk_session_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mk_session_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mk_session_id_seq OWNER TO postgres;

--
-- Name: mk_session_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mk_session_id_seq OWNED BY public.mk_session.id;


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
-- Name: post_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_id_seq OWNER TO postgres;

--
-- Name: post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.post_id_seq OWNED BY public.mk_post.id;


--
-- Name: mk_post id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mk_post ALTER COLUMN id SET DEFAULT nextval('public.post_id_seq'::regclass);


--
-- Name: mk_session id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mk_session ALTER COLUMN id SET DEFAULT nextval('public.mk_session_id_seq'::regclass);


--
-- Name: mk_user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mk_user ALTER COLUMN id SET DEFAULT nextval('public.mk_user_id_seq'::regclass);


--
-- Data for Name: mk_post; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mk_post (id, title, description) FROM stdin;
8	title 1	hongxuan123
9	hongxuan123	new des
10	hongxuan	new des
12	hongxuan1	new des
13	hongxuan12	new des
\.


--
-- Data for Name: mk_session; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mk_session (id, content, post_id) FROM stdin;
2	the data of page	8
3	content 1	9
4	content 2	9
5	content 1	10
6	content 2	10
7	content 1	12
8	content 2	12
9	content 1	13
10	content 2	13
\.


--
-- Data for Name: mk_user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mk_user (id, username, pass) FROM stdin;
9	xuan	$2a$04$ZGH6eJHnGnFMxFh1dCsfquF5uqANF0SI.tARc8l4oZMOTct1h8Urm
\.


--
-- Name: mk_session_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mk_session_id_seq', 10, true);


--
-- Name: mk_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mk_user_id_seq', 9, true);


--
-- Name: post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.post_id_seq', 13, true);


--
-- Name: mk_post mk_post_title_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mk_post
    ADD CONSTRAINT mk_post_title_key UNIQUE (title);


--
-- Name: mk_session mk_session_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mk_session
    ADD CONSTRAINT mk_session_pkey PRIMARY KEY (id);


--
-- Name: mk_user mk_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mk_user
    ADD CONSTRAINT mk_user_pkey PRIMARY KEY (id);


--
-- Name: mk_post post_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mk_post
    ADD CONSTRAINT post_pkey PRIMARY KEY (id);


--
-- Name: mk_session mk_session_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mk_session
    ADD CONSTRAINT mk_session_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.mk_post(id);


--
-- PostgreSQL database dump complete
--

