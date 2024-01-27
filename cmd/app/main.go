package main

import (
	"fmt"

	"github.com/saboyagustavo/go-hello/internal/app/person"
)

var version = 1.1

func main() {
	var i *float64 = &version

	p := person.Person{}
	fmt.Println("Hello world!")
	fmt.Println("This program is running in its", *i, "version")

	p.GetDataInput()
	p.PrintSalutation()

	version = 1.2
	fmt.Println("Thank you, see in the next version! Hopefully", *i, "version")
}
