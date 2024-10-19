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
		nama, _ := components.Input(map[string]interface{}{"label": "Nama"})
		identitas, _ := components.Input(map[string]interface{}{"label": "Identitas"})
		tujuan, _ := components.Input(map[string]interface{}{"label": "Tujuan"})
		kelas, _ := components.Input(map[string]interface{}{"label": "Kelas"})

		penumpang := models.Penumpang{}
		penumpang.IsiNama(fmt.Sprintf("%v", nama))
		penumpang.IsiIdentitas(fmt.Sprintf("%v", identitas))
		penumpang.IsiTujuan(fmt.Sprintf("%v", tujuan))
		penumpang.IsiKelas(fmt.Sprintf("%v", kelas))
		penumpang.Tambah()

		menuItem, err := components.Input(map[string]interface{}{"type": "number", "label": "Tekan Enter untuk menambahkan lagi atau Masukkan [0] untuk Kembali Ke Menu Utama"})
		fmt.Println(menuItem, err)
		switch {
		case err != nil:
			continue
		case menuItem == 0:
			break loop
		}
	}

}
