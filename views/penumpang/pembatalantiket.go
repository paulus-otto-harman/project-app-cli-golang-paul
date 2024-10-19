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

	utils.ClearScreen()
	penumpang := models.Penumpang{}

	fmt.Println("*** Pembatalan Tiket ***")
	identitas, _ := components.Input(map[string]interface{}{"label": "Nomor Identitas"})
	penumpang.Hapus(penumpang.CariIdentitas(fmt.Sprintf("%v", identitas)))

}
