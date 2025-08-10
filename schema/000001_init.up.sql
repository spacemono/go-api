CREATE TABLE users
(
    id            serial       not null unique,
    email          varchar(255) not null unique,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE user_sessions
(
    id                 serial  NOT NULL PRIMARY KEY,
    user_id            integer NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
)
