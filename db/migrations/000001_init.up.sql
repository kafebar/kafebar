CREATE TABLE products(
    id SERIAL PRIMARY KEY,
    name TEXT,
    price REAL,
    picture TEXT
);

CREATE TABLE product_options(
    id SERIAL PRIMARY KEY,
    product_id INTEGER,
    option TEXT,

	CONSTRAINT "product_options_product_id" FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    name TEXT,
    is_archived BOOLEAN
);

CREATE TABLE order_items(
    id SERIAL PRIMARY KEY,
    order_id INTEGER,
    product_id INTEGER,
    status TEXT,

	CONSTRAINT "order_items_order_id" FOREIGN KEY (order_id) REFERENCES orders(id),
	CONSTRAINT "order_items_product_id" FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE order_item_options(
    id SERIAL PRIMARY KEY,
    order_id INTEGER,
    order_item_id INTEGER,
    option TEXT,


	CONSTRAINT "order_item_options_order_id" FOREIGN KEY (order_id) REFERENCES orders(id),
	CONSTRAINT "order_item_options_order_item_id" FOREIGN KEY (order_item_id) REFERENCES order_items(id)
);
