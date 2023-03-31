package api

import (
	"github.com/xuri/excelize/v2"
	"log"
)

func ExcelDemo() {
	newExcelFile := excelize.NewFile()
	defer func() {
		if err := newExcelFile.Close(); err != nil {
			log.Println("close the excel file fail")
		}
	}()
}
