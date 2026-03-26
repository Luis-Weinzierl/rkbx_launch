package widgets

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func NewHeader(text string) *canvas.Text {
	header := canvas.NewText(text, fyne.CurrentApp().Settings().Theme().Color(theme.ColorNameForeground, 0))
	header.TextSize = 32
	header.Alignment = fyne.TextAlignCenter
	header.TextStyle.Bold = true

	return header
}

func NewSubheader(text string) *canvas.Text {
	header := canvas.NewText(text, fyne.CurrentApp().Settings().Theme().Color(theme.ColorNameForeground, 0))
	header.TextSize = 24
	header.Alignment = fyne.TextAlignCenter
	header.TextStyle.Bold = true

	return header
}

func NewTitle(text string) *canvas.Text {
	header := canvas.NewText(fmt.Sprintf(" %s", text), fyne.CurrentApp().Settings().Theme().Color(theme.ColorNameForeground, 0))
	header.TextSize = 20
	header.Alignment = fyne.TextAlignLeading
	header.TextStyle.Bold = true

	return header
}
