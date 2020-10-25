package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/shinnosuke-K/Battery-Watcher/refactoring/command"
)

func main() {

	cmd := command.New()
	cmd.Set()

	if err := cmd.Pipe(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	strOut, err := cmd.GetStrOut()
	if err != nil {
		log.Fatal(err)
	}
	sliceOuts := strings.Split(strings.ReplaceAll(strOut, " ", ""), "\n")[:3]
	fmt.Println(sliceOuts)

	if err := cmd.Done(); err != nil {
		log.Fatal(err)
	}
}
