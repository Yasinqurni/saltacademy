package usecase

import (
	"context"
	"salt-academy_learn_week2/domain/entity"
)

type MahasiswaUseCase interface {
	GetListMahasiswaUC(ctx context.Context) ([]*entity.Mahasiswa, error)
}
