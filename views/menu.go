package views

import "project/utils"

type Menu struct {
}

func (menu Menu) Render() {
	utils.ClearScreen()
	println("*** PENJUALAN TIKET KERETA API ***")
	println("[1] Penjualan")
	println("[2] Pembatalan")
	println("[3] Ubah Tujuan")
	println("[4] Pindah Kelas")
	println("[5] Daftar Penumpang")
	println("[6] Keluar dari Program")

}
