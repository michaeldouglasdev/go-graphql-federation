CREATE TABLE pets (
  id varchar(36) NOT NULL PRIMARY KEY,
  name varchar(24) NOT NULL,
  birthday varchar(10) NOT NULL,
  breed ENUM ('DOG', 'CAT') NOT NULL
);