package excelprocessor

import (
	"encoding/csv"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"os"
	"strings"
	"sync"
)

// ExcelRowProcessor 定义了如何处理每一行数据的函数原型。
type ExcelRowProcessor func(row []string) []string

// ReadData 读取数据文件（支持 Excel 和 CSV）。
func ReadData(fileName string, rows int) ([][]string, error) {
	if strings.HasSuffix(strings.ToLower(fileName), ".xlsx") {
		return readExcel(fileName, rows)
	} else if strings.HasSuffix(strings.ToLower(fileName), ".csv") {
		return readCSV(fileName, rows)
	}
	return nil, fmt.Errorf("unsupported file type")
}

// readExcel 使用 excelize 读取 Excel 文件。
func readExcel(fileName string, rows int) ([][]string, error) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var data [][]string
	sheetName := f.GetSheetName(1) // 获取第一个工作表的名称
	rowsIterator, err := f.Rows(sheetName)
	if err != nil {
		return nil, err
	}

	rowCount := 0
	for rowsIterator.Next() {
		if rows != -1 && rowCount >= rows {
			break
		}
		row, err := rowsIterator.Columns()
		if err != nil {
			return nil, err
		}
		data = append(data, row)
		rowCount++
	}
	return data, nil
}

// readCSV 专门读取 CSV 文件。
func readCSV(fileName string, rows int) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var data [][]string
	for i := 0; rows == -1 || i < rows; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		data = append(data, record)
	}
	return data, nil
}

// ProcessData 并发处理数据。
func ProcessData(data [][]string, maxConcurrency int, processor ExcelRowProcessor) [][]string {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxConcurrency)
	results := make([][]string, len(data))

	for i, row := range data {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(row []string, index int) {
			defer wg.Done()
			result := processor(row)
			results[index] = result
			<-semaphore
		}(row, i)
	}

	wg.Wait()
	return results
}

// WriteCSV 将结果写入到CSV文件中。
func WriteCSV(fileName string, data [][]string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	return nil
}
