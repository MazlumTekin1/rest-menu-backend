CREATE TABLE rest_user.categories (
	category_id bigserial NOT NULL,
	menu_id int8 NOT NULL,
	category_name varchar(100) NOT NULL,
	category_image_url text NULL,
	created_date timestamptz NULL,
	updated_date timestamptz NULL,
	is_deleted bool DEFAULT false NULL,
	create_user_id int8 NOT NULL,
	update_user_id int8 NOT NULL,
	CONSTRAINT categories_pkey PRIMARY KEY (category_id)
);


-- rest_user.categories foreign keys

ALTER TABLE rest_user.categories ADD CONSTRAINT fk_rest_user_restaurant_menus_categories FOREIGN KEY (menu_id) REFERENCES rest_user.restaurant_menus(menu_id);


CREATE TABLE rest_user.products (
	product_id bigserial NOT NULL,
	menu_id int8 NOT NULL,
	category_id int8 NULL,
	product_name varchar(100) NOT NULL,
	product_price numeric(10, 2) NOT NULL,
	product_description text NULL,
	product_image_url text NULL,
	created_date timestamptz NULL,
	updated_date timestamptz NULL,
	is_deleted bool DEFAULT false NULL,
	create_user_id int8 NOT NULL,
	update_user_id int8 NOT NULL,
	CONSTRAINT products_pkey PRIMARY KEY (product_id)
);


-- rest_user.products foreign keys

ALTER TABLE rest_user.products ADD CONSTRAINT fk_rest_user_categories_products FOREIGN KEY (category_id) REFERENCES rest_user.categories(category_id);



CREATE TABLE rest_user.restaurant_menus (
	menu_id bigserial NOT NULL,
	restaurant_id int8 NOT NULL,
	menu_name varchar(100) NOT NULL,
	created_date timestamptz NULL,
	updated_date timestamptz NULL,
	is_deleted bool DEFAULT false NULL,
	create_user_id int8 NOT NULL,
	update_user_id int8 NOT NULL,
	CONSTRAINT restaurant_menus_pkey PRIMARY KEY (menu_id)
);


CREATE TABLE rest_user.users (
	user_id serial4 NOT NULL,
	username varchar(50) NOT NULL,
	"password" varchar(100) NOT NULL,
	email varchar(100) NOT NULL,
	roles _text DEFAULT '{}'::text[] NOT NULL,
	created_date timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_date timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	is_deleted bool DEFAULT false NOT NULL,
	create_user_id int4 NOT NULL,
	update_user_id int4 NOT NULL,
	CONSTRAINT users_email_unique UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (user_id),
	CONSTRAINT users_username_unique UNIQUE (username)
);
CREATE INDEX idx_users_email ON rest_user.users USING btree (email) WHERE (NOT is_deleted);
CREATE INDEX idx_users_username ON rest_user.users USING btree (username) WHERE (NOT is_deleted);