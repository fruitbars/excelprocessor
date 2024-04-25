


# ExcelProcessor

`ExcelProcessor` 是一个用于高效读取和处理 Excel (.xlsx) 和 CSV 文件的 Go 语言库。它利用 `excelize` 和 `encoding/csv` 标准库的功能，提供了一套简单的 API 来并行处理大量数据并将结果输出到 CSV 文件中。

## 功能

- 读取 Excel (.xlsx) 和 CSV 文件。
- 可配置的行数处理，支持处理全部数据。
- 支持并行数据处理。
- 结果输出到 CSV 文件。

## 安装

确保你已安装 Go 1.13 或更高版本。通过以下命令安装 `ExcelProcessor`：

```bash
go get github.com/fruitbars/excelprocessor
```

确保同时安装了 `excelize` 库：

```bash
go get github.com/xuri/excelize/v2
```

## 快速开始

下面是如何使用 `ExcelProcessor` 来读取文件、处理数据并输出到 CSV 的简单示例。

### 示例代码

```go
package main

import (
	"log"
	"excelprocessor"
)

func main() {
	data, err := excelprocessor.ReadData("path/to/your/file.xlsx", -1) // 支持 .xlsx 或 .csv 文件
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
```

### API 参考

#### `ReadData(fileName string, rows int) ([][]string, error)`

读取指定的 Excel 或 CSV 文件。`rows` 参数指定要处理的最大行数；如果设置为 `-1`，则处理所有行。

#### `ProcessData(data [][]string, maxConcurrency int, processor ExcelRowProcessor) [][]string`

并发处理数据。`maxConcurrency` 参数控制并发处理的最大数量。

#### `WriteCSV(fileName string, data [][]string) error`

将处理后的数据写入 CSV 文件。

## 贡献

欢迎通过 GitHub 提交问题报告和拉取请求。

## 许可证

本项目使用 MIT 许可证。有关详细信息，请参阅 LICENSE 文件。