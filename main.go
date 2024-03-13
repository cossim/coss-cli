package main

import (
	"github.com/cossim/coss-cli/cmd"
	"log"
	"os"
)

func main() {
	err := cmd.App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
