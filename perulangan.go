package main

import "fmt"


var nama [999]string
var nim [999]string
var ipk [999]float32
var indeks = 0

func cetak(){
	for i := 1; i <= indeks; i++{
		fmt.Println("Data ",i)
		fmt.Println("Nama\t: ",nama[i])
		fmt.Println("NIM\t: ",nim[i])
		fmt.Println("IPK\t: ",ipk[i])
	}
}

func input(){
	var banyak int
	fmt.Println("Input Data")
	
	fmt.Printf("Banyak data : ")
	fmt.Scanf("%d",&banyak) 
	fmt.Scanln()
	for i := 1; i <= banyak; i++{
		fmt.Println("Data ",i)
		fmt.Printf("Nama\t: ")
		fmt.Scanln()
		fmt.Scanf("%s",&nama[indeks])
		fmt.Scanln()
		fmt.Printf("NIM\t: ")
		fmt.Scanf("%s",&nim[indeks])
		fmt.Scanln()
		fmt.Printf("IPK\t: ")
		fmt.Scanf("%f",&ipk[indeks])
		
		indeks++
	}

	cetak()
}

func main()  {
	for{ 
		var pilih int
		fmt.Println("::::: Menu ::::")
		fmt.Println("1. Input Data")
		fmt.Println("2. Tampilkan Data")
		fmt.Println("3. Keluar")
		fmt.Printf("Pilih : ")
		fmt.Scanf("%d",&pilih)
		fmt.Scanln()

		switch pilih {
			case 1: input()
			case 2: cetak()
			case 3: os.Exit(0)	
			}
	}
}
