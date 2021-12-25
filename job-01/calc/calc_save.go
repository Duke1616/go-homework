package calculator

var suggests []string
var fatrates []float64
var bmis []float64
var names []string


func SaveHealthInformation(name string, bmi float64, fatRate float64, suggest string) {
	names = append(names, name)
	bmis = append(bmis, bmi)
	fatrates = append(fatrates, fatRate)
	suggests = append(suggests, suggest)
}