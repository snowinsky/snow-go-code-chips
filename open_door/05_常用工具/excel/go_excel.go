package main

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"log"
)

func main() {
	//创建excel文件
	xlsx := excelize.NewFile()

	//创建新表单
	index, _ := xlsx.NewSheet("成绩表")

	//写入数据
	data := map[string]string{
		//学科
		"B1": "语文",
		"C1": "数学",
		"D1": "英语",
		"E1": "理综",

		//姓名
		"A2": "啊俊",
		"A3": "小杰",
		"A4": "老王",

		//啊俊成绩
		"B2": "112",
		"C2": "115",
		"D2": "128",
		"E2": "255",

		//小杰成绩
		"B3": "100",
		"C3": "90",
		"D3": "110",
		"E3": "200",

		//老王成绩
		"B4": "70",
		"C4": "140",
		"D4": "60",
		"E4": "265",
	}
	for k, v := range data {
		//设置单元格的值
		xlsx.SetCellValue("成绩表", k, v)
	}

	//设置默认打开的表单
	xlsx.SetActiveSheet(index)

	//保存文件到指定路径
	err := xlsx.SaveAs("./成绩表.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	//#########################################
	f, err := excelize.OpenFile("./成绩表.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	/*
	   //读取某个单元格的值
	   value, err := f.GetCellValue("成绩表", "D2")
	   if err != nil {
	       log.Fatal(err)
	   }
	   fmt.Println(value)
	*/

	//读取某个表单的所有数据
	rows, err := f.GetRows("成绩表")
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range rows {
		for _, value := range row {
			fmt.Printf("\t%s", value)
		}
		fmt.Println()
	}
}
