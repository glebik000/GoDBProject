-- Table: public.group_services

-- DROP TABLE IF EXISTS public.group_services;

CREATE TABLE IF NOT EXISTS public.group_services
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    code character varying(200) COLLATE pg_catalog."default" NOT NULL,
    name character varying(500) COLLATE pg_catalog."default" NOT NULL,
    hidden boolean NOT NULL,
    CONSTRAINT group_services_pkey PRIMARY KEY (id)
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.group_services
    OWNER to postgres;

-- Table: public.measure_unit

-- DROP TABLE IF EXISTS public.measure_unit;

CREATE TABLE IF NOT EXISTS public.measure_unit
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    short_name character varying(20) COLLATE pg_catalog."default",
    CONSTRAINT measure_unit_pkey PRIMARY KEY (id)
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.measure_unit
    OWNER to postgres;



ALTER TABLE IF EXISTS public.prod_to_service
    OWNER to postgres;

-- Table: public.products

-- DROP TABLE IF EXISTS public.products;

CREATE TABLE IF NOT EXISTS public.products
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    code character varying(200) COLLATE pg_catalog."default" NOT NULL,
    name character varying(500) COLLATE pg_catalog."default" NOT NULL,
    basecost numeric(17,2),
    hidden boolean,
    measure_id integer,
    CONSTRAINT products_pkey PRIMARY KEY (id)
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.products
    OWNER to postgres;

-- Table: public.services

-- DROP TABLE IF EXISTS public.services;

CREATE TABLE IF NOT EXISTS public.services
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name character varying(200) COLLATE pg_catalog."default" NOT NULL,
    code character varying(200) COLLATE pg_catalog."default" NOT NULL,
    basecost numeric(17,2),
    hidden boolean,
    group_id integer,
    CONSTRAINT services_pkey PRIMARY KEY (id)
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.services
    OWNER to postgres;

-- Table: public.prod_to_service

-- DROP TABLE IF EXISTS public.prod_to_service;

CREATE TABLE IF NOT EXISTS public.prod_to_service
(
    prod_id integer NOT NULL,
    service_id integer NOT NULL,
    count_of_prod integer NOT NULL,
    CONSTRAINT prod_to_service_pkey PRIMARY KEY (prod_id, service_id),
    CONSTRAINT product_key FOREIGN KEY (prod_id)
    REFERENCES public.products (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID,
    CONSTRAINT service_key FOREIGN KEY (service_id)
    REFERENCES public.services (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID
    )

    TABLESPACE pg_default;