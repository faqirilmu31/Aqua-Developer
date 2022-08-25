--------------------

-- CREATE SECTION --

--------------------

create database day_3_rdb;

create table customers(id serial primary key not null, customer_name char(50) not null);

create table products(id serial primary key not null, name char(50) not null);

create table orders(order_id serial primary key not null, customer_id int not null, product_id int not null, order_date date not null, total float not null, foreign key (customer_id) references customers(id), foreign key (product_id) references products(id));


--------------------

-- INSERT SECTION --

--------------------

insert into public.customers (customer_name) values ('Reni'), ('Jeje');

insert into public.products (name) values ('eMall'), ('eFish');

insert into public.orders (customer_id, product_id, order_date, total) values (1, 1, '2022-08-24', 10), (2, 1, '2022-08-24', 15), (2, 2, '2022-08-24', 10);



--------------------

-- UPDATESECTION --

--------------------

update customers set customer_name = 'Jennie' where id = 2;

update products set name = 'eFresh' where id = 2;

update orders set total = 15 where order_id = 3;


--------------------

-- DELETE SECTION --

--------------------

delete from orders where order_id = 2;



--------------------

-- JOIN SECTION --

--------------------

select o.order_id, p.name product_name, c.customer_name, o.order_date, o.total from orders o inner join products p on o.product_id = p.id inner join customers c on o.customer_id = c.id;