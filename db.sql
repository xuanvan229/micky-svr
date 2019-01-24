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
-- Name: mk_post_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mk_post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mk_post_id_seq OWNER TO postgres;

--
-- Name: mk_post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mk_post_id_seq OWNED BY public.mk_post.id;


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
-- Name: mk_post id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mk_post ALTER COLUMN id SET DEFAULT nextval('public.mk_post_id_seq'::regclass);


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
1	title12243 3333	new descri22222ption change
2	title12243 3333	new descri22222ption change
\.


--
-- Data for Name: mk_session; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mk_session (id, content, post_id) FROM stdin;
1	conte2232323nt 11111123123123	1
2	conte2323233nt 11111312312321	1
3	conte2323233nt 11111312312321	1
4	conte2232323nt 11111123123123	1
5	conte2323233nt 11111312312321	1
6	conte2323233nt 11111312312321	1
\.


--
-- Data for Name: mk_user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mk_user (id, username, pass) FROM stdin;
5	xuan	$2a$04$kOdxaaWQAMwmGL4W7CaVJOtsJkSqDQXnFsaFj7cMdJb4iLFuosH.2
7	hongxuan	$2a$04$IDRBs7EC.1NnylfRiqthW.U20HcD0hJIhyI.DJ8/7V/LIpmPktJa2
8	hongxuan123	$2a$04$CHII3zusOT72VhqaRVXHSeVZdGQffud0yLnyqXIbx2ougS/JnGiHu
\.


--
-- Name: mk_post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mk_post_id_seq', 2, true);


--
-- Name: mk_session_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mk_session_id_seq', 6, true);


--
-- Name: mk_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mk_user_id_seq', 8, true);


--
-- Name: mk_post mk_post_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mk_post
    ADD CONSTRAINT mk_post_pkey PRIMARY KEY (id);


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
-- Name: mk_session mk_session_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mk_session
    ADD CONSTRAINT mk_session_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.mk_post(id);


--
-- PostgreSQL database dump complete
--

