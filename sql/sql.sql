CREATE DATABASE IF NOT EXISTS devbook;
-- \c devbook (only works in psql CLI)

DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id serial PRIMARY KEY,
  name varchar(50) NOT NULL,
  nick varchar(50) NOT NULL UNIQUE,
  email varchar(50) NOT NULL UNIQUE,
  password varchar(20) NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP
);