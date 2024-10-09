-- +goose up
CREATE TABLE tokens (
  id SERIAL PRIMARY KEY,
  user_id UUID REFERENCES users (id),
  token TEXT NOT NULL
);

-- +goose down
DROP TABLE tokens;
