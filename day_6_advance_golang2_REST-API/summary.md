# **Ringkasan Materi #06 - Advance Golang (REST API)**

## **API**

Appication Programming Interface atau API adalah sebuah software yang memungkinkan para developer untuk mengintegrasikan dan mengizinkan dua aplikasi yang berbeda secara bersamaan untuk saling terhubung satu sama lain.

![API](https://miro.medium.com/max/1021/1*1efY4vE8NIr92fNImYK1pQ.jpeg)

## **REST API**
Representational State Transfer (REST) merupakan arsitektur untuk web service kaena berjalan di protokol HTTP di level website. REST API sangat populer karena lebih mudah dan datanya lebih readable karena menggunakan JSON, tidak seperti SOAP yang menggunakan XML. Selain itu, REST API juga mempunyai method yang memiliki fungsi khusus.

![](https://images.squarespace-cdn.com/content/v1/60345439404e83020b68ca5f/d9348fbb-dd71-4f25-95d7-faca8af78fdf/main-qimg-2e9beabab9ce3bfcae6b70c889d5fb6e.png)

_Lalu apakah bedanya REST API dengan RESTful API?_

REST API itu arsitekturnya sedangkan RESTFul API itu service yang menggunakan arsitektur REST API.


## **Macam REST API**

  **1. Hypertext Transfer Protocol (HTTP)**

  HTTP adalah protokol tingkat aplikasi untuk sistem informasi terdistribusi, kolaboratif, dan hypermedia.

**2. Hypertext Transfer Protocol Secure (HTTPS)**

  HTTPS merupakan HTTP yanga menggunakan protokol enkripsi untuk mengenkripsi komunikasi.Atau mudahnya, versi securenya HTTP.


  Pada HTTP/1
  Ketika client melakkan reuest makan harus menunggu response dari server. Sehingga, harus menunggu dan tidak bisa terus menerus response dari server.

  Sedangkan HTTP 2 tidak harus menunggu response



HTTP message / body / payload are how data is exxhanged between a server and a client. Ada 

HTTP header let the client and the server pass additional information with an HTTP request or response.

## **Methods**

- **GET**

  digunakan untuk mengambil data


- **POST**

  (tidak ada query params)
  digunakan untuk create dan proses login

- **PUT**

  digunakan untuk megupdate data dan akan mereplace semua existing data, namun apabila data tidak ada maka API akan menentukan untuk membiat resource atau tidak.

  contoh :

- **PATCH**

  Ketika ingin melakukan update pada spesigik data.

- **DELETE**

  digunakan untuk mengapus resource


## 

2xxx : success - mengindikasi client rewuest sukses. misal untu kcreate data

4xx : client error - request yang dikirimkan client tidak valid

5xx : serer error - 
misal di end point maasih ada bug di server. terkadang menyebabkan down atau infinite loop

200 - OK - menginidkasikan request sukse
201 - created - mengindikasikan bahwa request sukses dan sebuah resource baru dibuat
202 - accepted - mengindikasikan mengindikasikan bahwa request telah diterima tetapi tidak complete.
misal create user, tetapi membutuhkan approval untuk masuk ke sistem

400 - bad request - request tidak dapat dikenali karena error sintaks. Perlu untuk dimodifikasi

401 - unauthorized - request membutuhkan authentication information.
403 - forbidden - unauthorized request
misal staff mengakses admin
404 - not found - 
405 - method not allowed - 
misal client request method get tapi dari server 


500 bisa dikirim dari web server maupun gateway
500 - internal server error - server encountered an 

502 - bad gateway

504 - gateway timeout
misal dari sisi client request ke BE dengan maksimmal 1 menit

JWT mempunyai expired yang ditentukan dari sisi server

API gateway efishery = KONG


- apakah URL itu bener2 merespon
- untuk bisa mengerti bagaimana response dari REST API


Di eFishery menggunakan echo framework.
High performance, extensible, dan minimalist


get, get by id, create , update, delete, dan log semua aktivitas di dalam endpoint
export semua collection