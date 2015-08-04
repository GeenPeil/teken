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
	"ID" serial NOT NULL,
	insert_time timestamp NOT NULL,
	iphash char(32) NOT NULL,
	CONSTRAINT handtekeningen_id_primary PRIMARY KEY ("ID")

);
-- ddl-end --
ALTER TABLE public.handtekeningen OWNER TO postgres;
-- ddl-end --

-- object: public.nawhash | type: TABLE --
-- DROP TABLE IF EXISTS public.nawhash CASCADE;
CREATE TABLE public.nawhash(
	hash char(32) NOT NULL,
	CONSTRAINT nawhash_unique UNIQUE (hash)

);
-- ddl-end --
ALTER TABLE public.nawhash OWNER TO postgres;
-- ddl-end --


