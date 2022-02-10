CREATE USER gopher
WITH PASSWORD '12345';

CREATE DATABASE shortener
    WITH OWNER gopher
    TEMPLATE = 'template0'
    ENCODING = 'utf-8'
    LC_COLLATE = 'C.UTF-8'
    LC_CTYPE = 'C.UTF-8';

\c shortener
SET ROLE gopher;

CREATE SCHEMA IF NOT EXISTS shortener;

DROP TABLE IF EXISTS shortener.shortener CASCADE;
CREATE TABLE shortener.shortener (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    short_url VARCHAR(255) NOT NULL,
    original_URL TEXT NOT NULL,
    visitors_counter INT DEFAULT 0
);

CREATE INDEX ON shortener.shortener(short_url);