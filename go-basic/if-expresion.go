package main

import "fmt"

func main() {

	var nilai int = 85

	//fmt.Println("Masukan nilai ujian : ")
	//fmt.Scan(&nilai)

	if nilai >= 85 {
		fmt.Println("Nilai anda A")
	} else if nilai >= 75 {
		fmt.Println("Nilai anda B")
	} else if nilai >= 65 {
		fmt.Println("Nilai Anda C")
	} else {
		fmt.Println("Niai anda bukan E")
	}

	name := "ahmad kurniawan"

	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
