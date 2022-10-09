package models

import "time"

type NilaiMahasiswa struct {
	ID           uint      `json:"id"`
	Indeks       string    `json:"indeks"`
	Skor         uint      `json:"skor"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	MahasiswaId  uint      `json:"mahasiswa_id"`
	MataKuliahId uint      `json:"mata_kuliah_id"`
}
