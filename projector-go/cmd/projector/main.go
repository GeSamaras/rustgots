package main

import (
	"fmt"
	"log"

	"github.com/theprimeagen/projector/pkg/projector"
)

func main() {
	opts, err := projector.GetOpts()
	if err != nil {
		log.Fatalf("Error getting opts: %v", err)
	}

	fmt.Printf("opts: %v\n", opts)
}