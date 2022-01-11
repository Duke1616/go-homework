package main

import "testing"

func TestCase01(t *testing.T) {
	e := Elevator{
		floor: 5,
	}
	if e.floor != 5 {
		t.Fatalf("预计电梯停留在5层, 但目前提留在%d", e.floor)
	}
}


func TestCase02(t *testing.T) {
	e := Elevator{
		floor: 1,
	}
	p := Person{}
	p.InputElevatorNumber(3)
	e.MovingDirection(&p)
	e.SaveRankArr(&p)
	e.goRun(&p)
	if e.floor != 3 {
		t.Fatalf("预计电梯停留在3层, 但目前提留在%d", e.floor)
	}
}

func TestCase03(t *testing.T) {
	e := Elevator{
		floor: 3,
	}
	p := Person{}
	p.InputElevatorNumber(4)
	p.InputElevatorNumber(2)
	e.MovingDirection(&p)
	e.SaveRankArr(&p)
	e.goRun(&p)
	if e.floor != 2 {
		t.Fatalf("预计电梯停留在2层, 但目前提留在%d", e.floor)
	}
}

func TestCase04(t *testing.T) {
	e := Elevator{
		floor: 3,
	}
	p := Person{}
	p.InputElevatorNumber(4)
	p.InputElevatorNumber(5)
	p.InputElevatorNumber(2)
	e.MovingDirection(&p)
	e.SaveRankArr(&p)
	e.goRun(&p)
	if e.floor != 2 {
		t.Fatalf("预计电梯停留在2层, 但目前提留在%d", e.floor)
	}
}




