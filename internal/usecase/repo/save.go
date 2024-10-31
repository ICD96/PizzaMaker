package repo

import (
	"bytes"
	"controller/internal/entity"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func Save(orders *entity.OrderMap) {
	writer, file, err := createCSVWriter("data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	records := [][]string{}
	temp := ""
	for k, v := range orders.O {
		for i := range v.OrderedPizza {
			temp = temp + string(i)
		}
		records = append(records, []string{string(k), string(v.Id), v.TimeCreated.String(), v.CreatorName, v.Status, string(v.OrderCost), string(v.TimeToPrepare), temp})
	}

	for _, record := range records {
		writer.Write(record)
	}
	writer.Flush()
}

func createCSVWriter(filename string) (*csv.Writer, *os.File, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, nil, err
	}
	writer := csv.NewWriter(f)
	return writer, f, nil
}

func Read() {
	data, err := readCVSFile("data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := parseCSV(data)
	processCSV(reader)
}

func readCVSFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func parseCSV(data []byte) *csv.Reader {
	reader := csv.NewReader(bytes.NewReader(data))
	return reader
}

func processCSV(reader *csv.Reader) {
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading CSV data:", err)
			break
		}
		fmt.Println(record)
	}
}
