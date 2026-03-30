package widgets

import (
	"com/rkbx_launch/globalisation"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func NewBoolConfig(label string, binding binding.Bool) *fyne.Container {
	w := widget.NewCheckWithData(globalisation.Get(label), binding)
	return container.NewHBox(w, NewInfoHover(label))
}

func NewBoolConfigWithSubmenu(text string, bind binding.Bool, submenu *fyne.Container) *widget.Check {
	w := widget.NewCheckWithData(globalisation.Get(text), bind)

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
