-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.8.1-alpha1
-- PostgreSQL version: 9.4
-- Project Site: pgmodeler.com.br
-- Model Author: ---


-- Database creation must be done outside an multicommand file.
-- These commands were put in this file only for convenience.
-- -- object: geenpeil | type: DATABASE --
-- -- DROP DATABASE IF EXISTS geenpeil;
-- CREATE DATABASE geenpeil
-- ;
-- -- ddl-end --
-- 

-- object: public.handtekeningen | type: TABLE --
-- DROP TABLE IF EXISTS public.handtekeningen CASCADE;
CREATE TABLE public.handtekeningen(
	id serial NOT NULL,
	insert_time timestamp NOT NULL,
	iphash bytea NOT NULL,
	CONSTRAINT handtekeningen_id_primary PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.handtekeningen OWNER TO postgres;
-- ddl-end --

-- object: public.nawhashes | type: TABLE --
-- DROP TABLE IF EXISTS public.nawhashes CASCADE;
CREATE TABLE public.nawhashes(
	hash bytea NOT NULL,
	CONSTRAINT nawhash_unique UNIQUE (hash)

);
-- ddl-end --
ALTER TABLE public.nawhashes OWNER TO postgres;
-- ddl-end --


