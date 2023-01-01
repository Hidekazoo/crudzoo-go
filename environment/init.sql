CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE tasks (
    task_id uuid DEFAULT uuid_generate_v4 () NOT NULL,
    subject text,
    link text,
    body text,
    created_at timestamp with time zone
);

INSERT INTO tasks VALUES ('F5D06514-D694-4422-BC9E-50FF24009DDF', 'subject', 'https://example.com', 'body', now());
INSERT INTO tasks(subject, link, body, created_at) VALUES ('subject', 'https://example.com', 'body', now());