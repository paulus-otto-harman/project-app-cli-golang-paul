package models

var DaftarBuku = []Buku{{"11", "11", "111"}, {"2222", "22", "222"}}

type Buku struct {
	judul, penulis, isbn string
}

func (buku Buku) Judul() string {
	return buku.judul
}

func (buku Buku) Penulis() string {
	return buku.penulis
}

func (buku Buku) Isbn() string {
	return buku.isbn
}
