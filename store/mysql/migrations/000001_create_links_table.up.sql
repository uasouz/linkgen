CREATE TABLE links
(
    id          int auto_increment primary key,
    shortid     varchar(10)   not null,
    originalURL varchar(1000) not null
);