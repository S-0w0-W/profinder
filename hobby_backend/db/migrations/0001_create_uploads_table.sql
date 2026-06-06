-- +goose Up
CREATE TABLE upload_file (
  id SERIAL PRIMARY KEY,
  uploader VARCHAR(50) NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT NOW()
);
CREATE TABLE jobs (
    id         SERIAL PRIMARY KEY,
    title      TEXT NOT NULL,
    status     TEXT NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE upload_file;
DROP TABLE IF EXISTS jobs;
