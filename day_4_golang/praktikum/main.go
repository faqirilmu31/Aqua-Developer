//Praktikum Hari ke-4 - Basic Golang
//Nama : Reni Setyaningsih

/* TUGAS : Membuat deklarasi variable menggunakan struct, map, array, dan slice. Kemudian assign data menjadi data yang baru*/

package main

import "fmt"

type mahasiswa struct {
	nama string
	usia int
	univ universitas
}

type universitas struct {
	fakultas string
	prodi string
}

func main() {
	// Initialize Struct
	fmt.Println("================= Struct =================")

	fmt.Println("\n----------- Initialize Struct -----------")
	m0520067 := mahasiswa{
		nama: "Reni Setyaningsih",
		usia: 20,
		univ: universitas{
			fakultas: "Fakultas Matemaatika dan Ilmu Pengetahuan Alam",
			prodi: "S1 Informatika",
		},
	}

	fmt.Printf(
		"Nama \t\t: %s\nUsia \t\t: %d\nFakultas \t: %s\nProdi \t\t: %s\n",
		m0520067.nama,
		m0520067.usia,
		m0520067.univ.fakultas,
		m0520067.univ.prodi,
	)

	fmt.Println("\n----------- Edit Struct -----------")
	m0520067.univ.fakultas = "Fakultas Teknologi Informasi dan Sains Data"

	fmt.Printf(
		"Nama \t\t: %s\nUsia \t\t: %d\nFakultas \t: %s\nProdi \t\t: %s\n",
		m0520067.nama,
		m0520067.usia,
		m0520067.univ.fakultas,
		m0520067.univ.prodi,
	)

	fmt.Println("\n================= Map =================")
	fmt.Println("\n----------- Initialize Map -----------")
	var mapMahasiswa map[string]mahasiswa
	mapMahasiswa = make(map[string]mahasiswa)

	mapMahasiswa["M0520067"] = mahasiswa{
		nama: "Reni Setyaningsih",
		usia: 20,
		univ: universitas{
			fakultas: "Fakultas Matemaatika dan Ilmu Pengetahuan Alam",
			prodi: "S1 Informatika",
		},
	}

	for key, value := range mapMahasiswa {
		fmt.Println("Key \t\t: ", key)
		fmt.Printf(
			"Nama \t\t: %s\nUsia \t\t: %d\nFakultas \t: %s\nProdi \t\t: %s\n",
			value.nama,
			value.usia,
			value.univ.fakultas,
			value.univ.prodi,
		)
	}

	fmt.Println("\n----------- Edit Map -----------")
	mapMahasiswa["M0520067"] = mahasiswa{
		nama: "Reni Setyaningsih",
		usia: 20,
		univ: universitas{
			fakultas: "Fakultas Teknologi Informasi dan Sains Data",
			prodi: "S1 Informatika",
		},
	}

	for key, value := range mapMahasiswa {
		fmt.Println("Key \t\t: ", key)
		fmt.Printf(
			"Nama \t\t: %s\nUsia \t\t: %d\nFakultas \t: %s\nProdi \t\t: %s\n",
			value.nama,
			value.usia,
			value.univ.fakultas,
			value.univ.prodi,
		)
	}

	fmt.Println("\n================= Array =================")
	fmt.Println("\n----------- Initialize Array -----------")
	var arrayMahasiswa = [...]mahasiswa {
		{
			nama: "Budi Prakoso",
			usia: 21,
			univ: universitas{
				fakultas: "Fakultas Kedokteran",
				prodi: "S1 Pendidikan Kedokteran",
			},
		},
		{
			nama: "Bunga Citra Lestari",
			usia: 20,
			univ: universitas{
				fakultas: "Fakultas Pertanian",
				prodi: "S1 Agribisnis",
			},
		},
	}

	for _, array := range arrayMahasiswa {
		// fmt.Println("Indeks ke : ", i)
		fmt.Printf(
			"Nama \t\t: %s\nUsia \t\t: %d\nFakultas \t: %s\nProdi \t\t: %s\n\n",
			array.nama,
			array.usia,
			array.univ.fakultas,
			array.univ.prodi,
		)
	}

	fmt.Println("\n----------- Edit Array -----------")
	arrayMahasiswa[1] = mahasiswa{
		nama: "Bunga Citra Lestari",
		usia: 20,
		univ: universitas{
			fakultas: "Fakultas Hukum",
			prodi: "S1 Hukum",
		},
	}

	for _, array := range arrayMahasiswa {
		// fmt.Println("Indeks ke : ", i)
		fmt.Printf(
			"Nama \t\t: %s\nUsia \t\t: %d\nFakultas \t: %s\nProdi \t\t: %s\n\n",
			array.nama,
			array.usia,
			array.univ.fakultas,
			array.univ.prodi,
		)
	}

	fmt.Println("\n================= Slice =================")
	fmt.Println("\n----------- Initialize Slice -----------")
	var sliceMahasiswa = []mahasiswa{
		{
			nama: "Gita Gutawa",
			usia: 22,
			univ: universitas{
				fakultas: "Fakultas Teknik",
				prodi: "S1 Teknik Industri",
			},
		},
		{
			nama: "Fedi Nuril",
			usia: 22,
			univ: universitas{
				fakultas: "Fakultas Kedokteran",
				prodi: "S1 Pendidikan Dokter",
			},
		},
	}

	for _, slice := range sliceMahasiswa {
		fmt.Printf(
			"Nama \t\t: %s\nUsia \t\t: %d\nFakultas \t: %s\nProdi \t\t: %s\n\n",
			slice.nama,
			slice.usia,
			slice.univ.fakultas,
			slice.univ.prodi,
		)
	}

	fmt.Println("\n----------- Edit Slice -----------")
	sliceMahasiswa[1] = mahasiswa{
		nama: "Fedi Nuril",
		usia: 22,
		univ: universitas{
			fakultas: "Fakultas Kedokteran",
			prodi: "S1 Psikologi",
		},
	}

	for _, slice := range sliceMahasiswa {
		fmt.Printf(
			"Nama \t\t: %s\nUsia \t\t: %d\nFakultas \t: %s\nProdi \t\t: %s\n\n",
			slice.nama,
			slice.usia,
			slice.univ.fakultas,
			slice.univ.prodi,
		)
	}
}