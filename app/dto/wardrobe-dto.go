package dto

import "rentoday.id/app/model"

type AddWardrobe struct {
	OwnerEmail string           `json:"ownerEmail"`
	LabelName  string           `json:"labelName"`
	Category   string           `json:"category"`
	Condition  string           `json:"condition"`
	LabelSize  string           `json:"labelSize"`
	SizeDetail model.SizeDetail `json:"sizeDetail"`
	// SDPanjangBaju		uint		`json:"panjangBaju"`
	// SDLingkarDada		uint		`json:"lingkarDada"`
	// SDPanjangTangan		uint		`json:"panjangTangan"`
	// SDLingkarLeher		uint		`json:"lingkarLeher"`
	// SDLingkarLengan		uint		`json:"lingkarLengan"`
	// SDLingkarPinggang	uint		`json:"lingkarPinggang"`
	// SDPanjangCelana		uint		`json:"panjangCelana"`
	// SDLingkarPinggul	uint		`json:"lingkarPinggul"`
	SDPesak  uint     `json:"pesak"`
	Material string   `json:"material"`
	Defect   []string `json:"defect"`
	Colors   string   `json:"colors"`
	Images   []string `json:"images"`
}
