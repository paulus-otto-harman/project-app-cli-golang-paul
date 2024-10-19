package views

import (
	"fmt"
	"project/coms"
	"project/models"
	"project/utils"
)

func HapusBuku() {
	utils.ClearScreen()
	buku := models.Buku{}

	fmt.Println("*** Hapus Buku ***")
	isbn, _ := coms.Input(map[string]interface{}{"label": "Isbn"})
	buku.Hapus(buku.CariIsbn(fmt.Sprintf("%v", isbn)))
}
