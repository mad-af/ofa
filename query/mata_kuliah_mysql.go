package query

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
	"tugas15/config"
	"tugas15/models"
)

const (
	matkul_table = "mata_kuliah"
)

func GetAllMataKuliah(ctx context.Context) ([]models.MataKuliah, error) {
	var matkuls []models.MataKuliah
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := "SELECT * FROM " + matkul_table + " Order By created_at DESC"
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var matkul models.MataKuliah
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&matkul.ID,
			&matkul.Nama,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		matkul.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		matkul.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		matkuls = append(matkuls, matkul)
	}
	return matkuls, nil
}

func InsertMataKuliah(ctx context.Context, matkul models.MataKuliah) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (nama, created_at, updated_at) values('%v', NOW(), NOW())", matkul_table, matkul.Nama)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func UpdateMataKuliah(ctx context.Context, matkul models.MataKuliah, idMataKuliah string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set nama='%v', updated_at=NOW() where id=%v", matkul_table,
		matkul.Nama,
		idMataKuliah,
	)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func DeleteMataKuliah(ctx context.Context, idMataKuliah string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id=%v", matkul_table, idMataKuliah)
	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()

	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
