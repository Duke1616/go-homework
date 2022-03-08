package main

import (
	"encoding/json"
	"fmt"
	gobmi "github.com/Duke1616/go-bmi"
	"learn.go/pkg/apis"
	"log"
	"os/exec"
	"strconv"
)

func NewRecordModification(filePath string) *Modification {
	return &Modification{
		filePath: filePath,
	}
}

type Modification struct {
	filePath string
}

func (Modification) ReadFileData() error {
	if len(items) == 0 {
		err := fmt.Errorf("没有用户注册，请先注册用户")
		return err
	} else {
		for i, v := range items {
			fmt.Println(i+1, v)
		}
	}
	return nil
}

func (m *Modification) UpdateFile(serial string) (name string, fatRateRank float64) {
	intSerial, _ := strconv.Atoi(serial)
	if intSerial > len(items) {
		fmt.Println("所选择编号不存在，请重新操作")
	} else {
		input := &inputFromStd{}
		pi := input.GetInput()

		bmi, _ := gobmi.BMI(pi.Weight, pi.Tall)
		fatRate := gobmi.CalcFatRate(bmi, pi.Age, pi.Sex)

		personAll := &apis.PersonalInformation{
			Name:    pi.Name,
			Sex:     pi.Sex,
			Tall:    pi.Tall,
			Weight:  pi.Weight,
			Age:     pi.Age,
			Bmi:     bmi,
			FatRate: fatRate,
		}

		data, _ := json.Marshal(personAll)
		m.ShellCommand(serial, data)

		items[intSerial-1].Sex = pi.Sex
		items[intSerial-1].Name = pi.Sex
		items[intSerial-1].Tall = pi.Tall
		items[intSerial-1].Weight = pi.Weight
		items[intSerial-1].Age = pi.Age
		items[intSerial-1].Bmi = bmi
		items[intSerial-1].FatRate = fatRate

		return pi.Name, fatRate
	}
	return "", 0
}

func (m *Modification) ShellCommand(serial string, data []byte) {
	CommandShell := "/usr/bin/sed"

	ArgsDelete := "d"
	CommandDelete := serial + ArgsDelete

	ArgsCreate := "i"
	CommandCreate := serial + ArgsCreate + " " + string(data)
	cmd1 := exec.Command(CommandShell, "-i", CommandDelete, m.filePath)
	cmd2 := exec.Command(CommandShell, "-i", CommandCreate, m.filePath)
	bytes1, err := cmd1.Output()
	if err != nil {
		log.Println(err)
	}
	resp1 := string(bytes1)
	log.Println(resp1)

	bytes2, err := cmd2.Output()
	if err != nil {
		log.Println(err)
	}
	resp2 := string(bytes2)
	log.Println(resp2)
}
