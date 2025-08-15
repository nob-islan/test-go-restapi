CREATE DATABASE snaildb;
USE snaildb;

CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT
    , name VARCHAR(20)
    , age INT
);

INSERT INTO users (
    name
    , age
) VALUES (
    "nob01"
    , 13
), (
    "nob02"
    , 706
);
