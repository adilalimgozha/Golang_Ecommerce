INSERT INTO roles (role_id, role_name) VALUES (1, 'user'),(2, 'admin');


INSERT INTO users (username, password_hash, email, role_id) 
VALUES ('admin1', 'hashed_password', 'admin1@mail.ru', 2);

INSERT INTO products
VALUES (2, 'T-Shirt', 'description T-Shirt', 1500, 30, 2);

INSERT INTO categories
VALUES (2, 'short', 'short clothes');

INSERT INTO product_images
VALUES (1, 1, 'url');


select * from roles
select * from users
select * from products
select * from categories
select * from product_images