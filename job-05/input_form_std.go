package main

import (
	"fmt"
	"learn.go/pkg/apis"
)

type inputFromStd struct {
}

func (inputFromStd) GetInput() *apis.PersonalInformation {
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var sex string
	fmt.Print("性別（男/女）：")
	fmt.Scanln(&sex)

	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)

	var weight float64
	fmt.Print("体重（公斤）：")
	fmt.Scanln(&weight)

	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)


	return &apis.PersonalInformation{
		Name:   name,
		Sex:    sex,
		Tall:   tall,
		Weight: weight,
		Age:    age,
	}
}
