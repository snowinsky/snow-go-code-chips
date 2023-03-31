package api

import (
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	"time"
)

func ExcelDemo() {
	fileName, err := writeFile()
	if err != nil {
		log.Println("write file error", err)
		return
	}
	readFile(fileName)
}

func writeFile() (string, error) {
	newExcelFile := excelize.NewFile()
	defer func() {
		if err := newExcelFile.Close(); err != nil {
			log.Println("close the excel file fail")
		}
	}()
	sheetIndex, err := newExcelFile.NewSheet("hebei")
	if err != nil {
		panic(errors.New("create sheet in excel file fail"))
	}
	newExcelFile.SetActiveSheet(sheetIndex)
	newExcelFile.SetCellBool("hebei", "B2", true)
	newExcelFile.SetCellStr("hebei", "B3", "true")
	newExcelFile.SetCellValue("hebei", "B4", "454548154545454545454")

	fileName := "B" + strconv.FormatInt(time.Now().Unix(), 16) + "Baoo1.xlsx"
	saveErr := newExcelFile.SaveAs(fileName)
	if saveErr != nil {
		return "", errors.Wrap(saveErr, "save fail")
	}
	return fileName, nil
}

func readFile(fileName string) {
	f, err := excelize.OpenFile(fileName)
	if nil != err {
		log.Println("open file fail", err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("close the open excel file error")
		}
	}()
	cellValue, err := f.GetCellValue("hebei", "B4")
	if nil != err {
		log.Println("get cell value fail", err)
		return
	}
	log.Println("get the cell value=", cellValue)
}
