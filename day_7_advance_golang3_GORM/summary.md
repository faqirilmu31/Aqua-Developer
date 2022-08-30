# **Ringkasan Materi #07 - Advance Golang3 (REST API)**

Pada hari ke-7 ini materi yang dibahas adalah tentang GORM, clean architecture, middleware, dan environment variable.

## **Go Object-Relationnal Mapping (GORM)**

ORM merupakan teknik yang memungkinkan kita melakaukan wuery dan manipulasi data

## **Clean Architecture**

onion layer

Manfaat :

- bisa ditest
- maintainable ada layering ada 
- gampang diuabh atau changeable
- easy to develop
- easy to deploy, biasanyaakan ketahuan bagianamana yg controller
- independent karena dipisahkan per modul dan abstraksinya menggunakan interface

repository yang emngurusi ssemua tentang database

module tempat dimana bisnis logic

delivery yang menghandel dan memvlidasi semua equest response

service 

3 komponen besar


delivery (semua hal yg rser perlukan)

membahas semua teknologi yan digunakan. misal REST API graphql, SOA

service -> logic

respotitory -> data source, mysql, no sql, cache, eksternal api

supaya ketika ada perubahan pada suatu layer tidak memperngaruhi layer yang lain

## **Middleware**

Seoah oleh membungkus aplikasi

swaggo digunakan utntuk mendokumentasikan 

dari sisi fe proses pengambilan data lebih cepat

## **Environment Variables**

variable yang disimpan dalam OS untuk keperluan sistem kita