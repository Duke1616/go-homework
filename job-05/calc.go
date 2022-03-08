package main

import (
	gobmi "github.com/Duke1616/go-bmi"
	"learn.go/pkg/apis"
)

type Calc struct {
}

func (Calc) BMI(person *apis.PersonalInformation) (float64, error) {
	bmi, err := gobmi.BMI(person.Weight, person.Tall)
	if err != nil {
		return -1, err
	}
	return bmi, err
}

func (c *Calc) FatRate(person *apis.PersonalInformation) (float64, error) {
	bmi, err := c.BMI(person)
	if err != nil {
		return -1, err
	}
	return gobmi.CalcFatRate(bmi, person.Age, person.Sex), nil
}
