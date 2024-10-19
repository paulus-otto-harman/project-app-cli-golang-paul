package penumpang

import (
	"fmt"
	"project/components"
	"project/models"
	"project/utils"
)

type PindahKelas struct{}

func (pindahKelas PindahKelas) Render() {
	utils.ClearScreen()
	penumpang := models.Penumpang{}

	fmt.Println("*** Pindah Gerbong (Ekonomi/Eksekutif/Bisnis) ***")
	identitas, _ := components.Input(map[string]interface{}{"label": "Nomor Identitas"})
	row, err := penumpang.CariIdentitas(fmt.Sprintf("%v", identitas))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(row)
}
