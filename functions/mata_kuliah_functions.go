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

func GetAllMataKuliah(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	nilai, err := query.GetAllMataKuliah(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(rw, nilai, http.StatusOK)

}

func PostMataKuliah(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var matkul models.MataKuliah
	if err := json.NewDecoder(r.Body).Decode(&matkul); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	if err := query.InsertMataKuliah(ctx, matkul); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(rw, res, http.StatusCreated)
}

func UpdateMataKuliah(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var matkul models.MataKuliah
	if err := json.NewDecoder(r.Body).Decode(&matkul); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	var idmatkul = ps.ByName("id")

	if err := query.UpdateMataKuliah(ctx, matkul, idmatkul); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(rw, res, http.StatusCreated)
}

func DeleteMataKuliah(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idmatkul = ps.ByName("id")

	if err := query.DeleteMataKuliah(ctx, idmatkul); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(rw, res, http.StatusCreated)
}
