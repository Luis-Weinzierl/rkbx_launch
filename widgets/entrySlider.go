package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewEntrySlider(title string, min int, max int, bind binding.Int) *fyne.Container {
	bindF := binding.IntToFloat(bind)
	entryBind := binding.IntToString(bind)

	slider := widget.NewSliderWithData(float64(min), float64(max), bindF)
	slider.Step = 1

	entry := widget.NewEntryWithData(entryBind)

	return container.NewVBox(
		NewInfoLabel(title),
		container.New(layout.NewGridLayoutWithColumns(2), slider, entry),
	)
}

func NewEntrySliderF(title string, min float64, max float64, bind binding.Float) *fyne.Container {
	return container.NewVBox(
		NewInfoLabel(title),
		NewEntrySliderFWithoutLabel(min, max, 0.01, bind),
	)
}

func NewEntrySliderFWithoutLabel(min float64, max float64, step float64, bind binding.Float) *fyne.Container {
	entryBind := binding.FloatToStringWithFormat(bind, "%.2f")

	slider := widget.NewSliderWithData(min, max, bind)
	slider.Step = step

	entry := widget.NewEntryWithData(entryBind)

	return container.New(layout.NewGridLayoutWithColumns(2), slider, entry)
}
