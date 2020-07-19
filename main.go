package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

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

	maxCap, err := strconv.ParseFloat(capacity["MaxCapacity"], 64)
	if err != nil {
		log.Fatalln(err)
	}

	designCap, err := strconv.ParseFloat(capacity["DesignCapacity"], 64)
	if err != nil {
		log.Fatalln(err)
	}

	capacity["CapacityRate"] = strconv.FormatFloat(maxCap/designCap, 'f', -1, 64)

	insertVal := make([]string, 8)
	insertVal[0] = capacity["CurrentCapacity"]
	insertVal[1] = capacity["MaxCapacity"]
	insertVal[2] = capacity["DesignCapacity"]
	insertVal[3] = capacity["CapacityRate"]

	now := []int{time.Now().Year(), int(time.Now().Month()), time.Now().Day(), time.Now().Hour()}
	for n := 0; n < len(now); n++ {
		insertVal[n+4] = strconv.Itoa(now[n])
	}

	fmt.Println(insertVal)

	file, err := os.OpenFile("cap.csv", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	csvFile := csv.NewWriter(file)
	if err := csvFile.Write(insertVal); err != nil {
		log.Fatalln(err)
	}

	csvFile.Flush()
	if err := csvFile.Error(); err != nil {
		log.Fatalln(err)
	}
}
