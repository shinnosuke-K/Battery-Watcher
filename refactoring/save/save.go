package save

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type InsertValue struct {
	Values  []string
	CSVFile *os.File
}

func New() *InsertValue {
	return &InsertValue{}
}

func (iv *InsertValue) SetValues(data map[string]string, now time.Time) {
	for _, v := range data {
		iv.Values = append(iv.Values, v)
	}
	iv.Values = append(iv.Values, strconv.Itoa(now.Year()))
	iv.Values = append(iv.Values, strconv.Itoa(int(now.Month())))
	iv.Values = append(iv.Values, strconv.Itoa(now.Day()))
	iv.Values = append(iv.Values, strconv.Itoa(now.Hour()))
}

func (iv *InsertValue) CreateFile(filePath string) error {
	var err error
	iv.CSVFile, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (iv *InsertValue) Do() error {
	file := csv.NewWriter(iv.CSVFile)
	if err := file.Write(iv.Values); err != nil {
		return err
	}

	file.Flush()
	if err := file.Error(); err != nil {
		return err
	}
	return nil
}

//func Do(path string) error {
//	filePath := filepath.Join(path, "cap.csv")
//	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
//	if err != nil {
//		return err
//	}
//
//	csvFile := csv.NewWriter(file)
//	if err := csvFile.Write(iv); err != nil {
//		return err
//	}
//
//	csvFile.Flush()
//	if err := csvFile.Error(); err != nil {
//		return err
//	}
//	return nil
//}
