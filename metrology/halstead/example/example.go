package example

import "fmt"

func main() {
	var (
		n   = 100
		sum int
	)
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum / n)
}
