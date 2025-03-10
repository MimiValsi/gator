-- +goose Up
CREATE TABLE feed_follows (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id) UNIQUE,
  feed_id UUID NOT NULL REFERENCES feeds(id) UNIQUE
);

-- +goose Down
DROP TABLE feed_follows;
