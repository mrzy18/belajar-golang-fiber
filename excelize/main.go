package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type User map[string]any

var data = []User{
	{"Name": "John", "Gender": "male", "Age": 28},
	{"Name": "Nabila", "Gender": "female", "Age": 22},
	{"Name": "Budi", "Gender": "male", "Age": 23},
}

func main() {
	// Write File Xlsx
	xlsx := excelize.NewFile()

	sheet1Name := "Sheet One"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)
	xlsx.SetCellValue(sheet1Name, "A1", "Name")
	xlsx.SetCellValue(sheet1Name, "B1", "Gender")
	xlsx.SetCellValue(sheet1Name, "C1", "Age")

	err := xlsx.AutoFilter(sheet1Name, "A1", "C1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	for i, each := range data {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each["Name"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each["Gender"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), each["Age"])
	}

	err = xlsx.SaveAs("./file1.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	// Read File Xlsx
	xlsx, err = excelize.OpenFile("./file1.xlsx")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	rows := make([]User, 0)
	for i := 2; i < 5; i++ {
		row := User{
			"Name":   xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i)),
			"Gender": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i)),
			"Age":    xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i)),
		}
		rows = append(rows, row)
	}

	fmt.Printf("%v \n", rows)
}
