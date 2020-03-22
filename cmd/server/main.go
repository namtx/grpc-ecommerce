package main

import (
	"fmt"
	"os"

	cmd "github.com/namtx/grpc-ecommerce/pkg/cmd/server"
)

var err error

func main() {
	/*
		if err = cmd.RunProductServiceServer(); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)

			os.Exit(1)
		}
	*/

	if err := cmd.RunOrderServiceServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}
}
