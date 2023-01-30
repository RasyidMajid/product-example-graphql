package model

type ProductModel struct {
	BaseModels
	NamaProduct  string  `gorm:"column:nama_product"`
	JenisProduct string  `gorm:"column:jenis_product"`
	HargaProduct float64 `gorm:"column:harga_product"`
}

func (ProductModel) TableName() string {
	return "product_table"
}
