package calc

func GetHealthinessSuggestions(age int, fatrate float64, sex string) (suggest string) {
	if sex == "男" {
		return getHealthinessSuggestionsForMale(age, fatrate)
	} else {
		return getHealthinessSuggestionForFemale(age, fatrate)
	}
}


func StateOfHealth(fatRate, a, b, c, d float64) string {
	if fatRate <= a {
		return "偏瘦"
	} else if fatRate > a && fatRate <= b {
		return "标准"
	} else if fatRate > b && fatRate <= c {
		return "偏胖"
	} else if fatRate > c && fatRate <= d {
		return "肥胖"
	} else {
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