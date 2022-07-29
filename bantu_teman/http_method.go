package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataMahasiswa struct {
	Nama string `json:"nama"`
	Nim  string `json:"nim"`
}

func main() {
	var DataMahasiswas []DataMahasiswa
	router := gin.Default()
	router.GET("/mahasiswa", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
			"data":   DataMahasiswas,
		})
	})
	// router.GET("/mahasiswa", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"method": "GET DUA",
	// 	})
	// })
	router.POST("/mahasiswa", func(c *gin.Context) {
		var input DataMahasiswa
		c.ShouldBind(&input)
		DataMahasiswas = append(DataMahasiswas, input)
		c.JSON(http.StatusCreated, gin.H{
			"method": "POST",
			"data":   DataMahasiswas,
		})
	})
	router.PATCH("/mahasiswa", func(c *gin.Context) {
		var input DataMahasiswa
		c.ShouldBind(&input)
		for i := range DataMahasiswas {
			if DataMahasiswas[i].Nama == input.Nama {
				DataMahasiswas[i].Nim = input.Nim
				break
			}
		}
		DataMahasiswas = append(DataMahasiswas, input)
		c.JSON(http.StatusCreated, gin.H{
			"method": "PATCH",
			"data":   DataMahasiswas,
		})
	})
	router.DELETE("/mahasiswa", func(c *gin.Context) {
		var input DataMahasiswa
		var data []DataMahasiswa
		c.ShouldBind(&input)
		//delete data
		for i := range DataMahasiswas {
			if DataMahasiswas[i].Nama != input.Nama {
				data = append(data, DataMahasiswas[i])
			}
		}
		DataMahasiswas = data
		c.JSON(http.StatusCreated, gin.H{
			"method": "DELETE",
			"data":   DataMahasiswas,
		})
	})
	/*
		Jika ditemukan GET /mahasiswa  dan GET /mahasiswa dan method requestnya sama maka akan mucul panic: handlers are already registered for path '/mahasiswa'
		walaupun endpointnya sama tapi beda http method request maka akan beda handler
		Format : HTTP_METHOD_REQ ENDPOINT
		Misal : POST /mahasiswa GET /mahasiswa PATCH /mahasiswa DELETE /mahasiswa
		Dihindari : POST /tambahmahasiswa GET /lihatmahasiswa dan sebagainya
	*/
	router.Run(":5001")
}
