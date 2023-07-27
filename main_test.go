package main_test

import (
	main_test "salt-academy_learn_week2"
	"context"
	mocks "salt-academy_learn_week2/mocks/repository"
	"salt-academy_learn_week2/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetListMahasiswaUC(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		mahasiswaMock := &mocks.MahasiswaMock{}
		data := &model.Mahasiswa{
			Name:       "Fajri",
			NIM:        "SA999",
			BirthPlace: "Uzbekistan",
			Handphone:  "6288888888",
			Gender:     "Laki-Laki",
			Address:    "Jalan Duren",
			BirthDate:  time.Now(),
		}
		listData := make([]*model.Mahasiswa, 0)
		listData = append(listData, data)
		listData = append(listData, data)
		listData = append(listData, data)
		mahasiswaMock.On("GetListMahasiswa", mock.Anything).Return(listData, nil)

		res, err := main_test.GetListMahasiswaUC(context.Background(), mahasiswaMock)
		assert.NoError(t, err)
		assert.Equal(t, listData, res)
	})
}
