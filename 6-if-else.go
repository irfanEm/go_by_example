package main

import "fmt"

func main() {

	fmt.Println("=======================================================")
	fmt.Println("1. Contoh basic if-else.")
	fmt.Println("=======================================================")
	if 7%2 == 0 {
		fmt.Println("7 adalah bilangan genap")
	} else {
		fmt.Println("7 adalah bilangan ganjil")
	}

	fmt.Println("")
	fmt.Println("=======================================================")
	fmt.Println("2. 'If' dapat digunakan tanpa menggunakan 'else'.")
	fmt.Println("=======================================================")
	if 8%2 == 0 {
		fmt.Println("8 adalah bilangan genap.")
	}

	fmt.Println("")
	fmt.Println("=======================================================")
	fmt.Println("3. Kita juga bisa menggunakan operatori logic seperti '&&' dan '||'.")
	fmt.Println("=======================================================")
	if 8%2 == 0 && 8 > 0 {
		fmt.Println("8 adalah bilangan genap dan bilangan positif.")
	}
}
