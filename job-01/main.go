package main

import (
	"fmt"
	calc "learn.go/job-01/calc"
	"runtime/debug"
)

func main() {
	for {
		mainFatRateBody()
		if cont := whetherContinue(); !cont {
			break
		}
	}
}

func mainFatRateBody() {
	name, weight, tall, age, _, sex := getMaterialsFromInput()
	fatRate, bmi := calcFatRate(weight, tall, age, sex)

	if fatRate <= 0 {
		panic("fat rate is not allowed to be negative")
	}

	var suggest string
	if sex == "男" {
		suggest = getHealthinessSuggestions(age, fatRate, getHealthinessSuggestionsForMale)
	} else {
		suggest = getHealthinessSuggestions(age, fatRate, getHealthinessSuggestionForFemale)
	}

	calc.SaveHealthInformation(name, bmi, fatRate, suggest)
}

func getHealthinessSuggestions(age int, fatRate float64, getSuggestion func(age int, fatRate float64) string) string {
	return getSuggestion(age, fatRate)
}

func StateOfHealth(fatRate, a, b, c, d float64) string {
	if fatRate <= a {
		//fmt.Println("目前是：偏瘦")
		return "偏瘦"
	} else if fatRate > a && fatRate <= b {
		//fmt.Println("目前是：标准")
		return "标准"
	} else if fatRate > b && fatRate <= c {
		//fmt.Println("目前是：偏胖")
		return "偏胖"
	} else if fatRate > c && fatRate <= d {
		//fmt.Println("目前是：肥胖")
		return "肥胖"
	} else {
		//fmt.Println("目前是：非常肥胖")
		return "非常肥胖"
	}
}

func getHealthinessSuggestionForFemale(age int, fatRate float64) string {
	// 编写女性的体脂率与体制状态表
	if age >= 18 && age <= 39 {
		return StateOfHealth(fatRate, 0.2, 0.27, 0.34, 0.39)
	} else if age >= 40 && age <= 59 {
		return StateOfHealth(fatRate, 0.21, 0.28, 0.35, 0.40)
	} else {
		return StateOfHealth(fatRate, 0.22, 0.29, 0.36, 0.41)
	}
}

func getHealthinessSuggestionsForMale(age int, fatRate float64) string {
	// 编写男性的体脂率与体脂状态表
	if age >= 18 && age <= 39 {
		return StateOfHealth(fatRate, 0.1, 0.16, 0.21, 0.26)
	} else if age >= 40 && age <= 59 {
		return StateOfHealth(fatRate, 0.11, 0.17, 0.22, 0.27)
	} else {
		return StateOfHealth(fatRate, 0.13, 0.19, 0.24, 0.29)
	}
}

func calcFatRate(weight float64, tall float64, age int, sex string) (fatRate float64, bmi float64) {
	// 计算体脂率
	bmi = calc.CalcBMI(tall, weight)
	fatRate = calc.CalcFatRate(bmi, age, sex)
	fmt.Println("体脂率是: ", fatRate)
	return fatRate, bmi
}

func recoverMainBody() {
	if re := recover(); re != nil {
		fmt.Printf("warning: catch critical error: %v\n", re)
		debug.PrintStack()
	}
}

func getMaterialsFromInput() (string, float64, float64, int, int, string) {
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

	var sexWeight int
	sex := "男"
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)

	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	return name, weight, tall, age, sexWeight, sex
}

func whetherContinue() bool {
	fmt.Println("########  继续录入信息 1 ########")
	fmt.Println("########  打印录入信息 2 ########")
	var whetherContinue string
	fmt.Print("是做出你的选择: ")
	fmt.Scanln(&whetherContinue)
	if whetherContinue == "2" {
		calc.QueryAllHealthInformation()
		return false
	} else if whetherContinue == "1"{
		return true
	} else {
		return false
	}
	return true
}