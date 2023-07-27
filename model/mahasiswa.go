package model

import "time"

type Mahasiswa struct {
	Name       string    `dbq:"nama"`
	NIM        string    `dbq:"nim"`
	BirthPlace string    `dbq:"tempat_lahir"`
	Handphone  string    `dbq:"no_hp"`
	Gender     string    `dbq:"jenis_kelamin"`
	Address    string    `dbq:"alamat"`
	BirthDate  time.Time `dbq:"tanggal_lahir"`
}

func (Mahasiswa) GetTableName() string {
	return `mahasiswa`
}

func ListRowsInsertMahasiswa() []string {
	return []string{
		"nama",
		"nim",
		"tempat_lahir",
		"no_hp",
		"jenis_kelamin",
		"alamat",
		"tanggal_lahir",
	}
}
