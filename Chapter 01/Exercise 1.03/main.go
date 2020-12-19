package main

import (
	"fmt"
	"time"
)

var (
	debug       bool      = false
	logLevel    string    = "info"
	startUpTime time.Time = time.Now()
)

func main() {
	fmt.Println(debug, logLevel, startUpTime)
}
