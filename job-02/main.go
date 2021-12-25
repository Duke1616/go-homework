package main

import (
	"fmt"
	calc "learn.go/job-02/calc"
)


func main() {
	weight, tall, age, sex := getMaterialsFromInput()
	bmi, err := calc.CalcBMI(tall, weight)
	if err != nil {
		panic(err)
	}

	fatRate,err  := calc.FateRate(bmi, age, sex)
	if err != nil {
		panic(err)
	}

	fmt.Println(calc.GetHealthinessSuggestions(age, fatRate, sex))

}



func getMaterialsFromInput() (float64, float64, int, string) {
	// 录入各项
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重（千克）：")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)
	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	var sex string
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)


	return weight, tall, age, sex
}