// 代码生成时间: 2025-09-01 02:38:55
package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"
    "github.com/tealeg/xlsx/v3"
)

// ExcelGenerator 定义一个Excel生成器结构体
type ExcelGenerator struct {
    SheetName string
    FileName  string
}

// NewExcelGenerator 创建一个ExcelGenerator实例
func NewExcelGenerator(sheetName, fileName string) *ExcelGenerator {
    return &ExcelGenerator{
        SheetName: sheetName,
        FileName:  fileName,
    }
}

// GenerateExcel 生成Excel文件
func (eg *ExcelGenerator) GenerateExcel(data [][]string) error {
    // 创建一个新的Excel文件
    file := xlsx.NewFile()
    // 创建一个工作表
    sheet, err := file.AddSheet(eg.SheetName)
    if err != nil {
        return err
    }

    // 添加数据到工作表
    for _, row := range data {
        if err := addRow(sheet, row); err != nil {
            return err
        }
    }

    // 保存Excel文件
    if err := file.Save(eg.FileName); err != nil {
        return err
    }

    return nil
}

// addRow 添加一行数据到工作表
func addRow(sheet *xlsx.Sheet, row []string) error {
    cell, err := sheet.AddRow()
    if err != nil {
        return err
    }
    for _, value := range row {
        cell.AddCell().Value = value
    }
    return nil
}

func main() {
    // 初始化Excel生成器
    generator := NewExcelGenerator("Example Sheet", "example.xlsx")

    // 准备数据
    data := [][]string{
        {"Header1", "Header2", "Header3"},
        {"Data1", "Data2", "Data3"},
        // ... 更多数据
    }

    // 生成Excel文件
    if err := generator.GenerateExcel(data); err != nil {
        fmt.Printf("Error generating Excel file: %v
", err)
    } else {
        fmt.Println("Excel file generated successfully.")
    }
}
