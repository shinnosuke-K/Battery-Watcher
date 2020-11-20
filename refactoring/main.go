package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/shinnosuke-K/Battery-Watcher/refactoring/capacity"
	"github.com/shinnosuke-K/Battery-Watcher/refactoring/command"
	"github.com/shinnosuke-K/Battery-Watcher/refactoring/save"
)

func do() {
	now := time.Now()

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

	s := save.New()
	s.SetValues(cap.Data, now)

	fmt.Println(s)

}

func main() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("@every 1s", do)
	c.Run()
}
