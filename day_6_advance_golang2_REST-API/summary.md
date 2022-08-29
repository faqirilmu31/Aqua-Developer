# **Ringkasan Materi #06 - Advance Golang (REST API)**

## **API**

Appication Programming Interface atau API adalah sebuah software yang memungkinkan para developer untuk mengintegrasikan dan mengizinkan dua aplikasi yang berbeda secara bersamaan untuk saling terhubung satu sama lain.

![API](https://miro.medium.com/max/1021/1*1efY4vE8NIr92fNImYK1pQ.jpeg)

## **REST API**
Representational State Transfer (REST) merupakan arsitektur untuk web service kaena berjalan di protokol HTTP di level website. REST API sangat populer karena lebih mudah dan datanya lebih readable karena menggunakan JSON, tidak seperti SOAP yang menggunakan XML. Selain itu, REST API juga mempunyai method yang memiliki fungsi khusus.

Di eFishery menggunakan KONG sebagai API gatewaynya.

![](https://images.squarespace-cdn.com/content/v1/60345439404e83020b68ca5f/d9348fbb-dd71-4f25-95d7-faca8af78fdf/main-qimg-2e9beabab9ce3bfcae6b70c889d5fb6e.png)

_Lalu apakah bedanya REST API dengan RESTful API?_

REST API itu arsitekturnya sedangkan RESTFul API itu service yang menggunakan arsitektur REST API.


## **Macam REST API**

  1. **Hypertext Transfer Protocol (HTTP)**

      HTTP adalah protokol tingkat aplikasi untuk sistem informasi terdistribusi, kolaboratif, dan hypermedia.

2. **Hypertext Transfer Protocol Secure (HTTPS)**

    HTTPS merupakan HTTP yanga menggunakan protokol enkripsi untuk mengenkripsi komunikasi.Atau mudahnya, versi securenya HTTP. Pada HTTP/1 ketika client melakukan request ke server maka harus menunggu response dari server. Sehingga tidak bisa terus menerus response dari server.

    Lain halnya dengan HTTP/2 yang tidak harus menunggu response dari server. Misal client mengirimkan 3 request maka ketiga request tersebut dapat diproses.


## **Bagian-bagian HTTP**

1. **HTTP message / body / payload**

    HTTP body merupakan bagaimana data dipertukarkan antara server dan client. Ada dua jenis pesan:
    - request, yang dikirim oleh klien untuk memicu tindakan di server
    - response, jawaban dari server.

2. **HTTP header**

    HTTP header memungkinkan client dan server menyampaikan informasi tambahan dengan permintaan atau respons HTTP. Header HTTP terdiri dari case-insensitive name yang diikuti oleh titik dua (:), kemudian dengan nilainya. Spasi sebelum nilai diabaikan.

## **REST API Methods**

1. `GET`

    GET berfungsi untuk membaca data/resource dari REST server.

2. `POST`

    POST berfungsi untuk membuat sebuah data/resource baru di REST server. Contohnya create dan proses login.

3. `PUT`

    PUT digunakan untuk megupdate data dan akan mereplace semua existing data, namun apabila data tidak ada maka API akan menentukan untuk membiat resource atau tidak.

4. `PATCH`

    PATCH digunakan untuk melakukan update pada spesifik data.

5. `DELETE`

    DELETE berfungsi untuk menghapus data/resource dari REST server


## **Status Code**

1. **2xx : success** - mengindikasi client request sukses. misal untuk create data

    - **200 - OK** - mengindikasikan request sukses
    - **201 - created** - mengindikasikan bahwa request sukses dan sebuah resource baru dibuat
    - **202 - accepted** - mengindikasikan mengindikasikan bahwa request telah diterima tetapi tidak complete. Misal create user, tetapi membutuhkan approval untuk masuk ke sistem

2. **4xx : client error** - request yang dikirimkan client tidak valid

    - **400 - bad request** - request tidak dapat dikenali karena error sintaks. Sintaks erlu untuk dimodifikasi.
    - **401 - unauthorized** - request membutuhkan authentication information. Client harus mengulangi request dengan authorisasi yang sesuai.
    - **403 - forbidden** - unauthorized request. misal staff mengakses admin
    - **404 - not found** - server tidak dapat menemukan resource yang diminta
    - **405 - method not allowed** - request HTTP dikenali tetapi disable dan tidak dapat digunakan untuk resource tersebut. Misal client request method GET tapi dari server harus menerima POST

3. **5xx : serer error** - 
misal di end point masih ada bug di server dan terkadang menyebabkan down atau infinite loop. Error code 5xx bisa dikirim dari web server maupun gateway.
    - **500 - internal server error** - server mengalami kondisi tak terduga yang mencegahnya memenuhi request. 
    - **502 - bad gateway** - server mendapat respons yang tidak valid saat bekerja sebagai gateway untuk mendapatkan respons yang diperlukan untuk menangani request.
    - **504 - gateway timeout** - server bertindak sebagai gateway dan tidak bisa mendapatkan respons tepat waktu untuk permintaan. Misal dari sisi client request ke BE lebih dari 1 menit.


## **Authorization**

Autorisasi kali ini berfokus pada JWT. **JSON Web Token (JWT)** - adalah standar terbuka (RFC 7519) yang mendefinisikan cara yang ringkas dan mandiri untuk mentransmisikan informasi antar pihak secara aman sebagai objek JSON. Informasi ini dapat diverifikasi dan dipercaya karena ditandatangani secara digital. JWT dapat ditandatangani menggunakan rahasia (dengan algoritma HMAC) atau pasangan kunci publik/pribadi menggunakan RSA atau ECDSA. JWT sendiri mempunyai expired yang ditentukan dari sisi server.

Format JWT adalah sebagai berikut :

`<header>.<payload>.<signature>`

- Header - header berisi algoritma dan type token

- Payload - payload berisi klaim. Klaim adalah pernyataan tentang suatu entitas (biasanya, pengguna) dan data tambahan.

- Signature - untuk membuat bagian signature kita harus mengambil header yang disandikan, payload yang disandikan, secret, algoritma yang ditentukan di header, dan menandatanganinya.

Contoh :

`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`


Untuk lebih lengkapnya dapat dilihat di [website JWT](https://jwt.io/) langsung.


## **Postman**

Postman adalah sebuah aplikasi yang berfungsi sebagai REST CLIENT untuk uji coba REST API. Dari sisi front-end, postman ini biasa digunakan untuk :

- mengetes apakah URL itu benar-benar merespon
- mengerti bagaimana response dari REST API

## **ECHO Framewrok**

Echo framework merupakan framework yang dominan digunakan oleh eFishery. Alasan penggunaan framework ini adalah high performance, extensible, dan minimalis.

Untuk dapat menginstallnya kita cukup jalankan command berikut :

`go get github.com/labstack/echo/v4` pada terminal. Kemudian Golang akan mendownload framework tersebut dan siap digunakan.
