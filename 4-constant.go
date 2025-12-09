package main

import (
	"fmt"
	"math"
)

const s string = "Constant"

func main() {
	fmt.Println(s)

	const j = 7

	const k = 22 / j
	fmt.Println(k)

	fmt.Println(int64(k))

	fmt.Println(math.Sin(j))
}

// Materi :

// 1. const = nilai tetap
// 	Tidak bisa diubah.
// 	Tidak bisa di-assign ulang.

// 2. const bisa ditulis di mana saja
// 	Di luar function (global)
// 	Di dalam function (local)

// 3. Constant bisa menghitung angka besar dengan presisi tinggi
// 	Go menghitung constant math dengan akurasi maksimal.
// 	Lebih presisi dibanding variable biasa.

// 4. Numeric constant tidak punya tipe
// 	Contoh: const n = 10
// 	Belum jadi int atau float sampai dipakai.

// 5. Constant baru punya tipe kalau dibutuhkan
// 	Cara constant dapat tipe:
// 	Assignment ke variable:
// 	var x int = n

// 	Dipakai di function yang butuh tipe tertentu:
// 	math.Sin(n) // otomatis jadi float64

// 	Explicit conversion:
// 	int64(d)

// 6. String, boolean, dan character juga bisa jadi constant

// 	Contoh:

// 	const s = "hello"
// 	const active = true
// 	const letter = 'a'

// 7. Constant = nilai cepat & aman

// 	Go memeriksa constant saat compile.
// 	Lebih cepat daripada variable biasa.
