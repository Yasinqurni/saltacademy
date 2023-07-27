package usecase

import (
	"salt-academy_learn_week2/domain/repository"
	"salt-academy_learn_week2/domain/usecase"
)

type mahasiswaInteractor struct {
	repo repository.MahasiswaTemplate
}

func NewMahasiswa(repo repository.MahasiswaTemplate) usecase.MahasiswaUseCase {
	return &mahasiswaInteractor{repo: repo}
}
