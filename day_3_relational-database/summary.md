# **Ringkasan Materi #03 - Basic Relational Database**

Database adalah sekumpulan data yang saling terhubung sehingga memudahkan dalam pengelolaan. Macam-macam database secara umum ada 5, yaitu operational database, database warehouse, distributed database, relational database, dan end-user database.

[Relational database](https://cloud.google.com/learn/what-is-a-relational-database#:~:text=A%20relational%20database%20(RDB)%20is,relationship%20between%20various%20data%20points.) atau RDB adalah cara untuk menyusun informasi dalam tabel, baris, dan kolom. RDB memiliki kemampuan untuk membuat link—atau hubungan—antara informasi dengan menggabungkan tabel, yang memudahkan untuk memahami dan memperoleh wawasan tentang hubungan antara berbagai titik data. Pada pembahasan kali ini akan difokuskan pada penggunaan postgreSQL.

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
