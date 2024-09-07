
CREATE TYPE  "user_position" AS ENUM (
    'ADMIN',
    'COURIER',
    'USER'
    );


CREATE TABLE  "users"
(
    "id"           uuid UNIQUE PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "first_name"   text,
    "last_name"    text,
    "email"         text,
    "password"      text,
    "position"      user_position not null default 'USER',
    "created_at"   timestamp                        DEFAULT (current_timestamp),
    "updated_at"   timestamp,
    "deleted_at"   timestamp
);