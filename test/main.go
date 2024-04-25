package main

import (
	"excelprocessor"
	"log"
)

func main() {
	data, err := excelprocessor.ReadData("path/to/your/file.xlsx", -1) // 可以是 Excel 或 CSV
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	processedData := excelprocessor.ProcessData(data, 10, YourCustomProcessor)

	err = excelprocessor.WriteCSV("output.csv", processedData)
	if err != nil {
		log.Fatalf("Failed to write CSV file: %s", err)
	}
}

// YourCustomProcessor 自定义处理函数
func YourCustomProcessor(row []string) []string {
	// 实现你的逻辑
	return []string{row[0], "Processed"}
}
