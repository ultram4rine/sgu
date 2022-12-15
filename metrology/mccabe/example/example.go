package example

import "fmt"

func main() {
	if true {
		fmt.Println("a")

		if true {
			fmt.Println("aa")
		} else if false {
			fmt.Println("ab")
		} else {
			fmt.Println("ac")
		}
	} else {
		fmt.Println("b")
	}

	if true {
		fmt.Println("c")
	} else {
		fmt.Println("d")
	}
}
