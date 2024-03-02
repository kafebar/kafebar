CREATE TABLE orders(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    status TEXT
);

CREATE TABLE order_items(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    order_id INTEGER,
    product_id INTEGER,
    status TEXT,

    FOREIGN KEY(order_id) REFERENCES orders(id)
    FOREIGN KEY(product_id) REFERENCES products(id)
);

CREATE TABLE order_item_options(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    order_id INTEGER,
    order_item_id INTEGER,
    option TEXT,

    FOREIGN KEY(order_id) REFERENCES orders(id),
    FOREIGN KEY(order_item_id) REFERENCES order_items(id)
);

CREATE TABLE products(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name INTEGER,
    price REAL,
    picture TEXT
);

CREATE TABLE product_options(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    product_id INTEGER,
    option TEXT,

    FOREIGN KEY(product_id) REFERENCES products(id)
);