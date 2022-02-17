package model

type SizeDetail struct {
	BaseModel
	PanjangBaju		uint	`json:"panjangBaju"`
	LingkarDada		uint	`json:"lingkarDada"`
	PanjangTangan	uint	`json:"panjangTangan"`
	LingkarLeher	uint	`json:"lingkarLeher"`
	LingkarLengan	uint	`json:"lingkarLengan"`
	LingkarPinggang	uint	`json:"lingkarPinggang"`
	PanjangCelana	uint	`json:"panjangCelana"`
	LingkarPinggul	uint	`json:"lingkarPinggul"`
	Pesak			uint	`json:"pesak"`
}