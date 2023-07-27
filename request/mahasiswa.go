package request

type BodyMahasiswa struct {
	Nama     string `json:"nama"`
	BirthDay string `json:"birth_day"`
	Gender   string `json:"gender"`
}
