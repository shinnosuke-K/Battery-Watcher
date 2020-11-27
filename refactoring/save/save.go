package save

import (
	"encoding/csv"
	"os"
	"path/filepath"
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

func (iv *InsertValue) CreateFile(path string, fileName string) error {
	filePath := filepath.Join(path, fileName)

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
