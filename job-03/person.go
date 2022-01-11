package main


type Person struct {
	floorNumber []int
}


func (p *Person) InputElevatorNumber(elevatorNumber int) {
	p.floorNumber = append(p.floorNumber, elevatorNumber)
}
