package main

import (
	"fmt"
	"sort"
	"time"
)

type Elevator struct {
	floor int	// 电梯当前层数
	runNum int  // 1 上行  2 下行
	floorArrUp []int
	floorArrDown []int
}

func (e *Elevator) RankElevator() {
	e.floor = 4
}

// 行进方向判断
func (e *Elevator) MovingDirection(p *Person){
	// 1 上行  2 下行
	if p.floorNumber[0] > e.floor {
		e.runNum = 1
	} else if p.floorNumber[0] < e.floor {
		e.runNum = 2
	}
}

// 存入数组
func (e *Elevator) SaveRankArr(p *Person) {
	sort.Slice(p.floorNumber, func(i, j int) bool {
		// fmt.Println(p.floorNumber[i], p.floorNumber[j])
		return p.floorNumber[i] < p.floorNumber[j]
	})
	for _,item := range p.floorNumber {
		if item > e.floor {
			e.floorArrUp = append(e.floorArrUp, item)
		} else {
			e.floorArrDown = append(e.floorArrDown, item)
		}
	}
}


func (e *Elevator) goRun(p *Person) {
	for {
		if len(e.floorArrUp) < 1 && len(e.floorArrDown) < 1 {
			break
		}
		if e.runNum == 1 {
			e.goRunUp()
		} else {
			e.goRunDown()
		}
	}
}

func (e *Elevator) goRunUp() {
	for index,_ := range e.floorArrUp {
		time.Sleep(1 * time.Second)
		e.floor = e.floorArrUp[index]
		if index + 1 == len(e.floorArrUp) {
			e.runNum = 2
			e.floorArrUp = e.floorArrUp[:0]
		}
		fmt.Println("到达楼层为： ", e.floor)
	}
}

func (e *Elevator) goRunDown() {
	sort.Sort(sort.Reverse(sort.IntSlice(e.floorArrDown)))
	for index,_ := range e.floorArrDown {
		time.Sleep(1 * time.Second)
		e.floor = e.floorArrDown[index]
		if index + 1  == len(e.floorArrDown) {
			e.runNum = 1
			e.floorArrDown = e.floorArrDown[:0]
		}
		fmt.Println("到达楼层为： ", e.floor)
	}
}