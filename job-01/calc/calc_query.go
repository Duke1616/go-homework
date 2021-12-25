package calculator

import "fmt"

func QueryAllHealthInformation() {
	var avg_fatrates float64
	for index,_:= range names {
		fmt.Printf("姓名：%s  BMI: %f  体脂率：%f 建议: %s \n", names[index], bmis[index], fatrates[index], suggests[index])
		avg_fatrates += fatrates[index]
	}
	fmt.Printf("共录入%d人 平局体脂率为%f \n", len(names), avg_fatrates / float64(len(names)))
}