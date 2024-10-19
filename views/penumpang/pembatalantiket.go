package penumpang

import (
	"fmt"
	"project/components"
	"project/models"
	"project/utils"
)

type Pembatalan struct {
}

func (pembatalan Pembatalan) Render() {
	penumpang := models.Penumpang{}
loop:
	for {
		utils.ClearScreen()
		fmt.Println("*** Pembatalan Tiket ***")
		identitas, _ := components.Input(map[string]interface{}{"label": fmt.Sprintf("%-15s:", "Nomor Identitas")})
		row, err := penumpang.CariIdentitas(fmt.Sprintf("%v", identitas))
		if err != nil {
			menuItem, err := components.Input(map[string]interface{}{"type": "number", "label": fmt.Sprintf("%s\n%s", err, "Tekan Enter untuk membatalkan tiket lain, atau masukkan [0] untuk Kembali Ke Menu Utama")})
			fmt.Println(menuItem, err)
			switch {
			case err != nil:
				continue
			case menuItem == 0:
				break loop
			}
			return
		}
		penumpang.Hapus(row)
		components.Input(map[string]interface{}{"label": fmt.Sprintf("%s\n%s", "Tiket berhasil dibatalkan.", "Tekan Enter untuk kembali ke Menu Utama ...")})

	}

}
