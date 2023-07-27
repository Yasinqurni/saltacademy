package mapper

import (
	"encoding/json"
	"salt-academy_learn_week2/domain/entity"
	"salt-academy_learn_week2/model"
)

type MahasiswaJSON struct {
	Name      string `json:"nama"`
	NIM       string `json:"nim"`
	BirthDate string `json:"birthDate"`
}

func MapperModelToJson(model *model.Mahasiswa) []byte {
	dataMahasiswa := &MahasiswaJSON{
		Name:      model.Name,
		NIM:       model.NIM,
		BirthDate: "DATE: " + model.BirthDate.String(),
	}

	jsonMahasiwa, _ := json.Marshal(dataMahasiswa)
	return jsonMahasiwa
}

func MapperCollectionModelToCollectionJSON(models []*model.Mahasiswa) []byte {
	var datas []*MahasiswaJSON
	for _, value := range models {
		dataMahasiswa := &MahasiswaJSON{
			Name:      value.Name,
			NIM:       value.NIM,
			BirthDate: "DATE: " + value.BirthDate.String(),
		}

		datas = append(datas, dataMahasiswa)
	}

	jsonMahasiswa, _ := json.Marshal(datas)
	return jsonMahasiswa
}

func MapperListModelToEntityMahasiswa(models []*model.Mahasiswa) []*entity.Mahasiswa {
	var entities []*entity.Mahasiswa
	for _, model := range models {
		entity, err := entity.NewMahasiswa(model.Name, model.NIM, model.Handphone, model.Address, model.Gender)
		if err != nil {
			return nil
		}
		entities = append(entities, entity)
	}

	return entities
}
