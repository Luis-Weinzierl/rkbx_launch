package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewFormEntry(title string, binding *string) *fyne.Container {
	entry := widget.NewEntry()
	entry.SetText(*binding)
	entry.OnChanged = func(s string) {
		*binding = s
	}

	return container.NewBorder(nil, nil, widget.NewLabel(title), nil, entry)
}
