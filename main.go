package main

import (
	"fmt"
	"os"

	"github.com/majd/ipatool/v2/lib"
)

func main() {
	_, err := lib.NewClient(lib.ClientOptions{
		Debug: false,
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing client: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("ipatool library initialized successfully")
}
