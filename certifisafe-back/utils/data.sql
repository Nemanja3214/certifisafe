-- create db if it doesn't already exist
SELECT 'CREATE DATABASE certisafe'
    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'certisafe');
DROP TABLE IF EXISTS certificates;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(50),
    password VARCHAR(30),
    first_name VARCHAR(30),
    last_name VARCHAR(30),
    is_admin BOOLEAN
);

CREATE TABLE certificates (
     id SERIAL PRIMARY KEY,
     name VARCHAR(30),
     valid_from date,
     valid_to date,
     subject_pk varchar(30),
     subject_id INTEGER REFERENCES users(id),
     issuer_id INTEGER REFERENCES users(id),
     signature  varchar(30)
);


INSERT INTO users(email, password, first_name, last_name, is_admin) VALUES('project.usertest+sladic@outlook.com', '', 'Goran', 'Sladic', true);
INSERT INTO users(email, password, first_name, last_name, is_admin) VALUES('project.usertest+majstorovic@outlook.com', '', 'Nemanja', 'Majstorovic', false);
INSERT INTO users(email, password, first_name, last_name, is_admin) VALUES('project.usertest+dutina@outlook.com', '', 'Nemanja', 'Dutina', false);
INSERT INTO users(email, password, first_name, last_name, is_admin) VALUES('project.usertest+milosavljevic@outlook.com', '', 'Branko', 'Milosavljevic', false);

INSERT INTO certificates(name, valid_from, valid_to, subject_id, subject_pk, issuer_id, signature)
VALUES('Certificate #1', '2022-01-01', '2024-01-01', 2, 'asd', 1, 'asdasd');