package main

import "fmt"

const (
	errorMessage = "Error"
	infoMessage = "Info"
)

// func {nama-function}({parameter}) ({return})
func hello(nama string) string {
	return "hallo " + nama 
}

func swap(a string, b string) (string, string) {
	return b, a
}

func helloVariadic(name ...string) string {
	msg := "hello, "
	for _, v := range name {
		msg += v + ", "
	}

	return msg
}

func main() {
	// var decimalNumber = 3.14

	// fmt.Println(decimalNumber)
	// fmt.Print("Hello, World!")
	// fmt.Printf("Decimal number menggunakan formating: %f\n", decimalNumber)

	// var isExist bool = true
	// fmt.Println("isExist: ", isExist)

	// var firstName = "Reni"
	// var lastName = "Setyaningsih"

	// var name = firstName + " " + lastName
	// fmt.Println("name: ", name)

	// // varibale decralration
	// // var with data type
	// var decimalNumber1 float64 = 3.1

	// // var without data type
	// var decimalNumber2 = 3.2

	// // short hand
	// decimalNumber3 := 3.3

	// // multi define
	// var decimalNumber4, decimalNumber5 float64

	const firstName string = "Reni"
	// firstName = "Reniii"

	fmt.Println("firstName: ", firstName)

	var buah = [2]string{"apel", "jeruk"}
	
	for i, value := range buah {
		fmt.Println("index: ", i)
		fmt.Println("value: ", value)
		fmt.Println("------------------------------")
	}

	for i := 0; i < 10; i++ {
		fmt.Println("number : ", i)
	}

	for i := 0; i < 10; i++ {
		if i == 2{
			continue
		}

		if i == 8 {
			break
		}

		fmt.Println("number : ", i)
	}

	fmt.Println(hello("Reni"))

	a, b := swap("var A", "var B")
	fmt.Println("var a :", a)
	fmt.Println("var b :", b)

	// a, _ := swap("var A", "var B")
	// fmt.Println("var a :", a)
	// //fmt.Println("var b :", b)

}