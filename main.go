package main

import (
	"project/components"
	"project/views"
	"project/views/penumpang"
)

func main() {
	var menuItem interface{}
	var err error
loop:
	for {
		RenderView(views.Menu{})
		menuItem, err = components.Input(map[string]interface{}{"type": "number", "label": "Masukkan Pilihan Anda :"})
		switch {
		case err != nil:
			continue
		case menuItem == 6:
			break loop
		case menuItem == 5:
			RenderView(penumpang.DaftarPenumpang{})
		case menuItem == 4:
			RenderView(penumpang.PindahKelas{})
		case menuItem == 3:
			RenderView(penumpang.UbahTujuan{})
		case menuItem == 2:
			RenderView(penumpang.Pembatalan{})
		case menuItem == 1:
			RenderView(penumpang.Penjualan{})
		}
	}
}
