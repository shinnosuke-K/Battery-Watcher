package main

import (
	"fmt"
	"log"
	"os"
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
	s.SetValues(cap.Data, cap.Name, now)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fileName := "/refactoring/cap.csv"

	if err := s.CreateFile(pwd, fileName); err != nil {
		log.Fatal(err)
	}

	if err := s.Do(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}

func main() {
	c := cron.New(cron.WithSeconds())
	if _, err := c.AddFunc("@every 1s", do); err != nil {
		log.Fatal(err)
	}
	c.Run()
}
