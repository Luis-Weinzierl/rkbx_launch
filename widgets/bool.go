package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func NewBoolConfig(label string, binding binding.Bool) *widget.Check {
	w := widget.NewCheckWithData(label, binding)
	return w
}

func NewBoolConfigWithSubmenu(label string, bind binding.Bool, submenu *fyne.Container) *widget.Check {
	w := widget.NewCheckWithData(label, bind)

	val, _ := bind.Get()

	if val {
		submenu.Show()
	} else {
		submenu.Hide()
	}

	w.OnChanged = func(b bool) {
		if b {
			submenu.Show()
		} else {
			submenu.Hide()
		}
	}

	return w
}
