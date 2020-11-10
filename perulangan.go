package main

import (
	"fmt"
	"os" //jangan lupa ditambah os. soalnya kamu gunakan untuk di exit
)

var nama [999]string
var nim [999]string
var ipk [999]float32
var indeks = 0

func cetak() {
	for i := 0; i < indeks; i++ {
		// index kamu mulai dari 0
		// harusnya perulangan juga dimulai dari nol
		// perhatikan bagian input
		fmt.Println("Data ", i+1)
		fmt.Println("Nama\t: ", nama[i])
		fmt.Println("NIM\t: ", nim[i])
		fmt.Println("IPK\t: ", ipk[i])
	}
}

func input() {
	var banyak int
	fmt.Println("Input Data")

	fmt.Printf("Banyak data : ")
	fmt.Scan(&banyak)
	fmt.Scanln()
	for i := 1; i <= banyak; i++ {
		fmt.Println("Data ", i)
		fmt.Printf("Nama\t: ")
		//	fmt.Scanln() ini harusnya dihilangkan.
		// kalau engga, setiap input, huruf pertamanya aka nhilang
		fmt.Scanf("%s", &nama[indeks])
		fmt.Scanln()
		fmt.Printf("NIM\t: ")
		fmt.Scanf("%s", &nim[indeks])
		fmt.Scanln()
		fmt.Printf("IPK\t: ")
		fmt.Scanf("%f", &ipk[indeks])
		fmt.Scanln() // tambahan baru biar ga numpuk yang selanjutnya
		indeks++
	}

	cetak()
}

func main() {
	for {
		var pilih int
		fmt.Println("::::: Menu ::::")
		fmt.Println("1. Input Data")
		fmt.Println("2. Tampilkan Data")
		fmt.Println("3. Keluar")
		fmt.Printf("Pilih : ")
		fmt.Scan(&pilih) // diubah ke scan aja
		// aku gatau kenapa di scanf datanya ga bisa nyimpen baru lagi
		// mungkin kena buffer. entahlah

		switch pilih {
		case 1:
			input()
		case 2:
			cetak()
		case 3:
			os.Exit(0)
		}
	}
}
