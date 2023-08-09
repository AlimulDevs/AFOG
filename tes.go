package main

import "fmt"

type Mahasiswa struct {
	Nama string
	Nim  string
	Umur int
}

func Tes() {
	mahasiswa := Mahasiswa{"Alice", "12345", 21}
	fmt.Printf("Mahasiswa: %s, Nim: %s, Umur: %d\n", mahasiswa.Nama, mahasiswa.Nim, mahasiswa.Umur)

	angka := []int{1, 2, 3, 4, 5}
	for _, nilai := range angka {
		fmt.Println("Angka:", nilai)
	}
}
