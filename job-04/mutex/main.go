package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

type RankItem struct {
	Name	     string
	FatRate      float64
}

type FatRateRank struct {
	items		 []RankItem
	lock		 sync.RWMutex
	PersonNumber int
	min			 float64
	max			 float64
}


func (f *FatRateRank) UserRegister(i int) {
	f.lock.Lock()
	defer f.lock.Unlock()
	oName := fmt.Sprintf("TempUser%s", strconv.Itoa(i))
	FatRate := f.min + rand.Float64() * (f.max - f.min)
	f.items = append(f.items, RankItem{
		Name:    oName,
		FatRate: FatRate,
	})
	fmt.Println(oName, "用户注册成功")
}

func (f *FatRateRank) GetUserFatRateRank(i int) {
	f.lock.RLock()
	defer f.lock.RUnlock()
	oName := fmt.Sprintf("TempUser%s", strconv.Itoa(i))
	for _, item := range f.items {
		if oName == item.Name {
			rank, FatRate := f.getRank(oName)
			fmt.Println("GetUserFatRateRank", " 用户", oName, " 排名:", rank, " 体脂率:", FatRate)
		}
	}
}

func (f *FatRateRank) UpdateUserFatRate(i int) {
	f.lock.Lock()
	defer f.lock.Unlock()
	oName := fmt.Sprintf("TempUser%s", strconv.Itoa(i))
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
			rank, FatRate := f.getRank(oName)
			fmt.Println("UpdateUserFatRate", " 用户", oName, " 排名:", rank, " 体脂率:", FatRate)
		}
	}
}


func goRun() {
	f := &FatRateRank{
		PersonNumber: 1000,
		min:          0,
		max:          0.4,
	}

	for i:=0; i<f.PersonNumber; i++{
		go func(i int) {
			f.UserRegister(i)
		}(i)
	}


	for {
		for i:=0; i < f.PersonNumber; i++{
			go func(i int) {
				f.GetUserFatRateRank(i)
				f.UpdateUserFatRate(i)
			}(i)
		}
		time.Sleep(2 * time.Second)
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
