package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	gobmi "github.com/Duke1616/go-bmi"
	"io"
	"learn.go/pkg/apis"
	"os"
	"sync"
)

func NewRecord(filePath string) *WriteFile {
	return &WriteFile{
		filePath: filePath,
	}
}

type WriteFile struct {
	filePath string
}

type WriteMysql struct {
}

func (w *WriteFile) Write(pi *apis.PersonalInformation) {
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
	w.WriteFileSave(data)

	items = append(items, &apis.PersonalInformation{
		Name:    pi.Name,
		Sex:     pi.Sex,
		Tall:    pi.Tall,
		Weight:  pi.Weight,
		Age:     pi.Age,
		Bmi:     bmi,
		FatRate: fatRate,
	})
}

func (w *WriteFile) WriteFileSave(data []byte) {
	file, err := os.OpenFile(w.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("无法打开文件", w.filePath, err)
	}

	defer file.Close()

	file.Write(append(data, '\n'))
}

func (w *WriteMysql) Write() {
	// TODO
}

func (w *WriteFile) initGlobalRankStandard() {
	init := sync.Once{}
	init.Do(func() {
		fp, err := os.OpenFile(w.filePath, os.O_RDONLY, 0755)
		if err != nil {
			fmt.Println("读取json文件失败", err)
			return
		}

		defer fp.Close()
		br := bufio.NewReader(fp)
		serialize := &apis.PersonalInformation{}

		for {
			JsonPerson, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			err = json.Unmarshal(JsonPerson[:len(JsonPerson)], serialize)
			if err != nil {
				fmt.Println("解析数据失败", err)
				return
			}
			//fmt.Printf("%+v\n", serialize)
			items = append(items, &apis.PersonalInformation{
				Name:    serialize.Name,
				Sex:     serialize.Sex,
				Tall:    serialize.Tall,
				Weight:  serialize.Weight,
				Age:     serialize.Age,
				Bmi:     serialize.Bmi,
				FatRate: serialize.FatRate,
			})
		}
	})
}
