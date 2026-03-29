package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func NewSelectEntry(title string, bind binding.String, options binding.StringList) *fyne.Container {
	opt, _ := options.Get()

	sel := widget.NewSelectWithData(opt, bind)

	options.AddListener(binding.NewDataListener(func() {
		v, _ := options.Get()
		b, _ := bind.Get()

		sel.SetOptions(v)
		sel.SetSelected(b)
	}))

	return container.NewVBox(
		widget.NewLabel(title),
		sel,
	)
}
