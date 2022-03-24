package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("golang:123456@tcp(39.97.251.109:3306)/circle?charset=utf8mb4&parseTime=true&loc=Local"))
	conn.AutoMigrate(&PersonalInformation{})
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func main() {
	conn := connectDb()
	var circleServer CircleOfFriends = NewDbCircle(conn)
	r := gin.Default()
	r.POST("/poststate", func(c *gin.Context) {
		var pi *PersonalInformation
		if err := c.BindJSON(&pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "无法解析注册信息" + err.Error(),
			})
			return
		}
		if err := circleServer.PostState(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "注册失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	r.DELETE("/deletestate/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := circleServer.DeleteState(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "删除用户失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		}
	})
	r.GET("getstate", func(c *gin.Context) {
		if info, err := circleServer.GetState(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "获取用户失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, info)
		}
	})
	// TODO
	r.DELETE("/deleteall", func(c *gin.Context) {
		if err := circleServer.DeleteAll(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errorMessage": "删除失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		}
	})
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
