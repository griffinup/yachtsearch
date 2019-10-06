DROP TABLE IF EXISTS yachts;
CREATE TABLE yachts (
  id integer PRIMARY KEY,
  name VARCHAR(32) NOT NULL,
  company integer NOT NULL,
  model integer NOT NULL
);

DROP TABLE IF EXISTS companies;
CREATE TABLE companies (
  id integer PRIMARY KEY,
  name VARCHAR(64) NOT NULL
);

DROP TABLE IF EXISTS models;
CREATE TABLE models (
  id integer PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  builder integer NOT NULL
);

DROP TABLE IF EXISTS builders;
CREATE TABLE builders (
  id integer PRIMARY KEY,
  name VARCHAR(64) NOT NULL
);