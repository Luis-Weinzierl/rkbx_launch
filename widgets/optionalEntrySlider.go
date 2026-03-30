package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func NewOptionalEntrySliderF(label string, min float64, max float64, step float64, bindValue binding.Float, bindActive binding.Bool) *fyne.Container {
	entry := NewEntrySliderFWithoutLabel(min, max, step, bindValue)
	check := widget.NewCheckWithData("", bindActive)

	return container.NewVBox(
		widget.NewLabel(label),
		container.NewBorder(
			nil, nil,
			check,
			nil,
			entry,
		),
	)
}
