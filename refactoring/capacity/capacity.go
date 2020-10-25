package capacity

import (
	"log"
	"strconv"
	"strings"
)

type Capacity struct {
	Data map[string]string
}

func New() *Capacity {
	return &Capacity{
		Data: map[string]string{},
	}
}

func (cap *Capacity) SetData(outputs []string) {
	for _, output := range outputs {
		replacedOutput := strings.Replace(output, "||\"", "", 1)
		replacedOutput = strings.Replace(replacedOutput, "\"", "", 1)

		slicedOutput := strings.Split(replacedOutput, "=")
		cap.Data[slicedOutput[0]] = slicedOutput[1]
	}
}

func (cap *Capacity) CalcRate() error {
	maxCap, err := strconv.ParseFloat(cap.Data["MaxCapacity"], 64)
	if err != nil {
		return err
	}

	designCap, err := strconv.ParseFloat(cap.Data["DesignCapacity"], 64)
	if err != nil {
		log.Fatal(err)
	}

	cap.Data["CapacityRate"] = strconv.FormatFloat(maxCap/designCap, 'f', -1, 64)
	return nil
}
