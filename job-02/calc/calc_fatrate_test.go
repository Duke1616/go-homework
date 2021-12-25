package calc

import (
	"testing"
)

// 体脂率单侧编写
func TestCaseFatrate(t *testing.T) {
	// 测试男性体脂率
	{
		bmi, _ := CalcBMI(1.75, 75)

		fatRate, _ := FateRate(bmi, 23, "男")

		if fatRate != 0.18247755102040816 {
			t.Fatalf("预计体脂率为，但实际得到的是 %f", fatRate)
		}
	}
	// 测试女性体脂率
	{
		bmi, _ := CalcBMI(1.75, 75)
		fatRate, _ := FateRate(bmi, 23, "女")

		if fatRate != 0.2904775510204082 {
			t.Fatalf("预计体脂率为，但实际得到的是 %f", fatRate)
		}
	}
	// 测试性别非男非女
	{
		bmi, _ := CalcBMI(1.75, 75)
		fatRate, err := FateRate(bmi, 23, "中")
		t.Logf("实际得到：%f, error: %v", fatRate, err)
		if err == nil {
			t.Fatalf("预计体脂率计算报错，但正常返回 %f", fatRate)
		}
	}
	// 测试年龄大于150岁
	{
		bmi, _ := CalcBMI(1.75, 75)
		fatRate, err := FateRate(bmi, 151, "男")
		t.Logf("实际得到：%f, error: %v", fatRate, err)
		if err == nil {
			t.Fatalf("预计体脂率计算报错，但正常返回 %f", fatRate)
		}
	}
	// 测试年龄为负数
	{
		bmi, _ := CalcBMI(1.75, 75)
		fatRate, err := FateRate(bmi, -3, "男")
		t.Logf("实际得到：%f, error: %v", fatRate, err)
		if err == nil {
			t.Fatalf("预计体脂率为，但实际得到的是 %f", fatRate)
		}
	}

	// 测试BMI非法录入
	{
		bmi, _ := CalcBMI(0, 75)
		fatRate, err := FateRate(bmi, 23, "男")
		t.Logf("实际得到：%f, error: %v", fatRate, err)
		if fatRate < 0 {
			t.Fatalf("输入bmi异常 %f", fatRate)
		}
	}

}



