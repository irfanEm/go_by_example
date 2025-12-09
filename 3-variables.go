// Didalam Go variable dideklarasikan secara explisit dan digunakan oleh compiler untuk memeriksa keakuratan tipe panggilan fungsi
// 'var' dapat mendeklarasikan satu atau lebih variabel.
// bahkan kita bisa mendeklarasikan beberapa variabel sekaligus.
// Go akan menentukan tipe dari variabel yang diinisialisasi
// Variabel yang dideklarasikan tanpa korespondensi inisialisasi maka akan bernila kosong (zero-valued).
// Sebagai contoh nilai kosong dari sebuah int adalah 0.
// Syntax ':=' adalah singkatan (shorthand) dari deklarasi dan inisialisasi sebuah variable
// sebagai contoh : var f string = 'apple' bisa dituliskan f:='apple'.
// Akan tetapi syntax ini hanya tersedia di dalam function.

package main

import "fmt"

func main() {
	var a = 27
	fmt.Println(a)

	var i, m int = 27, 11
	fmt.Println(i, m)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "f for respect"
	fmt.Println(f)
}
