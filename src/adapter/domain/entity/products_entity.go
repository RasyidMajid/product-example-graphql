package entity

type Product struct {
	ID           *string
	NamaProduct  *string
	JenisProduct *string
	HargaProduct *float64
}
type ProductRequest struct {
	ID           string
	NamaProduct  string
	JenisProduct string
	HargaProduct float64
}
type ProductResponse struct {
	StatusCode string
	StatusDesc string
	Data       *Product
	DataList   []*Product
}
