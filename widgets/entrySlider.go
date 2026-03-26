package widgets

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewEntrySlider(title string, min int, max int, binding *int) *fyne.Container {
	slider := widget.NewSlider(float64(min), float64(max))
	slider.Step = 1
	slider.SetValue(float64(*binding))

	entry := widget.NewEntry()
	entry.SetText(fmt.Sprintf("%d", *binding))

	slider.OnChanged = func(v float64) {
		val := int(v)
		entry.SetText(fmt.Sprintf("%d", val))
		*binding = val
	}

	entry.OnChanged = func(v string) {
		if val, err := strconv.Atoi(v); err == nil {
			slider.SetValue(float64(val))
			trueVal := int(slider.Value)
			*binding = trueVal
		}
	}

	entry.OnSubmitted = func(v string) {
		trueVal := int(slider.Value)
		entry.SetText(fmt.Sprintf("%d", trueVal))
	}

	// TODO: Do the onSubmitted stoff when focus is lost

	return container.NewVBox(
		widget.NewLabel(title),
		container.New(layout.NewGridLayoutWithColumns(2), slider, entry),
	)
}

func NewEntrySliderF(title string, min float64, max float64, binding *float64) *fyne.Container {
	slider := widget.NewSlider(min, max)
	slider.Step = 0.01
	slider.SetValue(*binding)

	entry := widget.NewEntry()
	entry.SetText(fmt.Sprintf("%.2f", *binding))

	slider.OnChanged = func(v float64) {
		entry.SetText(fmt.Sprintf("%.2f", v))
		*binding = v
	}

	entry.OnChanged = func(v string) {
		if val, err := strconv.ParseFloat(v, 64); err == nil {
			slider.SetValue(val)
			trueVal := slider.Value
			*binding = trueVal
		}
	}

	entry.OnSubmitted = func(v string) {
		entry.SetText(fmt.Sprintf("%.2f", slider.Value))
	}

	// TODO: Do the onSubmitted stoff when focus is lost

	return container.NewVBox(
		widget.NewLabel(title),
		container.New(layout.NewGridLayoutWithColumns(2), slider, entry),
	)
}
