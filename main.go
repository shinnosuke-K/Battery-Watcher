package main

import (
	"fmt"
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
	cmdSlices := strings.Split(strings.ReplaceAll(cmdResults, " ", ""), "\n")

	//capacity := make(map[string]int)
	for _, cs := range cmdSlices {
		trimName := strings.TrimLeftFunc(cs, func(r rune) bool {
			if string(r) == "\"" {
				return false
			}
			return true
		})

		capName := strings.TrimRightFunc(trimName, func(r rune) bool {
			if string(r) == "\"" {
				return false
			}
			return true
		})

		fmt.Println(strings.ReplaceAll(capName, "\"", ""))

		trimVol := strings.TrimLeftFunc(cs, func(r rune) bool {
			if string(r) == "=" {
				return false
			}
			return true
		})

		capVol := strings.ReplaceAll(trimVol, "=", "")
		fmt.Println(capVol)
	}
}
