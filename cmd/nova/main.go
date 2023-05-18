package main

import (
	"fmt"
	"os"

	"github.com/Nova-foundation/go-nova/cmd/nova/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
