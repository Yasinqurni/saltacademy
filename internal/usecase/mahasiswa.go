package usecase

import (
	"context"
	"salt-academy_learn_week2/domain/entity"
)

func (m *mahasiswaInteractor) GetListMahasiswaUC(ctx context.Context) ([]*entity.Mahasiswa, error) {

	data, err := m.repo.GetListMahasiswa(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
