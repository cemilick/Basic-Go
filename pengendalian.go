package main

import "fmt"

func main() {
	var angka = 9387

	fmt.Println("angka = ")
	fmt.Println(angka)

	if angka%2==0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}
}