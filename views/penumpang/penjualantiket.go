package penumpang

import (
	"fmt"
	"project/components"
	"project/models"
	"project/utils"
)

type Penjualan struct {
}

func (penjualan Penjualan) Render() {

loop:
	for {
		utils.ClearScreen()
		fmt.Println("*** Tambah penumpang ***")
		nama, _ := components.Input(map[string]interface{}{"label": fmt.Sprintf("%-15s:", "Nama")})
		identitas, _ := components.Input(map[string]interface{}{"label": fmt.Sprintf("%-15s:", "Identitas")})
		tujuan, _ := components.Input(map[string]interface{}{"label": fmt.Sprintf("%-15s:", "Tujuan")})
		gerbong, _ := components.Input(map[string]interface{}{"type": "number", "label": fmt.Sprintf("\n%s\n%s\n%s\n%-15s:", "Kelas (Pilih 1-3)", "[1]Ekonomi", "[2]Bisnis", "[3]Eksekutif")})

		penumpang := models.Penumpang{}
		penumpang.IsiNama(fmt.Sprintf("%v", nama))
		penumpang.IsiIdentitas(fmt.Sprintf("%v", identitas))
		penumpang.IsiTujuan(fmt.Sprintf("%v", tujuan))

		if kelas, ok := gerbong.(int); ok {
			penumpang.IsiKelas(kelas)
		}
		penumpang.Tambah()

		menuItem, err := components.Input(map[string]interface{}{"type": "number", "label": fmt.Sprintf("%s\n%s", "Penumpang berhasil ditambahkan", "Tekan Enter untuk menambahkan lagi atau Masukkan [0] untuk Kembali Ke Menu Utama")})
		fmt.Println(menuItem, err)
		switch {
		case err != nil:
			continue
		case menuItem == 0:
			break loop
		}
	}

}
