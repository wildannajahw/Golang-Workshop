package main

import (
	"fmt"
	"time"
)

var (
	debug       = false
	logLevel    = "info"
	startUpTime = time.Now()
)

func main() {
	fmt.Println(debug, logLevel, startUpTime)
}
