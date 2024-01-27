package person

import (
	"fmt"

	"github.com/saboyagustavo/go-hello/utils"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) PrintSalutation() {
	fmt.Println("Hello", p.Name, "your age is", p.Age)
}

func (p *Person) GetDataInput() {
	fmt.Println("Enter your name:")
	fmt.Scan(&p.Name)

	fmt.Println("Enter your age:")
	for {
		_, err := fmt.Scan(&p.Age)
		if err != nil {
			fmt.Println("Enter a valid number:")
			utils.ClearInputBuffer()
		} else {
			fmt.Println("Got: ", p.Age)
			break
		}
		utils.ClearInputBuffer()
	}
}
