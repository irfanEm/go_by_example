// Materi :

// Didalam Go array adalah sebuah urutan element bernomor dengan panjang tertentu.
// Pada code umum Go, 'slices' lebih sering digunakan, sedangkan 'array' akan sangat berguna di beberapa skenario spesial.

package main

import "fmt"

func main() {

	// Disini kita akan membuat array yang akan menyimpan 5 integer secara tepat (pas, tidak lebih / kurang)
	// Dimana tipe dan panjang element, keduanya adalah bagian dari tipe array.
	// Secara default array akan zero-valued(bernilai kosong) yang artinya dalam int adalah 0.
	// contoh :
	var a [5]int
	fmt.Println("emp : ", a) // output = emp :  [0 0 0 0 0]

	// Kita dapat mengatur nilai dari sebuah index array dengan menggunakan format = array[index] = nilai.
	// dan untuk mendapatkan nilai dari index array menggunakan = array[index].
	// contoh :
	a[4] = 100
	fmt.Println("set", a)    // output = set [0 0 0 0 100]
	fmt.Println("get", a[4]) // output = get 100

	// method bawaan len() akan mengembalikan panjang dari sebuah array.
	// contoh :
	fmt.Println("len :", len(a)) // output = len : 5

	// Kita juga bisa mendeklarasikan dan langsung menginisialisasi dalam satu baris
	// contoh :
	b := [5]int{27, 11, 19, 97, 0}
	fmt.Println("dcl :", b) // output = dcl : [27 11 19 97 0]

	// Kita juga bisa memiliki compiler count (compiler penghitung) untuk menghitung jumlah element array menggunakan '...' (titik tiga)
	// contoh :
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl :", b)

	// Kita juga bisa menentukan index dengan tanda ":" (titik dua).
	// contoh:
	b = [...]int{200, 3: 500, 580}
	// pada contoh diatas, index diantara index ke 3 dan ke 0 akan dikosongkan (0)
	fmt.Println("idx:", b) // output = [200 0 0 500 580]

	// Tipe array itu satu dimensi, tapi kita bisa menyusun type untuk membangun struktur data multi dimensi.
	// contoh :
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d :", twoD)

}
