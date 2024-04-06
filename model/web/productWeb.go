package web


type ProductWeb struct {
	ProdukName string `json:"name_produk"`
	Deskripsi  string `json:"deskripsi"`
	Category   string `json:"category"`
	Userid     int `json:"user_id"`
	Harga int `json:"harga"`
	Quantity int `json:"quantity"`
}