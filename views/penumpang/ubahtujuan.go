package penumpang

import (
	"fmt"
	"project/components"
	"project/models"
	"project/utils"
)

type UbahTujuan struct{}

func (ubahTujuan UbahTujuan) Render() {
	penumpang := models.Penumpang{}
loop:
	for {
		utils.ClearScreen()
		fmt.Println("*** Ubah Stasiun Tujuan ***")
		identitas, _ := components.Input(map[string]interface{}{"label": fmt.Sprintf("%-15s:", "Nomor Identitas")})
		row, err := penumpang.CariIdentitas(fmt.Sprintf("%v", identitas))
		if err != nil {
			menuItem, err := components.Input(map[string]interface{}{"type": "number", "label": fmt.Sprintf("%s\n%s", err, "Tekan Enter untuk memindah penumpang lain, atau masukkan [0] untuk Kembali Ke Menu Utama")})
			fmt.Println(menuItem, err)
			switch {
			case err != nil:
				continue
			case menuItem == 0:
				break loop
			}
			return
		}
		tujuan, err := components.Input(map[string]interface{}{"label": fmt.Sprintf("%-15s:", "Tujuan")})
		if err != nil {
			panic("Perubahan tujuan gagal karena input tidak valid")
		}
		p := &models.DaftarPenumpang[row]
		p.IsiTujuan(fmt.Sprintf("%v", tujuan))

		components.Input(map[string]interface{}{"label": fmt.Sprintf("%s\n%s", "Penumpang berhasil mengubah tujuan", "Tekan Enter untuk kembali ke Menu Utama ...")})
		return
	}
}
