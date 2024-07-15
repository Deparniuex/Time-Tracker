CREATE TABLE tasks(
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    task_starts timestamp(0) without time zone,
    task_name text NOT NULL,
    task_description text NOT NULL,
    task_ends timestamp(0) without time zone
);