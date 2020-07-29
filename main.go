package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/pipe.v2"
)

func pipeLine() ([]byte, error) {
	p := pipe.Line(
		pipe.Exec("ioreg", "-l"),
		pipe.Exec("grep", "-v", "Apple"),
		pipe.Exec("grep", "-v", "BatteryData"),
		pipe.Exec("grep", "-e", "MaxCapacity", "-e", "DesignCapacity", "-e", "CurrentCapacity"),
	)
	return pipe.CombinedOutput(p)
}

type Cap struct {
	Data map[string]string
}

func (cap Cap) extraData(data []string) {
	for _, d := range data {
		trimName := strings.TrimLeftFunc(d, func(r rune) bool {
			return string(r) != "\""
		})

		trimName = strings.TrimRightFunc(trimName, func(r rune) bool {
			return string(r) != "\""
		})

		capName := strings.ReplaceAll(trimName, "\"", "")

		trimVol := strings.TrimLeftFunc(d, func(r rune) bool {
			return string(r) != "="
		})

		capVol := strings.ReplaceAll(trimVol, "=", "")
		cap.Data[capName] = capVol
	}
}

func (cap Cap) calRate() error {
	maxCap, err := strconv.ParseFloat(cap.Data["MaxCapacity"], 64)
	if err != nil {
		return err
	}

	designCap, err := strconv.ParseFloat(cap.Data["DesignCapacity"], 64)
	if err != nil {
		return err
	}

	cap.Data["CapacityRate"] = strconv.FormatFloat(maxCap/designCap, 'f', -1, 64)
	return nil
}

func setCap(iv []string, c Cap) {
	iv[0] = c.Data["CurrentCapacity"]
	iv[1] = c.Data["MaxCapacity"]
	iv[2] = c.Data["DesignCapacity"]
	iv[3] = c.Data["CapacityRate"]
}

func setDate(iv []string) {
	now := []int{time.Now().Year(), int(time.Now().Month()), time.Now().Day(), time.Now().Hour()}
	for n := 0; n < len(now); n++ {
		iv[n+4] = strconv.Itoa(now[n])
	}
}

func save(iv []string, path string) error {
	file, err := os.OpenFile(path+"/cap.csv", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	csvFile := csv.NewWriter(file)
	if err := csvFile.Write(iv); err != nil {
		return err
	}

	csvFile.Flush()
	if err := csvFile.Error(); err != nil {
		return err
	}
	return nil
}

func main() {
	byteResults, err := pipeLine()
	if err != nil {
		log.Fatal(err)
	}

	cmdResults := fmt.Sprintf("%s", byteResults)
	cmdSlices := strings.Split(strings.ReplaceAll(cmdResults, " ", ""), "\n")[:3]

	c := Cap{Data: make(map[string]string)}
	c.extraData(cmdSlices)

	if err := c.calRate(); err != nil {
		log.Fatal(err)
	}

	insertVal := make([]string, 8)
	setCap(insertVal, c)
	setDate(insertVal)

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	path := flag.String("path", currentDir, "string flag")
	flag.Parse()

	if err := save(insertVal, *path); err != nil {
		log.Fatal(err)
	}
}
