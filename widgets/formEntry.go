package widgets

import (
	"com/rkbx_launch/globalisation"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func NewFormEntry(title string, bind binding.String) *fyne.Container {
	entry := widget.NewEntryWithData(bind)

	return container.NewBorder(nil, nil,
		widget.NewLabel(globalisation.Get(title)), nil,
		entry)
}
