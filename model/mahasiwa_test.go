package model_test

import (
	"github.com/stretchr/testify/assert"
	"salt-academy_learn_week2/model"
	"testing"
	"time"
)

func TestGetTableMahasiswaName(t *testing.T) {
	model := model.Mahasiswa{
		Name:       "Maleakhi",
		NIM:        "001",
		BirthPlace: "Jakarta",
		Handphone:  "081373100015",
		Gender:     "Men",
		Address:    "JL Kebon Jeruk",
		BirthDate:  time.Now(),
	}
	assert.Equal(t, "m_mahasiswa", model.GetTableName())
}

func TestGetListRowsInsertMahasiswa(t *testing.T) {
	assert.Equal(t, []string{
		"nama",
		"nim",
		"tempat_lahir",
		"no_hp",
		"jenis_kelamin",
		"alamat",
		"tanggal_lahir",
	}, model.ListRowsInsertMahasiswa())
}
