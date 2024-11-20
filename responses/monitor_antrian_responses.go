package responses

type MonitorAntrianSaatIniResponse struct {
	Loket           string `json:"loket"`
	KategoriAntrian string `json:"kategori_antrian"`
	Seqno           int    `json:"seqno"`
	NomorAntrian    string `json:"nomor_antrian"`
	Aktif           int    `json:"aktif"`
	Panggil         int    `json:"panggil"`
}

type MonitorAntrianResponse struct {
	JumlahKunjungan int                             `json:"jumlah_kunjungan"`
	KategoriAntrian string                          `json:"kategori_antrian"`
	Seqno           int                             `json:"seqno"`
	NomorPanggil    string                          `json:"nomor_panggil"`
	Loket           string                          `json:"loket"`
	TextPanggilan   string                          `json:"text_panggilan"`
	DataAntrian     []MonitorAntrianSaatIniResponse `json:"antrian"`
}
