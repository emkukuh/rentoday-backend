package model

type SizeDetail struct {
	BaseModel
	PanjangBaju     uint `json:"panjangBaju,omitempty"`
	LingkarDada     uint `json:"lingkarDada,omitempty"`
	PanjangTangan   uint `json:"panjangTangan,omitempty"`
	LingkarLeher    uint `json:"lingkarLeher,omitempty"`
	LingkarLengan   uint `json:"lingkarLengan,omitempty"`
	LingkarPinggang uint `json:"lingkarPinggang,omitempty"`
	PanjangCelana   uint `json:"panjangCelana,omitempty"`
	LingkarPinggul  uint `json:"lingkarPinggul,omitempty"`
	Pesak           uint `json:"pesak,omitempty"`
}
