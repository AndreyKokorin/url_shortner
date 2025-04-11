-- Создание таблиц
CREATE TABLE public.users (
                              id integer NOT NULL,
                              email character varying(100),
                              password character varying(100),
                              CONSTRAINT users_pkey PRIMARY KEY (id),
                              CONSTRAINT unique_email UNIQUE (email)
);

CREATE TABLE public.links (
                              id integer NOT NULL,
                              original_url text,
                              short_url text,
                              click integer DEFAULT 0,
                              user_id integer,
                              CONSTRAINT links_pkey PRIMARY KEY (id),
                              CONSTRAINT links_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id)
);

-- Создание последовательностей
CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE SEQUENCE public.links_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- Привязка последовательностей к таблицам
ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
ALTER SEQUENCE public.links_id_seq OWNED BY public.links.id;

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
ALTER TABLE ONLY public.links ALTER COLUMN id SET DEFAULT nextval('public.links_id_seq'::regclass);

-- Вставка данных
COPY public.users (id, email, password) FROM stdin;
3	andrey.kokorin.007@gmail.com	$2a$10$FZx6R8K/FpAuY2YckAN99Odk6UugW0.r/JiTknwcpii670c6hmqRa
6	kokorina816@@gmail.com	$2a$10$26Nlh4SFtVD8DQmIr6OjnuvMMr5y2zV48gRoEWOyNUp7WMs3vEOmK
7	newEmail@gmail.com	$2a$10$dk2j.akOBoq2PoHsD68uTuMqOjKESiUaYMuHX/MwQ.3gzk5yCLYuG
\.

COPY public.links (id, original_url, short_url, click, user_id) FROM stdin;
27	https://grok.com/	zKjWXN	2	7
\.

-- Установка значений последовательностей
SELECT pg_catalog.setval('public.users_id_seq', 7, true);
SELECT pg_catalog.setval('public.links_id_seq', 27, true);