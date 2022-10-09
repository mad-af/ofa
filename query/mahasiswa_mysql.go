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
	mahasiswa_table = "mahasiswa"
)

func GetAllMahasiswa(ctx context.Context) ([]models.Mahasiswa, error) {
	var students []models.Mahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := "SELECT * FROM " + mahasiswa_table + " Order By created_at DESC"
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var student models.Mahasiswa
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&student.ID,
			&student.Nama,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		student.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		student.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		students = append(students, student)
	}
	return students, nil
}

func InsertMahasiswa(ctx context.Context, student models.Mahasiswa) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (nama, created_at, updated_at) values('%v', NOW(), NOW())", mahasiswa_table, student.Nama)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func UpdateMahasiswa(ctx context.Context, student models.Mahasiswa, idNilai string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set nama='%v', updated_at=NOW() where id=%v", mahasiswa_table,
		student.Nama,
		idNilai,
	)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func DeleteMahasiswa(ctx context.Context, idMahasiswa string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cannot Connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id=%v", mahasiswa_table, idMahasiswa)
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
