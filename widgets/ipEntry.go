package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

func NewIpPortEntry(label string, bind binding.String) *fyne.Container {
	entry := widget.NewEntryWithData(bind)
	entry.TextStyle.Monospace = true

	entry.Validator = validation.NewRegexp(`^(\d{1,3}\.){3}\d{1,3}(:\d{1,5})$`, "Invalid IP address format")

	return container.NewVBox(
		widget.NewLabel(label),
		entry,
	)
}

func NewIpEntry(label string, bind binding.String) *fyne.Container {
	entry := widget.NewEntryWithData(bind)
	entry.TextStyle.Monospace = true

	entry.Validator = validation.NewRegexp(`^(\d{1,3}\.){3}\d{1,3}$`, "Invalid IP address format")

	return container.NewVBox(
		widget.NewLabel(label),
		entry,
	)
}
