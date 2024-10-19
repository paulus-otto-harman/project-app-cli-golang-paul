package controllers

import (
	"fmt"
	"project/coms"
	"project/models"
)

func TambahBuku() {

}

func HapusBuku() {

}

func IndexBuku() interface{} {
	fmt.Println(models.DaftarBuku)
	coms.Input(map[string]interface{}{"label": "Tekan Enter untuk kembali ke Menu Utama"})
	return nil
}
