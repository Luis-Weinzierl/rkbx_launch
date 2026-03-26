package widgets

import (
	"bytes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func NewLogoImage(data []byte, name string) *canvas.Image {
	logo := canvas.NewImageFromReader(bytes.NewReader(data), name)
	logo.SetMinSize(fyne.Size{Width: 75, Height: 75})
	logo.FillMode = canvas.ImageFillContain
	return logo
}
