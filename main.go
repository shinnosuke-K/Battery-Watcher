package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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

	capacity := make(map[string]string)
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
		capacity[capName] = capVol
	}

	fmt.Println(capacity)

	file, err := os.OpenFile("cap.csv", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	// Turn: MaxCapacity, CurrentCapacity, DesignCapacity
	csvFile := csv.NewWriter(file)
	var insertVal []string
	for k, v := range capacity {
		fmt.Println(k, v)
		insertVal = append(insertVal, v)
	}

	if err := csvFile.Write(insertVal); err != nil {
		log.Fatalln(err)
	}

	csvFile.Flush()
	if err := csvFile.Error(); err != nil {
		log.Fatalln(err)
	}
}
