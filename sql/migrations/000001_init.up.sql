CREATE TABLE users (
  id varchar(36) NOT NULL PRIMARY KEY,
  username varchar(24) NOT NULL,
  password varchar(60) NOT NULL,
  name text NOT NULL,
  type ENUM ('BASIC', 'ADMIN') NOT NULL
)

