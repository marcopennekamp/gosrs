-- +goose Up

CREATE TABLE member (
  -- General information.
  id bigserial PRIMARY KEY,
  name text NOT NULL,  -- UNIQUE
  email text NOT NULL, -- UNIQUE
  password bytea NOT NULL,
  registration_date timestamptz NOT NULL,

  -- SRS-specific.
  level int NOT NULL
);

CREATE UNIQUE INDEX member_name_index ON member(name);
CREATE UNIQUE INDEX member_email_index ON member(email);


CREATE TABLE kanji (
  id serial PRIMARY KEY,
  character character(1) NOT NULL -- UNIQUE
);

CREATE UNIQUE INDEX kanji_character_index ON kanji(character);


CREATE TABLE kanji_status (
  member_id bigint NOT NULL REFERENCES member(id),
  kanji_id integer NOT NULL REFERENCES kanji(id),
  is_known boolean NOT NULL
);


CREATE TABLE word (
  id serial PRIMARY KEY,
  expression text NOT NULL, -- UNIQUE
  required_level integer NOT NULL
);

CREATE UNIQUE INDEX word_expression_index ON word(expression);


CREATE TABLE word_kanji_dependencies (
  word_id integer NOT NULL REFERENCES word(id),
  kanji_id integer NOT NULL REFERENCES kanji(id)
);


CREATE TABLE meaning (
  id serial PRIMARY KEY,
  expression text NOT NULL -- UNIQUE
);

CREATE UNIQUE INDEX meaning_expression_index ON meaning(expression);


CREATE TABLE reading (
  id serial PRIMARY KEY,
  expression text NOT NULL -- UNIQUE
);

CREATE UNIQUE INDEX reading_expression_index ON reading(expression);


CREATE TABLE word_meaning (
  word_id integer NOT NULL REFERENCES word(id),
  meaning_id integer NOT NULL REFERENCES meaning(id)
);


CREATE TABLE word_reading (
  word_id integer NOT NULL REFERENCES word(id),
  reading_id integer NOT NULL REFERENCES reading(id)
);


CREATE TABLE word_status (
  member_id bigint NOT NULL REFERENCES member(id),
  word_id integer NOT NULL REFERENCES word(id),
  level integer NOT NULL
);


CREATE TABLE review (
  member_id bigint NOT NULL REFERENCES member(id),
  word_id integer NOT NULL REFERENCES word(id),
  due timestamptz NOT NULL
);

CREATE INDEX review_due_index ON review(due);


-- +goose Down

DROP TABLE review;
DROP TABLE word_status;
DROP TABLE word_reading;
DROP TABLE word_meaning;
DROP TABLE reading;
DROP TABLE meaning;
DROP TABLE word_kanji_dependencies;
DROP TABLE word;
DROP TABLE kanji_status;
DROP TABLE kanji;
DROP TABLE member;


