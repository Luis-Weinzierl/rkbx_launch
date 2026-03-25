package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func NewLogoImage(uri string) *canvas.Image {
	logo := canvas.NewImageFromFile(uri)
	logo.SetMinSize(fyne.Size{Width: 75, Height: 75})
	logo.FillMode = canvas.ImageFillContain
	return logo
}

func NewStateImage(uri string) *canvas.Image {
	logo := canvas.NewImageFromFile(uri)
	logo.SetMinSize(fyne.Size{Width: 50, Height: 50})
	logo.FillMode = canvas.ImageFillContain
	logo.Hide()
	return logo
}
