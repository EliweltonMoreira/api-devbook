DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id serial PRIMARY KEY,
  name varchar(50) NOT NULL,
  nick varchar(50) NOT NULL UNIQUE,
  email varchar(50) NOT NULL UNIQUE,
  password varchar(100) NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE followers (
  user_id int NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  follower_id int NOT NULL,
  FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
  PRIMARY KEY (user_id, follower_id)
);