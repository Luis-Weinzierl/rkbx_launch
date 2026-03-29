package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func NewSelectEntry(title string, bind binding.String, options []string) *fyne.Container {
	sel := widget.NewSelectWithData(options, bind)

	return container.NewVBox(
		widget.NewLabel(title),
		sel,
	)
}
