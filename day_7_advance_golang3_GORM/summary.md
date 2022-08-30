# **Ringkasan Materi #07 - Advance Golang pt.3 **

Pada hari ke-7 ini materi yang dibahas adalah tentang GORM, clean architecture, middleware, dan environment variable.

## **Golang Object-Relationnal Mapping (GORM)**

ORM merupakan teknik yang memungkinkan kita untuk melakukan query dan manipulasi data dari database menggunakan paradigma object-oriented.

Untuk melakukan installasi, kita hanya perlu menjalankan command berikut pada terminal :

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

Kemudian, untuk mengoneksikan dengan database, khususnya postgreSQL yang berada di lokal, kita dapat menambahkan :

```
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

dsn := "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable TimeZome=Asia/Jakarta"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```


## **Middleware**

Middleware ini seperti pipeline atau seoah-olah membungkus aplikasi sebelum menuju ke end pointnya. Penggunaan middleware ini dari sisi front end developer memudahkan melakukan pengambilan data dan lebih cepat.

![](https://miro.medium.com/max/945/1*RgPEcCE3mHSGR-fS5lXTCQ.png)

Ada banyak third-party yang dapat digunakan di go untuk mendokumentasikannya, yaitu swaggo digunakan untuk mendokumentasikan, yaitu echo-contrib, oapi-codegen, swaggo, ziflex, echozap, darkweak, dan lain-lain. 

Di eFishery menggunakan swaggo untuk mendokumentasikannya. Swaggo dikenal mudah karena kita hanya perlu membuat comment di dalam REST API maka akan otomatis tergenerate.


## **Environment Variables**

Environment variable adalah variabel yang berada di lapisan runtime sistem operasi. Karena environment variable adalah variabel seperti biasa, kita dapat melakukan operasi seperti mengubah nilainya atau mengambil nilainya.

Untuk menambahkan environment variable, kita dapat melakukannya dengan :

```
confAppName := os.Getenv("APP_NAME")
if confAppName == "" {
	e.Logger.Fatal("APP_NAME config is required")
}
```

## **Clean Architecture**

Clean architecture adalah sebuah filosofi design softare yang membuat sebuah sistem mudah untuk dipahami, dikembangkan, dan di maintain, ke dalam sebuah ring level atau onion layer. Clean architecture ini digunakan supaya ketika ada perubahan pada suatu layer tidak mempengaruhi layer yang lain.

Manfaat penggunaan clean architecture :

- mudah untuk di-testing
- mudah untuk di-maintain karena berbentuk layer 
- mudah diubah atau changeable
- easy to develop
- easy to deploy karena biasanyaakan ketahuan bagian mana yg controller, service, dan sebagainya.
- independent karena dipisahkan per modul dan abstraksinya menggunakan interface

![](https://cdn-media-1.freecodecamp.org/images/oVVbTLR5gXHgP8Ehlz1qzRm5LLjX9kv2Zri6)

Terdapat 3 komponen utama pada clean architecture, yaitu :

1. **delivery** (semua hal yg user perlukan akan ada di sini)

	Delivery layer berfungsi untuk :
	- memvalidasi inputan
	- pasing body JSON
	- return http response

2. **service** (berisi logic)

	Service layer berfungsi :
	- sebagai business process
	- logic
	- olah data

3. **respository** (berisi data source, mysql, no sql, cache, eksternal api)

	Repository layer berfungsi :
	- untuk komunikasi dengan database
	- query ke database
	- memanggil API jika menggunakan micro service