package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewSelectEntry(title string, binding *string, options []string) *fyne.Container {
	sel := widget.NewSelect(options, func(s string) {
		*binding = s
	})

	sel.SetSelected(*binding)

	return container.NewVBox(
		widget.NewLabel(title),
		sel,
	)
}
