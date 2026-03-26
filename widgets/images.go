package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func NewLogoImage(res fyne.Resource) *canvas.Image {
	logo := canvas.NewImageFromResource(res)
	logo.SetMinSize(fyne.Size{Width: 75, Height: 75})
	logo.FillMode = canvas.ImageFillContain
	return logo
}
