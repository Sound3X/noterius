BEGIN;

CREATE TABLE players (
	id serial NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	username text NOT NULL UNIQUE,
	email text NOT NULL,
	password text NOT NULL,
	pincode text,
	is_active boolean NOT NULL
);


CREATE TABLE maps (
	id serial NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	name text[] NOT NULL UNIQUE
);


CREATE TABLE characters (
	id serial NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	player_id bigint REFERENCES players(id),
	name text NOT NULL UNIQUE,
	job text NOT NULL,
	map_id bigint REFERENCES maps(id),
	level integer NOT NULL,
	race integer NOT NULL,
	enabled boolean NOT NULL
);


COMMIT;
