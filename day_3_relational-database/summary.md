# **Ringkasan Materi #03 - Basic Relational Database**

Database adalah sekumpulan data yang saling terhubung sehingga memudahkan dalam pengelolaan. Macam-macam database secara umum ada 5, yaitu operational database, database warehouse, distributed database, relational database, dan end-user database.

[Relational database](https://cloud.google.com/learn/what-is-a-relational-database#:~:text=A%20relational%20database%20(RDB)%20is,relationship%20between%20various%20data%20points.) atau RDB adalah cara untuk menyusun informasi dalam tabel, baris, dan kolom. RDB memiliki kemampuan untuk membuat link—atau hubungan—antara informasi dengan menggabungkan tabel, yang memudahkan untuk memahami dan memperoleh wawasan tentang hubungan antara berbagai titik data. Pada pembahasan kali ini akan difokuskan pada penggunaan [postgreSQL](https://www.postgresql.org/).

Berikut adalah beberapa basic pemahaman tentang relational database pada postgreQSL :

## **Data Definition Language (DDL)**
DDL merupakan perintah SQL yang digunakan untuk mendefinisikan struktur seperti skema, database, tabel, constraint, dsb. Contoh sintaksnya adalah :

|query|contoh|fungsi|
|:------------- | :--------- |:---------- |
|create| `create database <database_name>;` | membuat database dan tabel|
|alter| `alter table <table_name> rename to <table_name_change>;` | menambah, menghapus, atau mengubah kolom dalam tabel|
|drop| `drop table <table_name>;` | menghapus database atau tabel|

## **Data Manipulation Language (DML)**
DML adalah perintah SQL yang berhubungan dengan manipulasi atau pengolahan data dalam table.
Contoh sintaksnya adalah :

|query|contoh|fungsi|
|:------------- | :--------- |:---------- |
|insert| `insert into public.<database_name> (name, age) values ('reni', 20);` | menambahkan data ke dalam tabel|
|select| `select * from product;` | menampilkan data dari tabel|
|update| `update public.<table_name> set age = 21 where name = 'reni';` | mengupdate data pada tabel|
|delete| `delete from public.<table_name> where name = 'reni';` | menghapus data pada tabel|

## **Definition Control Language (DCL)**
DCL biasa digunakan untuk merubah hak akses, memberikan roles, dan isu lain yang berhubungan dengan keamanan database. Contoh dari DCL adalah `grant` dan `revoke`.

## **Join**
Join digunakan untuk menggabungkan data atau baris dari satu (self-join) atau lebih tabel berdasarkan bidang yang sama di antara mereka. Bidang umum ini umumnya adalah primary key dari tabel pertama dan foreign key dari tabel lainnya.

Tipe-tipe join :

![](https://4.bp.blogspot.com/-_HsHikmChBI/VmQGJjLKgyI/AAAAAAAAEPw/JaLnV0bsbEo/s1600/sql%2Bjoins%2Bguide%2Band%2Bsyntax.jpg)

## **Agregation**
Agregation function digunakan untuk menghasilkan hasil yang diringkas dan beroperasi pada set baris. Agregation mengembalikan hasil berdasarkan kelompok baris. Secara default, semua baris dalam tabel diperlakukan sebagai satu grup.

Secara umum fungsi agregasi yang sering digunakan adalah `count`, `sum`, `min`, `max`, `avg`.

Untuk lebih rincinya, dapat dilihat pada halaman [postreSQL aggregation function](https://www.postgresql.org/docs/9.5/functions-aggregate.html).

## **Subquery**
Subquery adalah query SQL yang bersarang di dalam query yang lebih besar atau bisa disebut query dalam query.

Contoh :
`update products set stock = subquery.stock - 2 from(select id, stock from products where id = 1) as subquery where products.id = 1`

Pada query di atas kita akan mengambil data dari query sarangnya terlebih dahulu yaitu product dengan id = 1. Kemudian barulah akan dikurangkan 2.

## **Function**

Function juga dikenal dengan stored procedures, mmemungkinkan untuk melakukan operasi yang biasanya mengambil beberapa query dan melakukan beberapa hal dalam satu fungsi dalam database.

contoh :
`CREATE FUNCTION kurangi_stock(INT, INT) RETURNS products AS 'update products set stock = subquery.stock - $2 from(select id, stock from products where id = $1) as subquery where products.id = $1;
select * from products where id = $1'
LANGUAGE = 'sql';`

Function di atas sama halnya dengan aggregation di atas, yaitu untuk mengurangi stock - 2.

