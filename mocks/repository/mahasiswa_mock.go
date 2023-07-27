package mockrepository

import (
	"context"
	"salt-academy_learn_week2/domain/entity"
	"salt-academy_learn_week2/model"

	"github.com/stretchr/testify/mock"
)

type MahasiswaMock struct {
	mock.Mock
}

func (m *MahasiswaMock) GetMahasiswa(ctx context.Context) (*model.Mahasiswa, error) {
	args := m.Called(ctx)
	return args.Get(0).(*model.Mahasiswa), args.Error(1)
}

func (m *MahasiswaMock) GetListMahasiswa(ctx context.Context) ([]*entity.Mahasiswa, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Mahasiswa), args.Error(1)
}

func (m *MahasiswaMock) AddMahasiswa(ctx context.Context, Mahasiswa model.Mahasiswa) error {
	args := m.Called(ctx)
	return args.Error(0)
}
