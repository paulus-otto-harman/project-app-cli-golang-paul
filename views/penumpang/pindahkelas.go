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
	p := penumpang.CariIdentitas(fmt.Sprintf("%v", identitas))
	fmt.Println(p)
	components.Input(map[string]interface{}{"label": "test"})
}
