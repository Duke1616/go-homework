package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	gobmi "github.com/Duke1616/go-bmi"
	"io"
	"learn.go/pkg/apis"
	"log"
	"os"
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
	_, err1 := cmd1.Output()
	if err1 != nil {
		log.Println(err1)
	}

	fp, err := os.OpenFile(m.filePath, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("读取json文件失败", err)
		return
	}
	defer fp.Close()

	for {
		br := bufio.NewReader(fp)
		_, _, c := br.ReadLine()
		if c == io.EOF {
			m.WriteFileSave(data)
			break
		}
		cmd2 := exec.Command(CommandShell, "-i", CommandCreate, m.filePath)
		_, err2 := cmd2.Output()
		if err2 != nil {
			log.Println(err2)
		}
		break
	}


	//cmd2 := exec.Command(CommandShell, "-i", CommandCreate, m.filePath)
	//bytes2, err2 := cmd2.Output()
	//if err2 != nil {
	//	log.Println(err2)
	//}
	//resp2 := string(bytes2)
	//log.Println(resp2)
}

func (m *Modification) WriteFileSave(data []byte) {
	file, err := os.OpenFile(m.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("无法打开文件", m.filePath, err)
	}

	defer file.Close()

	file.Write(append(data, '\n'))
}
