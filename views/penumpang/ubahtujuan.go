package penumpang

import (
	"fmt"
	"project/components"
	"project/models"
	"project/utils"
)

type UbahTujuan struct{}

func (ubahTujuan UbahTujuan) Render() {
	utils.ClearScreen()
	penumpang := models.Penumpang{}

	fmt.Println("*** Ubah Destinasi ***")
	identitas, _ := components.Input(map[string]interface{}{"label": "Nomor Identitas"})
	penumpang.CariIdentitas(fmt.Sprintf("%v", identitas))
}
