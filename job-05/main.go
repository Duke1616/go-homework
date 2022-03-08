package main

import (
	"fmt"
	"learn.go/pkg/apis"
	"os"
)

var items []*apis.PersonalInformation

func main() {
	records := NewRecord("./person.json")
	records.initGlobalRankStandard()

	RecordModification := NewRecordModification("./person.json")
	rankSvc := SelectMethod()
	for {
		var choice string
		fmt.Print("1: 添加用户  2: 修改用户  3: 退出程序: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			input := &inputFromStd{}
			calc := &Calc{}
			pi := input.GetInput()
			rankSvc.Write.Write(pi)
			fatRate, _ := calc.FatRate(pi)
			rank := rankSvc.rank.GetRank(fatRate)
			fmt.Println("姓名: ", pi.Name, "排名：", rank, "体脂率: ", fatRate)
		case "2":
			RecordModification.ReadFileData()
			var serial string
			fmt.Println("根据编号选择修改的数据: ")
			fmt.Scanln(&serial)
			name, fatRate := RecordModification.UpdateFile(serial)
			rank := rankSvc.rank.GetRank(fatRate)
			fmt.Println("姓名: ", name, "排名：", rank, "体脂率: ", fatRate)
		case "3":
			os.Exit(1)
		default:
			continue
		}
	}
}
