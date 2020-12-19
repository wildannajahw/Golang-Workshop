package main

import (
	"fmt"
	"time"
)

func main() {
	debug := false
	logLevel := "info"
	startUpTime := time.Now()
	fmt.Println(debug, logLevel, startUpTime)
}
