CREATE TABLE users
(
    id            serial  not null unique,
    name          varchar not null,
    username      varchar not null unique,
    password_hash varchar not null
);

CREATE TABLE todo_lists
(
    id          serial  not null unique,
    title       varchar not null,
    description varchar
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_items
(
    id          serial  not null unique,
    title       varchar not null,
    description varchar,
    done        boolean not null default false
);

CREATE TABLE items_lists
(
    id      serial                                           not null unique,
    user_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);