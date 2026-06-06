-- +goose Up
INSERT INTO jobs (title, status) VALUES
    ('Backend Engineer', 'open'),
    ('Frontend Engineer', 'open'),
    ('DevOps Engineer',   'closed');

-- +goose Down
DELETE FROM jobs WHERE title IN ('Backend Engineer', 'Frontend Engineer', 'DevOps Engineer');