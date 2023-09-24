CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists (
    id SERIAL PRIMARY KEY,
    title varchar(255) not null,
    description varchar(255) not null
);

CREATE TABLE todo_items (
    id SERIAL PRIMARY KEY,
    title varchar(255) not null,
    description varchar(255) not null,
    done boolean not null default false
);

CREATE TABLE users_lists (
    id SERIAL PRIMARY KEY,
    user_id int not null,
    list_id int not null,
    FOREIGN KEY (user_id) REFERENCES users ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES users ON DELETE CASCADE
);

CREATE TABLE list_items (
    id SERIAL PRIMARY KEY,
    item_id int not null,
    list_id int not null,
    FOREIGN KEY (item_id) REFERENCES users ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES users ON DELETE CASCADE
);
