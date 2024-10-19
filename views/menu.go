package views

import "project/utils"

func HomeMenu() {
	utils.ClearScreen()
	println("[1] Tambah Buku")
	println("[2] Hapus Buku")
	println("[3] Tampilkan Buku")
	println("[4] Keluar dari Program")
}
