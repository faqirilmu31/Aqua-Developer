# **Ringkasan Materi #04 - Basic Golang**

Go adalah salah satu programming language yang dikelola oleh Google pada tahun 2009dan open source. eFishery sendiri merupakan perusahaan yang menggunakan golang sebagai programming language mereka. Selain itu, ada Tokopedia, Google, Uber, Gojek, dan masih banyak lagi lainnya.

## **Standard Library (stdllib)**



## **Data Type**

- **Numeric Non-Decimal**

    |Data Type|Range|
    |:-------------|:------------- |
    |`uint8`| 0 - 255|
    |`uint16`| 0 - 65535|
    |`uint32`| 0 - 4294967295|
    |`uint64`| 0 - 18446744073709551615|
    |`uint`| sama dengan uint32 atau uint64 (tergantung nilai)|
    |`byte`| sama dengan uint8|
    |`int8`| -128 - 127|
    |`int16`| -32768 - 32767|
    |`int32`| -2147483648 - 2147483647|
    |`int64`| -9223372036854775808 - 9223372036854775807|
    |`int`| sama dengan int32 atau int64 (tergantung nilai)|
    |`rune`| sama dengan int32|


- **Numeric Decimal**

  contoh :

  `var decimalNumber = 2.62`

  `fmt.Printf("Bilangan desimal : %f\n", decimalNumber)
  fmt.Printf("Bilangan desimal : %.3f\n", decimalNumber)`
  
- **Boolean**

  Tipe data boolean dapat berupa salah satu dari dua nilai, `true` atau `false`, dan didefinisikan sebagai bool saat mendeklarasikannya sebagai tipe data.

  contoh :

  `myBool := 5 > 8`
  `fmt.Println(myBool)`

- **String**

  String adalah urutan dari satu atau lebih karakter (huruf, angka, simbol) yang dapat berupa konstanta atau variabel. String ada di dalam tanda back quotes ` `` ` atau tanda kutip ganda `" "`
  
  contoh :

  `a := `Say "hello" to Go!`
fmt.Println(a)`
  
## **Variable**

Pada Go terdapate beberapa cara untuk mendeklarasikan variabel. Berikut adalah contohnya :

- **Declaration Variables with data type**

  contoh :
  `var pi float64`
  
- **Declaration Variables without type data**

  contoh :

  `var firstName string = "John"
  lastName := "Doe"`

  `fmt.Println("Hallo, % %!\n", firstName, lastName)`

- **Declaration Multiple Variables**

  `var fname, lname string = "John", "Doe"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 2000`

	`fmt.Println(fname + lname)
	fmt.Println(m + n + o)
	fmt.Println(item, "-", price)`

- **Declaration Underscore Variables**

  Tujuan penggunaan underscore adalah untuk menyimpan variable yang tidak digunakan.

  contoh :

  `_ = "belajar Golang"
  _ = "Golang itu mudah"
  name, _ := "john", "wick"`

- **Constants**

  `const firstName string = "john"
  fmt.Print("Hallo ", firstName, "!\n")`

## **Condition**

- **if..else..then condition**

`x := 100
 
	if x == 50 {
		fmt.Println("Germany")
	} else if x == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}`
- **switch case condition**


## **Looping**

## **Function**

## **Struct**

## **Data Structure**

## **Slice**

## **Defer**
Defer digunakan untu mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai

MEMBUAT DEKLARASI VARIBALE MENGGUNAKAN MAP, STRUCT, SLICE, ARRAY
ASSIGN TYPE DATA MENJADI DATA YANG BARU.