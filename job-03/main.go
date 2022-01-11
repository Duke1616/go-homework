package main


func main() {
	// TODO
	p := Person{}
	p.InputElevatorNumber(2)
	p.InputElevatorNumber(10)
	p.InputElevatorNumber(8)
	p.InputElevatorNumber(1)
	e := Elevator{}
	e.RankElevator()
	e.MovingDirection(&p)
	e.SaveRankArr(&p)
	e.goRun(&p)
}