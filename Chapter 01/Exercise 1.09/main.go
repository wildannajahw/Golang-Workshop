package main

import "fmt"

func main() {

	var total float64 = 2 * 13
	fmt.Println("Sub :", total)

	total = total + (4 * 2.25)
	fmt.Println("Sub :", total)

	total -= 5
	fmt.Println("Sub :", total)

	tip := total * 0.1
	fmt.Println("Tip :", tip)

	total += tip
	fmt.Println("Total :", total)

	split := total / 2
	fmt.Println("Split :", split)

	visitCount := 24
	visitCount++

	remainder := visitCount % 5
	if remainder == 0 {
		fmt.Println("With this visit, you've earned a reward")
	}
}
