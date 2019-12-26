package main

import (
	"fmt"
	"goplayground/excelWriter"
	"strconv"
)

type MyData struct {
	Name      string
	Address   []Address
	Education Education
}

type Address struct {
	Pin      int
	State    string
	District string
	Taluka   string
}

type Education struct {
	CourseName string
	Fees       float64
}

func main() {

	data := getInputData()
	fmt.Println("DATA:", data)

	excelWriter.WriteExcel(data)

}

func getInputData() []MyData {
	data := []MyData{}
	for index := 0; index < 5; index++ {
		obj := MyData{}
		obj.Name = "Name" + strconv.Itoa(index)
		obj.Education.CourseName = "course" + strconv.Itoa(index)
		obj.Education.Fees = float64(index)

		for j := 0; j < 2; j++ {
			addrObj := Address{
				Pin:      j,
				State:    "State" + strconv.Itoa(j),
				District: "District" + strconv.Itoa(j),
				Taluka:   "Taluka" + strconv.Itoa(j),
			}

			obj.Address = append(obj.Address, addrObj)

		}
		data = append(data, obj)
	}

	return data
}
