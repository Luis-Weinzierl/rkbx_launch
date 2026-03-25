package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewBoolConfig(label string, configEntry *bool) *widget.Check {
	w := widget.NewCheck(label, func(b bool) {
		switchCallback(b, configEntry)
	})
	w.SetChecked(*configEntry)
	return w
}

func NewBoolConfigWithSubmenu(label string, configEntry *bool, submenu *fyne.Container) *widget.Check {
	w := widget.NewCheck(label, func(b bool) {
		switchCallback(b, configEntry)

		if b {
			(*submenu).Show()
		} else {
			(*submenu).Hide()
		}
	})
	w.SetChecked(*configEntry)
	return w
}

func switchCallback(b bool, configEntry *bool) {
	*configEntry = b
}
