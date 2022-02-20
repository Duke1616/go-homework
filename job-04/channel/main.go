package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

type favContext string

type RankItem struct {
	Name	     string
	FatRate      float64
}



type FatRateRank struct {
	items		 []RankItem
	wg           sync.WaitGroup
	readNameCh	 chan string
	getUserCh	 chan string
	updateCh	 chan string
	PersonNumber int
	min			 float64
	max			 float64
}


func (f *FatRateRank) CreateUser(n int) {
	f.readNameCh <- fmt.Sprintf("TempUser%s", strconv.Itoa(n))
}


func (f *FatRateRank) UserRegister() {
	for {
		select {
		case oName, ok := <- f.readNameCh:
			if !ok {
				time.Sleep(2 * time.Second)
				for _, item := range f.items {
					f.updateCh <- item.Name
				}
			} else {
				FatRate := f.min + rand.Float64() * (f.max - f.min)
				f.items = append(f.items, RankItem {
					Name:    oName,
					FatRate: FatRate,
				})
				fmt.Println("用户：", oName, "注册成功")
			}
		default:
		}
	}
}

func (f *FatRateRank) GetUserFatRateRank() {
	for {
		select {
		case oName := <- f.getUserCh:
			rank, FatRate := f.getRank(oName)
			fmt.Println("UpdateUserFatRate", " 用户", oName, " 排名:", rank, " 体脂率:", FatRate)
		default:
			time.Sleep(2 * time.Second)
			for _, item := range f.items {
				rank, FatRate := f.getRank(item.Name)
				fmt.Println("GetUserFatRateRank", " 用户", item.Name, " 排名:", rank, " 体脂率:", FatRate)
			}
		}
	}
}

func (f *FatRateRank) UpdateUserFatRate() {
	for oName := range f.updateCh {
		NewFatRate := f.min + rand.Float64() * (f.max - f.min)
		for index, item := range f.items {
			if oName == item.Name {
				if item.FatRate < NewFatRate && item.FatRate + 0.2 >= NewFatRate {
					item.FatRate = NewFatRate
					f.items[index] = item
				} else if item.FatRate > NewFatRate && NewFatRate + 0.2 >= item.FatRate {
					item.FatRate = NewFatRate
					f.items[index] = item
				} else {
					fmt.Println(item.Name, item.FatRate, NewFatRate, "不符合体脂修改范围")
					break
				}
				f.getUserCh <- oName
			}
		}
	}
}


func goRun() {
	f := &FatRateRank{
		readNameCh:   make(chan string, 1000),
		updateCh:     make(chan string),
		getUserCh:    make(chan string),
		PersonNumber: 1000,
		min:          0,
		max:          0.4,
	}

	f.wg.Add(f.PersonNumber)
	for i := 0; i < f.PersonNumber; i++ {
		go func(n int) {
			defer f.wg.Done()
			f.CreateUser(n)
		}(i)
	}

	go f.UserRegister()
	go f.GetUserFatRateRank()
	worker := 10

	for i:= 0; i< worker; i ++ {
		go f.UpdateUserFatRate()
	}


	f.wg.Wait()
	close(f.readNameCh)

	for{
	}
}

func main()  {
	goRun()
}

func (f *FatRateRank) getRank(name string) (rank int, fatRate float64) {
	frs := map[float64]struct{}{}
	for _, item := range f.items {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fatRate = item.FatRate
		}
	}
	rankArr := make([]float64, 0, len(frs))
	for k := range frs {
		rankArr = append(rankArr, k)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			break
		}
	}
	return rank, fatRate
}

