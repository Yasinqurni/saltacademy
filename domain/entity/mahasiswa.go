package entity

import "errors"

type Mahasiswa struct {
	name        string
	nim         string
	phoneNumber string
	address     string
	gender      string
}

func NewMahasiswa(MahasiwaName string, MahasiwaNIM string, PhoneNumber string,
	Address string, Gender string) (*Mahasiswa, error) {
	mahasiwa := &Mahasiswa{
		name:        MahasiwaName,
		nim:         MahasiwaNIM,
		phoneNumber: PhoneNumber,
		address:     Address,
		gender:      Gender,
	}

	if mahasiwa.name == "" {
		return nil, errors.New("NAME MUST BE STATED")
	}

	if mahasiwa.nim == "" {
		return nil, errors.New("NIM MUST BE STATED")
	}

	return mahasiwa, nil
}

func (mhs *Mahasiswa) GetName() string {
	return mhs.name
}

func (mhs *Mahasiswa) GetNIM() string {
	return mhs.nim
}
