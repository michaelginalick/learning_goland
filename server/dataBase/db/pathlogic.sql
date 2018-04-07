-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.0
-- PostgreSQL version: 9.6
-- Project Site: pgmodeler.com.br
-- Model Author: ---

-- -- object: michaelginalick | type: ROLE --
-- -- DROP ROLE IF EXISTS michaelginalick;
-- CREATE ROLE michaelginalick WITH 
-- 	SUPERUSER
-- 	CREATEDB
-- 	CREATEROLE
-- 	INHERIT
-- 	LOGIN
-- 	REPLICATION
-- 	ENCRYPTED PASSWORD '********';
-- -- ddl-end --
-- 

-- Database creation must be done outside an multicommand file.
-- These commands were put in this file only for convenience.
-- -- object: pathlogic | type: DATABASE --
-- -- DROP DATABASE IF EXISTS pathlogic;
-- CREATE DATABASE pathlogic
-- 	ENCODING = 'UTF8'
-- 	LC_COLLATE = 'en_US.UTF-8'
-- 	LC_CTYPE = 'en_US.UTF-8'
-- 	TABLESPACE = pg_default
-- 	OWNER = michaelginalick
-- ;
-- -- ddl-end --
-- 

-- object: public.sites | type: TABLE --
-- DROP TABLE IF EXISTS public.sites CASCADE;
CREATE TABLE public.sites(
	id integer NOT NULL DEFAULT nextval('sites_id_seq'::regclass),
	name text NOT NULL,
	id_menu bigint,
	CONSTRAINT sites_pkey PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.sites OWNER TO postgres;
-- ddl-end --

-- object: sites_name_index | type: INDEX --
-- DROP INDEX IF EXISTS public.sites_name_index CASCADE;
CREATE INDEX sites_name_index ON public.sites
	USING btree
	(
	  id
	)
	WITH (FILLFACTOR = 90);
-- ddl-end --

-- object: public.menu | type: TABLE --
-- DROP TABLE IF EXISTS public.menu CASCADE;
CREATE TABLE public.menu(
	id bigint NOT NULL,
	menu_id varchar,
	CONSTRAINT menu_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.menu OWNER TO postgres;
-- ddl-end --

-- object: menu_fk | type: CONSTRAINT --
-- ALTER TABLE public.sites DROP CONSTRAINT IF EXISTS menu_fk CASCADE;
ALTER TABLE public.sites ADD CONSTRAINT menu_fk FOREIGN KEY (id_menu)
REFERENCES public.menu (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: sites_uq | type: CONSTRAINT --
-- ALTER TABLE public.sites DROP CONSTRAINT IF EXISTS sites_uq CASCADE;
ALTER TABLE public.sites ADD CONSTRAINT sites_uq UNIQUE (id_menu);
-- ddl-end --

-- object: public.menu_items | type: TABLE --
-- DROP TABLE IF EXISTS public.menu_items CASCADE;
CREATE TABLE public.menu_items(
	id bigint,
	item_id varchar NOT NULL,
	id_wiz_step_items smallint
);
-- ddl-end --
ALTER TABLE public.menu_items OWNER TO postgres;
-- ddl-end --

-- object: public.wizards | type: TABLE --
-- DROP TABLE IF EXISTS public.wizards CASCADE;
CREATE TABLE public.wizards(
	id bigint NOT NULL,
	wiz_id varchar,
	start_step varchar,
	complete_step varchar,
	id_menu bigint,
	id_wiz_step_items smallint,
	CONSTRAINT wizards_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.wizards OWNER TO postgres;
-- ddl-end --

-- object: menu_fk | type: CONSTRAINT --
-- ALTER TABLE public.wizards DROP CONSTRAINT IF EXISTS menu_fk CASCADE;
ALTER TABLE public.wizards ADD CONSTRAINT menu_fk FOREIGN KEY (id_menu)
REFERENCES public.menu (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: public.wiz_step | type: TABLE --
-- DROP TABLE IF EXISTS public.wiz_step CASCADE;
CREATE TABLE public.wiz_step(
	id smallint NOT NULL,
	wiz_step_id varchar,
	wiz_items_id smallint,
	id_wizards bigint,
	id_wiz_step_items smallint,
	wizarditems bigint,
	CONSTRAINT wiz_step_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.wiz_step OWNER TO postgres;
-- ddl-end --

-- object: wizards_fk | type: CONSTRAINT --
-- ALTER TABLE public.wiz_step DROP CONSTRAINT IF EXISTS wizards_fk CASCADE;
ALTER TABLE public.wiz_step ADD CONSTRAINT wizards_fk FOREIGN KEY (id_wizards)
REFERENCES public.wizards (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: public.media | type: TABLE --
-- DROP TABLE IF EXISTS public.media CASCADE;
CREATE TABLE public.media(
	id smallint NOT NULL,
	media_id varchar,
	url varchar,
	CONSTRAINT media_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.media OWNER TO postgres;
-- ddl-end --

-- object: public.inventory_item | type: TABLE --
-- DROP TABLE IF EXISTS public.inventory_item CASCADE;
CREATE TABLE public.inventory_item(
	id smallint NOT NULL,
	name varchar,
	price float,
	type varchar,
	id_wiz_step_items smallint,
	CONSTRAINT inventory_item_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.inventory_item OWNER TO postgres;
-- ddl-end --

-- object: public.wiz_step_items | type: TABLE --
-- DROP TABLE IF EXISTS public.wiz_step_items CASCADE;
CREATE TABLE public.wiz_step_items(
	id smallint NOT NULL,
	CONSTRAINT wiz_step_items_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.wiz_step_items OWNER TO postgres;
-- ddl-end --

-- object: wiz_step_items_fk | type: CONSTRAINT --
-- ALTER TABLE public.wizards DROP CONSTRAINT IF EXISTS wiz_step_items_fk CASCADE;
ALTER TABLE public.wizards ADD CONSTRAINT wiz_step_items_fk FOREIGN KEY (id_wiz_step_items)
REFERENCES public.wiz_step_items (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: wiz_step_items_fk | type: CONSTRAINT --
-- ALTER TABLE public.menu_items DROP CONSTRAINT IF EXISTS wiz_step_items_fk CASCADE;
ALTER TABLE public.menu_items ADD CONSTRAINT wiz_step_items_fk FOREIGN KEY (id_wiz_step_items)
REFERENCES public.wiz_step_items (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: wiz_step_items_fk | type: CONSTRAINT --
-- ALTER TABLE public.inventory_item DROP CONSTRAINT IF EXISTS wiz_step_items_fk CASCADE;
ALTER TABLE public.inventory_item ADD CONSTRAINT wiz_step_items_fk FOREIGN KEY (id_wiz_step_items)
REFERENCES public.wiz_step_items (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: wiz_step_items_fk | type: CONSTRAINT --
-- ALTER TABLE public.wiz_step DROP CONSTRAINT IF EXISTS wiz_step_items_fk CASCADE;
ALTER TABLE public.wiz_step ADD CONSTRAINT wiz_step_items_fk FOREIGN KEY (id_wiz_step_items)
REFERENCES public.wiz_step_items (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: wiz_step_uq | type: CONSTRAINT --
-- ALTER TABLE public.wiz_step DROP CONSTRAINT IF EXISTS wiz_step_uq CASCADE;
ALTER TABLE public.wiz_step ADD CONSTRAINT wiz_step_uq UNIQUE (id_wiz_step_items);
-- ddl-end --


