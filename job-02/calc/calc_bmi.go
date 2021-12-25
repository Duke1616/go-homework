package calc

import "fmt"

func CalcBMI(heightM float64, weightKG float64) (bmi float64, err error) {
	if heightM <= 0 {
		return 0, fmt.Errorf("身高不能是0或者负数")
	}
	if weightKG <= 0 {
		return 0, fmt.Errorf("体重不能是0或者负数")
	}
	return weightKG / (heightM * heightM), nil
}

