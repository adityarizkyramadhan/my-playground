package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type DataBesar struct {
	Nama string
	Data string
	Info string
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	data := &DataBesar{
		Nama: "Aditya",
		Data: "Data banyak weee",
		Info: "Mantap brohhh",
	}
	key := "aditya"
	//set data to []byte
	dataByte, err := json.Marshal(data)
	fmt.Println("data byte: ", dataByte)
	if err != nil {
		panic(err)
	}
	client.Set(key, string(dataByte), 0)

	dataSave, err := client.Get(key).Result()
	fmt.Println("data save: ", dataSave)
	if dataSave == "" || err != nil {
		panic(err)
	}
	// perbandingan := fmt.Sprintf("get %s: ", key)
	// if strings.Contains(perbandingan, dataSave) {
	// 	dataSave = strings.Replace(dataSave, perbandingan, "", -1)
	// }
	fmt.Println("data save: ", dataSave)
	dataUnmarshal := &DataBesar{}
	err = json.Unmarshal([]byte(dataSave), dataUnmarshal)
	if err != nil {
		panic(err)
	}
	fmt.Println("data unmarshal nama : ", dataUnmarshal.Nama)

	/*
		Output :
		data byte:  [123 34 78 97 109 97 34 58 34 65 100 105 116 121 97 34 44 34 68 97 116 97 34 58 34 68 97 116 97 32 98 97 110 121 97 107 32 119 101 101 101 34 44 34 73 110 102 111 34 58 34 77 97 110 116 97 112 32 98 114 111 104 104 104 34 125]
		data save:  {"Nama":"Aditya","Data":"Data banyak weee","Info":"Mantap brohhh"}
		data save:  {"Nama":"Aditya","Data":"Data banyak weee","Info":"Mantap brohhh"}
		data unmarshal nama :  Aditya
	*/
}
