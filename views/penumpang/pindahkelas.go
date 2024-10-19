package penumpang

import (
	"fmt"
	"project/components"
	"project/models"
	"project/utils"
)

type PindahKelas struct{}

func (pindahKelas PindahKelas) Render() {
	penumpang := models.Penumpang{}
loop:
	for {
		utils.ClearScreen()
		fmt.Println("*** Pindah Gerbong (Ekonomi/Eksekutif/Bisnis) ***")
		identitas, _ := components.Input(map[string]interface{}{"label": fmt.Sprintf("%-15s:", "Nomor Identitas")})
		row, err := penumpang.CariIdentitas(fmt.Sprintf("%v", identitas))
		if err != nil {
			menuItem, err := components.Input(map[string]interface{}{"type": "number", "label": fmt.Sprintf("%s\n%s", err, "Tekan Enter untuk pindah penumpang lain, atau masukkan [0] untuk Kembali Ke Menu Utama")})
			fmt.Println(menuItem, err)
			switch {
			case err != nil:
				continue
			case menuItem == 0:
				break loop
			}
			return
		}
		gerbong, err := components.Input(map[string]interface{}{"type": "number", "label": fmt.Sprintf("\n%s\n%s\n%s\n%-15s:", "Kelas (Pilih 1-3)", "[1]Ekonomi", "[2]Bisnis", "[3]Eksekutif")})
		if err != nil {
			panic("Perpindahan Kelas/Gerbong Penumpang tidak valid")
		}
		if kelas, ok := gerbong.(int); ok {
			p := &models.DaftarPenumpang[row]
			p.IsiKelas(kelas)
		}

		components.Input(map[string]interface{}{"label": fmt.Sprintf("%s\n%s", "Penumpang berhasil pindah kelas", "Tekan Enter untuk kembali ke Menu Utama ...")})
		return
	}
}
