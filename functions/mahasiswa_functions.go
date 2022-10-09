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

func GetAllMahasiswa(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	nilai, err := query.GetAllMahasiswa(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(rw, nilai, http.StatusOK)

}

func PostMahasiswa(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var mhs models.Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	if err := query.InsertMahasiswa(ctx, mhs); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(rw, res, http.StatusCreated)
}

func UpdateMahasiswa(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var mhs models.Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	var idMhs = ps.ByName("id")

	if err := query.UpdateMahasiswa(ctx, mhs, idMhs); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(rw, res, http.StatusCreated)
}

func DeleteMahasiswa(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idMhs = ps.ByName("id")

	if err := query.DeleteMahasiswa(ctx, idMhs); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(rw, res, http.StatusCreated)
}
