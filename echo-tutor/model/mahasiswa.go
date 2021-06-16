package model

type Mahasiswa struct {
	ID      string `json:"id"`
	Jurusan string `json:"jurusan"`
	Name    string `json:"name"`
	NoTlp   string `json:"no_tlp"`
	Nim     string `json:"nim"`
}
