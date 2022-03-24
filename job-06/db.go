package main

import (
	gobmi "github.com/Duke1616/go-bmi"
	"gorm.io/gorm"
	"log"
)

var _ CircleOfFriends = &dbCircle{}

func NewDbCircle(conn *gorm.DB) CircleOfFriends {
	return &dbCircle{
		conn: conn,
	}
}

type dbCircle struct {
	conn *gorm.DB
}

func (d *dbCircle) PostState(pi *PersonalInformation) error {
	bmi, _ := gobmi.BMI(pi.Weight, pi.Tall)
	fatRate := gobmi.CalcFatRate(bmi, pi.Age, pi.Sex)
	newPi := &PersonalInformation{
		Desc:    pi.Desc,
		Name:    pi.Name,
		Sex:     pi.Sex,
		Tall:    pi.Tall,
		Weight:  pi.Weight,
		Age:     pi.Age,
		Bmi:     bmi,
		FatRate: fatRate,
	}
	resp := d.conn.Create(newPi)
	if err := resp.Error; err != nil {
		log.Printf("创建%s时失败：%v\n", pi.Name, err)
		return err
	}
	log.Printf("创建%s成功\n", pi.Name)
	return nil
}

func (d dbCircle) DeleteState(id string) error {
	// gorm软删除

	resp := d.conn.Where("id = ?", id).Delete(&PersonalInformation{})
	if err := resp.Error; err != nil {
		log.Printf("更新%s时失败：%v\n", id, err)
		return err
	}
	log.Printf("更新%s成功\n", id)
	return nil
}

func (d dbCircle) GetState() ([]*CircleOfFriendsInfo, error) {
	var c []*CircleOfFriendsInfo
	resp := d.conn.Raw("select * from personal_info where deleted_at is null;").Scan(&c)
	if err := resp.Error; err != nil {
		log.Printf("获取时失败：%v\n", err)
		return nil, err
	}
	log.Printf("获取成功\n")
	return c, nil
}

func (d dbCircle) DeleteAll() error {
	var c PersonalInformation
	// TODO 无法删除 不知道原因
	resp := d.conn.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&c)
	if err := resp.Error; err != nil {
		log.Printf("删除所有失败：%v\n", err)
		return err
	}
	log.Printf("删除所有成功\n")
	return nil
}
