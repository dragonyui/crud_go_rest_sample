create database products;
use products;
CREATE TABLE product
(
  id INT PRIMARY KEY AUTO_INCREMENT,
  sku VARCHAR(255) UNIQUE ,
  name VARCHAR(255),
  price DECIMAL,
  qty INT
);