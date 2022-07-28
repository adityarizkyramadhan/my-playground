package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ModelTest struct {
	gorm.Model
	Name string `json:"name"`
	Data string `json:"data"`
}

type ModelInput struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func main() {

	router := gin.Default()
	dnsXAMPP := "root:@tcp(127.0.0.1:3306)/playground_database?parseTime=true"
	db, err := gorm.Open(mysql.Open(dnsXAMPP), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&ModelTest{})
	if err != nil {
		panic(err)
	}
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.GET("/alldata", func(c *gin.Context) {
		var test []*ModelTest
		if err := db.Find(&test).Error; err != nil {
			panic(err)
		}
		//
		// obj := []*ModelTest{}
		// err = redis.ScanStruct(value, &obj)
		// if err != nil {
		// 	panic(err)
		// }
		c.JSON(200, gin.H{
			"message": test,
		})
	})
	router.GET("/search", func(c *gin.Context) {
		name := c.Query("name")
		conn := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		value, _ := conn.Get(name).Result()
		if value == "" {
			var test *ModelTest
			fmt.Println("Masuk Querry")
			if err := db.Where("name = ?", name).First(&test).Error; err != nil {
				panic(err)
			}
			conn.Set(name, test.Data, 0)
		} else {
			fmt.Println("Masuk Redis")
		}
		data := conn.Get(name).Val()
		c.JSON(200, gin.H{
			"message": data,
		})
		/*
			Sebagai penanda bahwa query pertama akan menquery database
			Masuk Querry
			[GIN] 2022/07/28 - 11:53:43 | 200 |     16.2775ms |             ::1 | GET      "/search?name=RIARIO"
			Query kedua akan mengambil data dari redis
			Masuk Redis
			[GIN] 2022/07/28 - 11:53:48 | 200 |     10.3964ms |             ::1 | GET      "/search?name=RIARIO"

		*/
	})
	router.POST("/add", func(c *gin.Context) {
		var input ModelInput
		c.ShouldBindJSON(&input)
		data := &ModelTest{
			Name: input.Name,
			Data: input.Data,
		}
		db.Create(data)
		c.JSON(201, gin.H{
			"data": data,
		})
	})
	router.Run(":8080")
}
