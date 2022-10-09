package main

import (
	"fmt"
	"log"
	"net/http"
	"tugas15/functions"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/nilai-mahasiswa", functions.GetAllNilai)
	router.POST("/nilai-mahasiswa/create", functions.PostNilai)
	router.PUT("/nilai-mahasiswa/:id/update", functions.UpdateNilai)
	router.DELETE("/nilai-mahasiswa/:id/delete", functions.DeleteNilai)

	router.GET("/mahasiswa", functions.GetAllMahasiswa)
	router.POST("/mahasiswa/create", functions.PostMahasiswa)
	router.PUT("/mahasiswa/:id/update", functions.UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id/delete", functions.DeleteMahasiswa)

	router.GET("/mata-kuliah", functions.GetAllMataKuliah)
	router.POST("/mata-kuliah/create", functions.PostMataKuliah)
	router.PUT("/mata-kuliah/:id/update", functions.UpdateMataKuliah)
	router.DELETE("/mata-kuliah/:id/delete", functions.DeleteMataKuliah)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
