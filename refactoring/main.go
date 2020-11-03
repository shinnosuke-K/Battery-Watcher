package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/shinnosuke-K/Battery-Watcher/refactoring/capacity"
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

	if err := cmd.Done(); err != nil {
		log.Fatal(err)
	}

	cap := capacity.New()
	sliceOuts := strings.Split(strings.ReplaceAll(strOut, " ", ""), "\n")[:3]
	cap.SetData(sliceOuts)
	if err := cap.CalcRate(); err != nil {
		log.Fatal(err)
	}

	cap.SetDate()

	fmt.Println(cap)
}
