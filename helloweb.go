package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting...")
	for i := 0; i <= 120; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf(".")
	}
	fmt.Println("Done.")
}
