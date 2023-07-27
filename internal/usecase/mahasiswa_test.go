package usecase_test

import (
	"context"
	"salt-academy_learn_week2/internal/usecase"
	mocks "salt-academy_learn_week2/mocks/repository"
	"salt-academy_learn_week2/model"
	"salt-academy_learn_week2/model/mapper"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetListMahasiswaUC(t *testing.T) {
	t.Run("positif cases", func(t *testing.T) {

		mahasiswaMock := new(mocks.MahasiswaMock)
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

		result := mapper.MapperListModelToEntityMahasiswa(listData)

		mahasiswaMock.On("GetListMahasiswa", mock.Anything).Return(result, nil)

		uc := usecase.NewMahasiswa(mahasiswaMock)

		res, err := uc.GetListMahasiswaUC(context.Background())

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		assert.Equal(t, result, res)
	})
}
