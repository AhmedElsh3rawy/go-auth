CREATE TABLE users (
  id UUID PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password TEXT NOT NULL
);

CREATE TABLE tokens (
  id SERIAL PRIMARY KEY,
  user_id UUID REFERENCES users (id),
  token TEXT NOT NULL
);
