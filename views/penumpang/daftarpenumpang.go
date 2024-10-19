package penumpang

import (
	"fmt"
	"project/components"
	"project/models"
	"project/utils"
	"strings"
)

type DaftarPenumpang struct{}

func (d DaftarPenumpang) Render() {
	const WidthNama = 50
	const WidthIdentitas = 40
	const WidthTujuan = 30
	const WidthKelas = 30
	Kelas := [3]string{"Ekonomi", "Bisnis", "Eksekutif"}

	utils.ClearScreen()
	// header
	fmt.Printf("%s%s%s%s%s%s%s%s%s\n", "┌─", strings.Repeat("─", WidthNama), "┬─", strings.Repeat("─", WidthIdentitas), "┬─", strings.Repeat("─", WidthTujuan), "┬─", strings.Repeat("─", WidthKelas), "┐")
	fmt.Printf("%s%-*s%s%-*s%s%-*s%s%-*s%s\n", "│ ", WidthNama, "Nama", "│ ", WidthIdentitas, "Identitas", "│ ", WidthTujuan, "Tujuan", "│ ", WidthKelas, "Kelas", "│")
	fmt.Printf("%s%s%s%s%s%s%s%s%s\n", "├─", strings.Repeat("─", WidthNama), "┼─", strings.Repeat("─", WidthIdentitas), "┼─", strings.Repeat("─", WidthTujuan), "┼─", strings.Repeat("─", WidthKelas), "┤")

	// body
	for _, penumpang := range models.DaftarPenumpang {
		if penumpang.Kelas() > 0 {
			fmt.Printf("%s%-*s%s%-*s%s%-*s%s%-*s%s\n", "│ ", WidthNama, penumpang.Nama(), "│ ", WidthIdentitas, penumpang.Identitas(), "│ ", WidthTujuan, penumpang.Tujuan(), "│ ", WidthKelas, Kelas[penumpang.Kelas()-1], "│")
		}

	}

	// footer
	fmt.Printf("%s%s%s%s%s%s%s%s%s\n", "├─", strings.Repeat("─", WidthNama), "┴─", strings.Repeat("─", WidthIdentitas), "┴─", strings.Repeat("─", WidthTujuan), "┴─", strings.Repeat("─", WidthKelas), "┤")
	fmt.Printf("%s%-*s%-*s%-*s%-*s%s\n", "│ ", WidthNama, fmt.Sprintf("%s%d%s", "Total : ", len(models.DaftarPenumpang), " penumpang"), WidthIdentitas+2, "", WidthTujuan+2, "", WidthKelas+2, "", "│")
	fmt.Printf("%s%s%s\n", "╘═", strings.Repeat("═", WidthNama+2+WidthIdentitas+2+WidthTujuan+2+WidthKelas), "╛")

	components.Input(map[string]interface{}{"label": "Tekan Enter untuk kembali ke Menu Utama ..."})
}
