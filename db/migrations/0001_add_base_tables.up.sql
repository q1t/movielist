CREATE TABLE movies (
	id SERIAL PRIMARY KEY,
	imdb_id text NOT NULL,
    watched boolean,
	user_rating decimal,
	list_id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
CREATE TABLE lists (
	id SERIAL PRIMARY KEY,
	title varchar(50) NOT NULL,
	user_id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	email text,
	username text NOT NULL UNIQUE CHECK(username <> ''),
	password text NOT NULL UNIQUE CHECK(username <> ''),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
)
