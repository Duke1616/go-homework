package main

import (
	"fmt"
	"os"
)

type rankService struct {
	Write Writer
	rank  Ranker
}

func SelectMethod() *rankService {
	var method string
	fmt.Print("1: 冒泡排序  2: 快速排序  3: 退出程序: ")
	fmt.Scanln(&method)
	switch method {
	case "1":
		return &rankService{
			rank:  &RankBubbleSort{},
			Write: NewRecord("./person.json"),
		}
	case "2":
		return &rankService{
			rank:  &RankQuickSort{},
			Write: NewRecord("./person.json"),
		}
	case "3":
		os.Exit(2)
	default:
	}
	return nil
}
