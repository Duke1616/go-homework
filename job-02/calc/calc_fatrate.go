package calc

import "fmt"

func FateRate(bmi float64, age int, sex string) (fatRate float64, err error) {
	if age <= 0 || age > 150 {
		return 0, fmt.Errorf("年龄不符合输入规范")
	}

	if bmi <= 0 {
		return 0, fmt.Errorf("bmi非法录入")
	}

	if sex == "男" {
		return (1.2*bmi + 0.22*float64(age) - 5.4 - 10.8*1) / 100, nil
	} else if sex == "女" {
		return (1.2*bmi + 0.22*float64(age) - 5.4 - 10.8*0) / 100, nil
	} else {
		return 0, fmt.Errorf("输入性别不是男也不是女")
	}
}