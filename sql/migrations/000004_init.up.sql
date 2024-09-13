CREATE TABLE items (
  id varchar(36) NOT NULL PRIMARY KEY,
  name varchar (50) NOT NULL,
  price decimal(10,2) NOT NULL,
  image text NOT NULL,
  expiration_time varchar(10),
  weight varchar(10),
  suitable_for ENUM ("DOGS", "CATS"),
  material varchar (20),
  type ENUM ("FOOD", "TOY") NOT NULL
)