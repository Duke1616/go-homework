package calculator

import (
	gobmi "github.com/Duke1616/go-bmi"
)

func CalcBMI(tall float64, weight float64) (bmi float64) {
	bmi,_ = gobmi.BMI(weight, tall)
	return
}
