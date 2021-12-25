package calc

import (
	"testing"
)

// BMI单侧编写
func TestCaseBmi(t *testing.T) {
	// 正常BMI计算
	{
		bmi, _ := CalcBMI(1.75, 75.0)
		if bmi != 24.489795918367346 {
			t.Fatalf("预计BMI等于，实际得到的是%f", bmi)
		}
	}
	// 测试身高为0触犯Error
	{
		bmi, _ := CalcBMI(0, 75.0)
		if bmi != 0 {
			t.Fatalf("预计BMI等于0，实际得到的是 %f", bmi)
		}
	}
	// 测试体重为0触犯Error
	{
		bmi, _ := CalcBMI(1.70, 0)
		if bmi != 0 {
			t.Fatalf("预计BMI等于0，实际得到的是 %f", bmi)
		}
	}
}