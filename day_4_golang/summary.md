# **Ringkasan Materi #04 - Basic Golang**

Go adalah salah satu programming language yang dikelola oleh Google pada tahun 2009 dan open source. eFishery sendiri merupakan perusahaan yang menggunakan golang sebagai programming language mereka. Selain itu, ada Tokopedia, Google, Uber, Gojek, dan masih banyak lagi lainnya.

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

  ```
  var decimalNumber = 2.62

  fmt.Printf("Bilangan desimal : %f\n", decimalNumber)
  fmt.Printf("Bilangan desimal : %.3f\n", decimalNumber)
  ```
  
- **Boolean**

  Tipe data boolean dapat berupa salah satu dari dua nilai, `true` atau `false`, dan didefinisikan sebagai bool saat mendeklarasikannya sebagai tipe data.

  contoh :

  ```
  myBool := 5 > 8
  fmt.Println(myBool)
  ```

- **String**

  String adalah urutan dari satu atau lebih karakter (huruf, angka, simbol) yang dapat berupa konstanta atau variabel. String ada di dalam tanda back quotes ` `` ` atau tanda kutip ganda `" "`
  
  contoh :

  ```
  a := `Say "hello" to Go!
  fmt.Println(a)
  ```
  
## **Variable**

Pada Go terdapate beberapa cara untuk mendeklarasikan variabel. Berikut adalah contohnya :

- **Declaration Variables with data type**

  contoh :
  ```var pi float64```
  
- **Declaration Variables without type data**

  contoh :

  ```
  var firstName string = "John"
  lastName := "Doe"

  fmt.Println("Hallo, % %!\n", firstName, lastName)
  ```

- **Declaration Multiple Variables**

  ```
  var fname, lname string = "John", "Doe"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 2000

	fmt.Println(fname + lname)
	fmt.Println(m + n + o)
	fmt.Println(item, "-", price)
  ```

- **Declaration Underscore Variables**

  Tujuan penggunaan underscore adalah untuk menyimpan variable yang tidak digunakan.

  contoh :

  ```
  _ = "belajar Golang"
  _ = "Golang itu mudah"
  name, _ := "john", "wick"
  ```

- **Constants**

  ```
  const firstName string = "john"
  fmt.Print("Hallo ", firstName, "!\n")
  ```

## **Condition**

- **if..else..then condition**

  ```
  x := 100

  if x == 50 {
      fmt.Println("Germany")
    } else if x == 100 {
      fmt.Println("Japan")
    } else {
      fmt.Println("Canada")
    }
    ```

- **switch case condition**

  ```
  var point = 6

  switch point {
    case 8:
      fmt.Println("perfect")
    case 7:
      fmt.Println("awesome")
    default:
      fmt.Println("not bad")
  }
  ```

  Pada penerapan if else dan switch case ini terdapat operator comparison, yaitu :
  
  |Operator|Name|Description|Example|
  |:-------------|:------------- |:------------- | :------------- |
  | `&&` | Logical And | Mengembalikan nilai `TRUE` apabila kedua statement `TRUE` | `x < y && x > z` |
  | `OR` | Logical Or | Mengembalikan nilai `TRUE` apabila salah satu statement `TRUE` | `x < y OR x > z` |
  | `!` | Logical Not | Membalikkan nilai, mengembalikan nilai `FALSE` apabila nilai `TRUE` | `!(x == y && x > z)` |

## **Looping**

- **for..range**

  ```
  var buah = [4]string{"apel", "pisang", "jeruk", "duku"}
  for i, b := range buah {
    fmt.Printf("Elemen %d : %s\n", i, b)
  }
  ```

- **for..loop**

  ```
  for i := 0; i < 10; i++ {
		fmt.Println("Angka = ", i)
	}
  ```

- **for..loop break continue**

  ```
  for i := 0; i < 10; i++ {
    if i % 2 == 1 {
      continue
    }
    
    if i > 8 {
      break
    }
		fmt.Println("Angka = ", i)
	}
  ```
## **Function**

- **function with parameters**

  ```
  func add(x int, y int) {
    total := 0
    total = x + y
    fmt.Println("Total = ", total)
  }

  func main() {
    add(20, 30)
  }
  ```

- **function with return value**

  ```
  func add(x int, y int) int {
    total := 0
    total = x + y
    return total
  }

  func main() {
    sum := add(20, 30)
    fmt.Println(sum)
  }
  ```


- **functions returning multiple values**

  ```
  func rectangle(l int, b int) (area int, parameter int) {
    parameter = 2 * (l + b)
    area = l * b
    return
  }

  func main() {
    var a, p int
    a, p = rectangle(10, 20)
    fmt.Println("Area = ", a)
    fmt.Println("Parameter = ", p)
  }
  ```

- **anonymous functions**

  ```
  var (
    area = func(l int, b int) int {
      return l * b
    }
  )

  func main() {
    fmt.Println(area(20, 30))
  }
  ```

- **closures functions**

  ```
  func main() {
    l := 20
    b := 30
    
    func() {
      var area int
      area = l * b
      fmt.Println("Area = ", area)
    }()
  }
  ```


- **higher order functions**

  ```
  func sum(x, y int) int {
    return x + y
  }

  func partialSum(x int) func(int) int {
    return func(y int) int {
      return sum(x, y)
    }
  }

  func main() {
    partial := partialSum(3)
    fmt.Println(partial(7))
  }
  ```

## **Struct**

Struct dalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu.

```
type student struct {
  name string
  grade int
}
```

## **Data Structure**

- **map**

  ```
  var cicken map[string]int
  chicken = map[string]int{}

  chicken["januari"] = 50
  chicken["februari"] = 40

  fmt.Println("januari", chicken["januari"])
  fmt.Println("mei", chicken["mei"])
  ```

- **array**

  Array adalah kumpulan data bertipe sama, yang disimpan dalam sebuah variabel.

  ```
  var names [4]string
  names[0] = "trafalgar"
  names[1] = "d"
  names[2] = "water"
  names[3] = "law"

  fmt.Println(names[0], names[1], names[2], names[3])
  ```

- **Slice**

  Slice adalah reference elemen array bedanya tidak perlu mendefinisikan jumlah elemen ketika awal deklarasi. 

  ```
  var fruits = []string{"apple", "grape", "banana", "melon"}
  fmt.Println(fruits[0])
  ```

## **Standard Library (stdllib)**

Golang memiliki banyak standard library, namun `fmt`, `string`, dan `time` adalah library yang sering digunakan.

Untuk penjelasan lebih lengkapnya dapat dilihat di [sini](https://pkg.go.dev/std)

## **Defer**

Defer digunakan untu mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai.

```
package main

import "fmt"

func main() {
    defer fmt.Println("halo")
    fmt.Println("selamat datang")
}
```
