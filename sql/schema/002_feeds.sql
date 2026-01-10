-- +goose up
CREATE TABLE feeds (
	id uuid PRIMARY KEY,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
	name TEXT NOT NULL,
	url TEXT NOT NULL UNIQUE,
	user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose down
DROP TABLE feeds;