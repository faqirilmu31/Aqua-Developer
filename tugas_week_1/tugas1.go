// Nama : Reni Setyaningsih

package main

import (
	"fmt"
)

type item struct {
	nama string
	harga int
	jumlahItem int
}

const (
	jumlahUang = 100000
	itemNumber  = 9
)

func MostItems(items []item) ([]item, int) {
	var bag [itemNumber + 1][jumlahUang + 1]int
	for i := 0; i <= itemNumber; i++ {
		bag[i][0] = 0
	}
	for b := 1; b <= jumlahUang; b++ {
		bag[0][b] = 0
		for i := 1; i <= itemNumber; i++ {
			s := bag[i-1][b]
			if items[i].harga <= b {
				sum := bag[i-1][b-items[i].harga] + items[i].jumlahItem
				if s < sum {
					s = sum
				}
			}
			bag[i][b] = s
		}
	}
	fmt.Println("Barang terbanyak yang dapat dibeli dengan uang Rp.100000 adalah: ")
	i := itemNumber
	b := jumlahUang
	aux := bag[i][b]

	for aux != 0 {
		if bag[i][b] != bag[i-1][b] {
			items[i].jumlahItem = 0
			fmt.Println(items[i].nama, "seharga ", items[i].harga)
			b -= items[i].harga
		}
		i--
		aux = bag[i][b]
	}
	return items, bag[itemNumber][jumlahUang]
}

func MinMaxHarga(items []item) []item {
	var result []item
	min := items[0].harga
	max := items[0].harga
	for _, item := range items {
		if item.harga < min {
			min = item.harga
		}
		if item.harga > max {
			max = item.harga
		}
	}
	fmt.Println("Item yang memiliki harga termurah adalah : ")
	for _, item := range items {
		if item.harga == min {
			result = append(result, item)
			fmt.Println(item.nama, "seharga ", item.harga)
		}
	}
	fmt.Println("Item yang memiliki harga termahal adalah : ")
	for _, item := range items {
		if item.harga == max {
			result = append(result, item)
			fmt.Println(item.nama, "seharga", max)
		}
	}
	return result
}

func itemWithSamePrice(items []item, price int) []item {
	var result []item
	fmt.Println("Item yang memiliki harga Rp.", price, " adalah : ")
	for _, item := range items {
		if item.harga == price {
			result = append(result, item)
			fmt.Println(item.nama, "seharga ", item.harga)
		}
	}
	return result
}

func main() {
	objects := []item{
		{"Benih Lele", 50000, 1},
		{"Pakan lele cap menara", 25000, 1},
		{"Probiotik A", 75000, 1},
		{"Probiotik Nila B", 10000, 1},
		{"Pakan Nila", 20000, 1},
		{"Benih Nila", 20000, 1},
		{"Cupang", 5000, 1},
		{"Benih Nila", 30000, 1},
		{"Benih Cupang", 10000, 1},
		{"Probiotik B", 10000, 1},
	}

	fmt.Println("\n================================================================")
	_, val := MostItems(objects)
	fmt.Println("Total item :", val)
	fmt.Println("\n================================================================")
	MinMaxHarga(objects)
	fmt.Println("\n================================================================")
	itemWithSamePrice(objects, 10000)
}
