package main

import (
	"com/rkbx_launch/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewModalWindow(a *fyne.App, title string, hero string, accept string, deny string, acceptedCallback func(), cancelledCallback func()) fyne.Window {
	modalWindow := (*a).NewWindow(title)

	modalWindow.SetContent(
		container.NewBorder(
			nil,
			container.NewVBox(
				widget.NewButton(accept, acceptedCallback),
				widget.NewButton(deny, cancelledCallback),
			),
			nil, nil,
			widgets.NewHero(hero),
		),
	)

	modalWindow.CenterOnScreen()
	modalWindow.FixedSize()
	modalWindow.Resize(fyne.Size{Width: 500, Height: 300})

	return modalWindow
}
