package main

import (
	"gorm.io/gorm"
	"time"
)

type PersonalInformation struct {
	gorm.Model
	Id        int64      `json:"id" gorm:"primaryKey;column:id"`
	Desc      string     `json:"desc" gorm:"column:desc"`
	Name      string     `json:"name" gorm:"column:name"`
	Sex       string     `json:"sex" gorm:"column:sex"`
	Tall      float64    `json:"tall" gorm:"column:tall"`
	Weight    float64    `json:"weight" gorm:"column:weight"`
	Age       int        `json:"age" gorm:"column:age"`
	Bmi       float64    `json:"bmi" gorm:"column:bmi"`
	FatRate   float64    `json:"fat_rate" gorm:"column:fat_rate"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

type CircleOfFriendsInfo struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        int64     `json:"id"`
	Desc      string    `json:"desc"`
	Name      string    `json:"name"`
	Sex       string    `json:"sex"`
	Tall      float64   `json:"tall"`
	Weight    float64   `json:"weight"`
	Age       int       `json:"age"`
	FatRate   float64   `json:"fat_rate"`
}

func (*PersonalInformation) TableName() string {
	return "personal_info"
}
