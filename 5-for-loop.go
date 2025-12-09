package main

import "fmt"

func main() {

	fmt.Println("Basic Looping")
	fmt.Println("==============================================================")
	// looping baling basic
	i := 1
	for i <= 3 {
		fmt.Println("Looping ke - ", i)
		i++
	}

	fmt.Println("")
	fmt.Println("Classicc Looping")
	fmt.Println("==============================================================")

	// looping clasic
	for j := 0; j < 3; j++ {
		fmt.Println("Classic loop ke - ", j)
	}

	fmt.Println("")
	fmt.Println("Range Looping")
	fmt.Println("==============================================================")

	// range loop
	for k := range 3 {
		fmt.Println("Range loop ke - ", k)
	}

	fmt.Println("")
	fmt.Println("Infinity Looping")
	fmt.Println("==============================================================")

	// infinity Loop
	for {
		fmt.Println("Infinite looping..")
		break
	}

	fmt.Println("")
	fmt.Println("Continue Looping")
	fmt.Println("==============================================================")

	// continue loop
	for l := range 11 {
		if l%3 == 0 {
			continue
		}

		fmt.Println("Continue looping ke - ", l)
	}
}

// Materi :

// - 'for' adalah satu satunya konstruksi looping (perulangan) Go, berikut adalah beberapa tipe basic dari 'for' :

// 	1. tipe paling basic adalah dengan kondisi tunggal.
// 	contoh :
// 	i:= 5
// 	for i <= 3 {
// 		fmt.Println(i)
// 		i = i+1
// 	}

// 	2. for klasik, yaitu nilai awal - kondisi - setelah.
// 	contoh:
// 	for j := 0; j < 3; j++ {
// 		fmt.Println(j)
// 	}

// 	3. cara lain untuk melakukan perulangan 'lakukan ini selama N kali' adalah dengan menggunakan range(rentang) dari bilangan bulat.
// 	contoh :
// 	for i := range 3 {
// 		fmt.Println("range", i)
// 	}

// 	4. Tanpa kondisi, perulangan for akan melakukan looping terus menerus hingga kita melakukan 'break', atau mengelakukan 'return' dari sebuah function closure.
// 	contoh:

// 	break :
// 	for {
// 		fmt.Println('loop')
// 		break
// 	}

// 	return :
// 	func main() {
// 		for {
// 			fmt.Println('loop')

// 			return
// 		}
// 	}

// 	5. Kita juga bisa melakukan 'continue' ke iterasi selanjutnya dari looping.
// 	contoh :

// 	for n := range 6 {
// 		if n%2 == 0 {
// 			continue
// 		}

// 		fmt.Println(n)
// 	}
