package main

import (
	"errors"

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

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       3,
	})

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
		// masih gagal karena speed querry terlalu cepat
		client.Set("alldata", test, 1*60*60)
		data := client.Get("alldata")
		if data == nil {
			panic(errors.New("data in redis was not set"))
		}
		c.JSON(200, gin.H{
			"message": data,
			"data":    test,
		})
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
