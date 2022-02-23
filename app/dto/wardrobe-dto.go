package dto

type wardrobeSizeDetail struct {
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

type AddWardrobeRequest struct {
	UserID uint
	// OwnerEmail string             `json:"ownerEmail"`
	LabelName string `json:"labelName"`
	Category  string `json:"category"`
	Condition string `json:"condition"`
	LabelSize string `json:"labelSize"`
	// SizeDetail wardrobeSizeDetail `json:"sizeDetail"`
	Material string `json:"material"`
	// Defects  []string `json:"defects,omitempty"`
	// Colors     []string           `json:"colors"`
	// Images     []string           `json:"images"`
}

type AddWardrobeResponse struct {
	// OwnerEmail string `json:"ownerEmail"`
	LabelName  string             `json:"labelName"`
	Category   string             `json:"category"`
	Condition  string             `json:"condition"`
	LabelSize  string             `json:"labelSize"`
	SizeDetail wardrobeSizeDetail `json:"sizeDetail"`
	Material   string             `json:"material"`
	// Defects    []string           `json:"defects,omitempty"`
	// Colors     []string           `json:"colors"`
	// Images     []string           `json:"images"`
}

type GetWardrobesResponse struct {
	Wardrobes []AddWardrobeResponse `json:"wardrobes"`
}
