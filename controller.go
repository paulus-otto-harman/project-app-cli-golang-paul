package main

import "project/views"

func RenderView(view views.View) {
	view.Render()
}
