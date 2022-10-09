package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"tugas15/models"
	"tugas15/query"
	"tugas15/utils"

	"github.com/julienschmidt/httprouter"
)

func getIndeksNilai(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70 && nilai < 80:
		return "B"
	case nilai >= 60 && nilai < 70:
		return "C"
	case nilai >= 50 && nilai < 60:
		return "D"
	default:
		return "E"
	}
}

func GetAllNilai(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	nilai, err := query.GetAllNilai(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(rw, nilai, http.StatusOK)

}

func PostNilai(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var nilaiMhs models.NilaiMahasiswa
	if err := json.NewDecoder(r.Body).Decode(&nilaiMhs); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	if nilaiMhs.Skor > 100 {
		http.Error(rw, "Nilai tidak boleh diinput lebih dari 100", http.StatusBadRequest)
		return
	}
	nilaiMhs.Indeks = getIndeksNilai(nilaiMhs.Skor)

	if err := query.InsertNilai(ctx, nilaiMhs); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(rw, res, http.StatusCreated)
}

func UpdateNilai(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var nilaiMhs models.NilaiMahasiswa
	if err := json.NewDecoder(r.Body).Decode(&nilaiMhs); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	var idNilaiMhs = ps.ByName("id")

	if nilaiMhs.Skor > 100 {
		http.Error(rw, "Nilai tidak boleh diinput lebih dari 100", http.StatusBadRequest)
		return
	}
	nilaiMhs.Indeks = getIndeksNilai(nilaiMhs.Skor)

	if err := query.UpdateNilai(ctx, nilaiMhs, idNilaiMhs); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(rw, res, http.StatusCreated)
}

func DeleteNilai(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idNilaiMhs = ps.ByName("id")

	if err := query.DeleteNilai(ctx, idNilaiMhs); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(rw, res, http.StatusCreated)
}
