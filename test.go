package main

import "fmt"

func main() {
	var firstname string = "Muhammad"
	var middlename string = "Bagus"
	lastname := "Harianto"
	lastname = "dunia"

	// fmt.Println("halo %s %s !\n", "My Name is", firstname, middlename, lastname)
	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstname, middlename, lastname)
	fmt.Println("halo", firstname, middlename, lastname+"!")
}
