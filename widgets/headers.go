package widgets

import (
	"com/rkbx_launch/globalisation"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func NewHero(text string) *canvas.Text {
	header := canvas.NewText(globalisation.Get(text), fyne.CurrentApp().Settings().Theme().Color(theme.ColorNameForeground, 0))
	header.TextSize = 32
	header.Alignment = fyne.TextAlignCenter

	return header
}

func NewHeader(text string) *canvas.Text {
	header := canvas.NewText(globalisation.Get(text), fyne.CurrentApp().Settings().Theme().Color(theme.ColorNameForeground, 0))
	header.TextSize = 32
	header.Alignment = fyne.TextAlignLeading
	header.TextStyle.Bold = true

	return header
}

func NewSubheader(text string) *canvas.Text {
	header := canvas.NewText(globalisation.Get(text), fyne.CurrentApp().Settings().Theme().Color(theme.ColorNameForeground, 0))
	header.TextSize = 24
	header.Alignment = fyne.TextAlignLeading
	header.TextStyle.Bold = true

	return header
}

func NewTitle(text string) *canvas.Text {
	header := canvas.NewText(globalisation.Get(text), fyne.CurrentApp().Settings().Theme().Color(theme.ColorNameForeground, 0))
	header.TextSize = 20
	header.Alignment = fyne.TextAlignLeading
	header.TextStyle.Bold = true

	return header
}

func NewVerticalSpacer() *canvas.Text {
	spacer := canvas.NewText("", color.RGBA{0, 0, 0, 0})
	spacer.TextSize = 24

	return spacer
}
