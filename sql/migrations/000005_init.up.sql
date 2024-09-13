CREATE TABLE carts (
  id varchar(36) NOT NULL PRIMARY KEY
);

CREATE TABLE cart_items (
  id varchar(36) NOT NULL PRIMARY KEY,
  item_id varchar(36) NOT NULL,
  qty smallint NOT NULL,
  cart_id varchar(36) NOT NULL,
  FOREIGN KEY (item_id) REFERENCES items(id),
  FOREIGN KEY (cart_id) REFERENCES carts(id),
  INDEX idx_item_id (item_id),
  INDEX idx_cart_id (cart_id)
);


