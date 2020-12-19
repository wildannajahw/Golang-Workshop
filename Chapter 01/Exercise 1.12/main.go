package main

import(
	"fmt"
	"time"
)

func main()  {
	var(
		count int
		discount float64
		debug bool
		message string
		emails []string
		startTime time.Time
	)
	fmt.Printf("Count  : %#v \n", count)
	fmt.Printf("Count  : %#v \n", discount)
	fmt.Printf("Count  : %#v \n", debug)
	fmt.Printf("Count  : %#v \n", message)
	fmt.Printf("Count  : %#v \n", emails)
	fmt.Printf("Count  : %#v \n", startTime)
}