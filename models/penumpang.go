package models

import "errors"

var DaftarPenumpang []Penumpang

type Penumpang struct {
	nama, identitas, tujuan string
	kelas                   int
}

func (penumpang *Penumpang) Nama() string {
	return penumpang.nama
}

func (penumpang *Penumpang) Identitas() string {
	return penumpang.identitas
}

func (penumpang *Penumpang) Tujuan() string {
	return penumpang.tujuan
}

func (penumpang *Penumpang) Kelas() int {
	return penumpang.kelas
}

func (penumpang *Penumpang) IsiNama(nama string) string {
	penumpang.nama = nama
	return penumpang.nama
}

func (penumpang *Penumpang) IsiIdentitas(identitas string) string {
	penumpang.identitas = identitas
	return penumpang.identitas
}

func (penumpang *Penumpang) IsiTujuan(tujuan string) string {
	penumpang.tujuan = tujuan
	return penumpang.tujuan
}

func (penumpang *Penumpang) IsiKelas(kelas int) int {
	penumpang.kelas = kelas
	return penumpang.kelas
}

func (penumpang *Penumpang) Tambah() string {
	DaftarPenumpang = append(DaftarPenumpang, *penumpang)
	return "Tiket Berhasil Terjual"
}

func (penumpang *Penumpang) CariIdentitas(identitas string) (int, error) {
	for i, penumpang := range DaftarPenumpang {
		if identitas == penumpang.identitas {
			return i, nil
		}
	}
	return -1, errors.New("Identitas tidak ditemukan")
}

func (penumpang *Penumpang) Hapus(index int) string {
	DaftarPenumpang = append(DaftarPenumpang[:index], DaftarPenumpang[index+1:]...)
	return "Tiket Berhasil Dibatalkan"
}
