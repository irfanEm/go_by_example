package main

import (
	"fmt"
	"time"
)

func main() {

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

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Wellcome on weekend !")
	default:
		fmt.Println("It's a weekday !")
	}
}
