package main

import (
	"backend/Backend"
	"fmt"
	"os"
)

func main() {
	errBackend := Backend.StartBackend()
	if errBackend != nil {
		fmt.Println(errBackend)
		os.Exit(1)
	}

	os.Exit(0)

}
