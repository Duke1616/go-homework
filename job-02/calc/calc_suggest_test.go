package calc

import (
	"testing"
)

// 测试健康建议是否符合预期
func TestCaseSuggest(t *testing.T) {
	// 男性健康建议
	{
		bmi, _ := CalcBMI(1.75, 75)
		fatRate, _ := FateRate(bmi, 23, "男")
		suggest := GetHealthinessSuggestions(23, fatRate, "男")
		t.Logf("预计得到偏胖, 实际得到：%s ", suggest)
		if suggest != "偏胖" {
			t.Logf("预计得到建议是偏旁, 但实际得到的是 %s", suggest)
		}
	}
	// 女性健康建议
	{
		bmi, _ := CalcBMI(1.75, 75)
		fatRate, _ := FateRate(bmi, 23, "女")
		suggest := GetHealthinessSuggestions(23, fatRate, "女")
		t.Logf("预计得到偏胖, 实际得到：%s ", suggest)
		if suggest != "偏胖" {
			t.Logf("预计得到建议是偏旁, 但实际得到的是 %s", suggest)
		}
	}
}

