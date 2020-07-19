package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"gopkg.in/pipe.v2"
)

func main() {
	p := pipe.Line(
		pipe.Exec("ioreg", "-l"),
		pipe.Exec("grep", "-v", "Apple"),
		pipe.Exec("grep", "-v", "BatteryData"),
		pipe.Exec("grep", "-e", "MaxCapacity", "-e", "DesignCapacity", "-e", "CurrentCapacity"),
	)

	byteResults, err := pipe.CombinedOutput(p)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	cmdResults := fmt.Sprintf("%s", byteResults)
	cmdSlices := strings.Split(strings.ReplaceAll(cmdResults, " ", ""), "\n")[:3]

	capacity := make(map[string]int)
	for _, cs := range cmdSlices {
		trimName := strings.TrimLeftFunc(cs, func(r rune) bool {
			return string(r) != "\""
		})

		trimName = strings.TrimRightFunc(trimName, func(r rune) bool {
			return string(r) != "\""
		})

		capName := strings.ReplaceAll(trimName, "\"", "")

		trimVol := strings.TrimLeftFunc(cs, func(r rune) bool {
			return string(r) != "="
		})

		capVol := strings.ReplaceAll(trimVol, "=", "")

		capacity[capName], err = strconv.Atoi(capVol)
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println(capacity)
}
