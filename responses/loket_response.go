package responses

type LoketResponse struct {
	Id        string `json:"id"`
	Kode      string `json:"kode"`
	Deskripsi string `json:"deskripsi"`
	Aktif     int    `json:"aktif"`
}
