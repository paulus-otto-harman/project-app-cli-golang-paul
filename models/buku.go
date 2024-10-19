package models

var DaftarBuku = []Buku{{"1", "11", "111"}, {"2", "22", "222"}}

type Buku struct {
	judul, penulis, isbn string
}
