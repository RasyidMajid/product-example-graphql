package model

type Product struct {
	ID           *string  `json:"id"`
	NamaProduct  *string  `json:"NamaProduct"`
	JenisProduct *string  `json:"JenisProduct"`
	HargaProduct *float64 `json:"HargaProduct"`
}

type ProductRequest struct {
	ID           string  `json:"id"`
	NamaProduct  string  `json:"NamaProduct"`
	JenisProduct string  `json:"JenisProduct"`
	HargaProduct float64 `json:"HargaProduct"`
}

type ProductResponse struct {
	StatusCode string     `json:"StatusCode"`
	StatusDesc string     `json:"StatusDesc"`
	Data       *Product   `json:"Data"`
	DataList   []*Product `json:"DataList"`
}
