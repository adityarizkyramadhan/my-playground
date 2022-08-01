package main

import "fmt"

type Siswa struct {
	Nama   string
	NIS    string
	Alamat string
}

var UpdateSiswaRefrence = func(siswa *Siswa, nama string) *Siswa {
	siswa.Nama = nama
	return siswa
}

var UpdateSiswaValue = func(siswa Siswa, nama string) Siswa {
	siswa.Nama = nama
	return siswa
}

var NewSiswa = func(Nama, NIS, Alamat string) *Siswa {
	return &Siswa{
		Nama:   Nama,
		NIS:    NIS,
		Alamat: Alamat,
	}
}

func main() {
	var siswa = NewSiswa("Rizki", "12345", "Jl. Raya")
	fmt.Println(siswa)
	//pass by reference
	UpdateSiswaRefrence(siswa, "Budi")
	fmt.Println("Pass by reference")
	fmt.Println(siswa)
	siswa2 := UpdateSiswaValue(*siswa, "Anton")
	fmt.Println("Pass by value")
	fmt.Println(siswa)
	fmt.Println(siswa2)
}
