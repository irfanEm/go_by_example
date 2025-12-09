package main

import (
	"fmt"
	"time"
)

func main() {

	// Materi :
	// Statement switch mengekspresikan pengkondisian lintas cabang (branch).

	// 1.
	// Berikut adalah contoh basic switch :

	i := 3
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("dua")
	case 3:
		fmt.Println("tiga")
	}

	// 2.
	// Kita bisa menggunakan tanda koma (,) untuk memisahkan beberapa expression pada statement case yang sama.
	// Slain itu pada contoh berikut kita juga menggunakan case optional 'default'.

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Wellcome on weekend !")
	default:
		fmt.Println("It's a weekday !")
	}

	// 3.
	// Jika kita menggunakan switch tanpa expression, itu sama seperti kita mengekspresikan(melakukan) logic if-else.
	// Pada contoh berikut juga ditunjukan bahwa expression case dapat bernilai non-constant (artinya bisa bernilai true / false).

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Ini masih pagi.")
	default:
		fmt.Println("Ini sudah siang.")
	}

	// 4.
	// Type switch adalah fitur khusus di Go yang bisa kita gunakan untuk memeriksa tipe dari sebuah nilai, bukan nilai-nya itu sendiri.
	// Ini sangat berguna ketika kita bekerja menggunakan 'interface{}' (empty interface).
	// Pada contoh berikut variabel 't' akan memiliki tipe yang sesuai dengan kondisinya.

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Aku adalah boolean")

		case int:
			fmt.Println("Aku adalah integer")

		default:
			fmt.Printf("Tipe tidak diketahui %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(3)
	whatAmI("hay")

	// Contoh lain dari type switch
	checkType := func(q interface{}) {
		switch m := q.(type) {
		case string:
			fmt.Printf("%s tipe data string (panjang: %d)\n", m, len(m))

		case int:
			fmt.Printf("Integer: %d (double: %d)\n", m, m*2)

		case float64:
			fmt.Printf("Float: %.2f\n", m)

		case bool:
			fmt.Printf("Boolean: %m\n", m)

		case []int:
			fmt.Printf("Slice of int with %d elements\n", len(m))

		default:
			fmt.Printf("Unknown type: %T\n", m)
		}
	}

	checkType("Hello")        // String: Hello (length: 5)
	checkType(42)             // Integer: 42 (double: 84)
	checkType(3.14)           // Float: 3.14
	checkType(true)           // Boolean: true
	checkType([]int{1, 2, 3}) // Slice of int with 3 elements
	checkType(make(chan int)) // Unknown type: chan int

}
