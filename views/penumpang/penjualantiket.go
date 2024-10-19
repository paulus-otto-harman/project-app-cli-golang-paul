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
		kelas, _ := components.Input(map[string]interface{}{"label": fmt.Sprintf("%s\n%s\n%s\n%-15s:", "Kelas", "[1]ekonomi", "[2]bisnis", "[3]eksekutif")})

		penumpang := models.Penumpang{}
		penumpang.IsiNama(fmt.Sprintf("%v", nama))
		penumpang.IsiIdentitas(fmt.Sprintf("%v", identitas))
		penumpang.IsiTujuan(fmt.Sprintf("%v", tujuan))
		penumpang.IsiKelas(fmt.Sprintf("%v", kelas))
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
