package excelWriter

import (
	"fmt"
	"reflect"

	"github.com/tealeg/xlsx"
)

func WriteExcel(data interface{}) error {
	excelFilePtr := xlsx.NewFile()
	sheet, err := excelFilePtr.AddSheet("My Data")
	if err != nil {
		fmt.Println("Erro while creating sheet:", err)
		return err
	}

	// err = excelFilePtr.Save("./test.xlsx")
	// if err != nil {
	// 	fmt.Println("eror:", err)
	// 	return err
	// }

	processData(data, sheet)
	return nil
}

func processData(data interface{}, sheet *xlsx.Sheet) {

	switch reflect.TypeOf(data).Kind() {

	case reflect.Slice:
		s := reflect.ValueOf(data)

		for i := 0; i < s.Len(); i++ {
			for j := 0; j < s.Index(i).NumField(); j++ {
				varName := s.Index(i).Type().Field(j).Name
				varType := s.Index(i).Type().Field(j).Type
				varValue := s.Index(i).Field(j).Interface()
				fmt.Printf("%v %v %v\n", varName, varType, varValue)
				fmt.Println("Datatype:", reflect.TypeOf(s.Index(i).Field(j).Interface()).Kind())
			}

			
		}

	default:
		fmt.Println("Invalid type")
	}
}
