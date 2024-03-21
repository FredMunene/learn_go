package main

import "fmt"

type student struct {
	name string
	age  int
}

func changeName(pointer *student, nameChosen string) {
	pointer.name = nameChosen
}

func main() {
	chris := student{}

	fmt.Println(chris)
	chris.name = "Chris"
	chris.age = 30

	fmt.Println(chris)
	changeName(&chris, "Joe")
	fmt.Println(chris)
}
