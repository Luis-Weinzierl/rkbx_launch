package main

import (
	"com/rkbx_launch/globalisation"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewModalWindow(a *fyne.App, title string, hero string, accept string, deny string, acceptedCallback func(), cancelledCallback func()) fyne.Window {
	modalWindow := (*a).NewWindow(globalisation.Get(title))

	hStyle := widget.RichTextStyleHeading
	hStyle.Alignment = fyne.TextAlignCenter

	pStyle := widget.RichTextStyleParagraph
	pStyle.Alignment = fyne.TextAlignCenter

	rt := widget.NewRichText(
		&widget.TextSegment{Text: globalisation.Get(title), Style: hStyle},
		&widget.TextSegment{Text: globalisation.Get(hero), Style: pStyle})

	modalWindow.SetContent(
		container.NewBorder(
			nil,
			container.NewVBox(
				widget.NewButton(globalisation.Get(accept), acceptedCallback),
				widget.NewButton(globalisation.Get(deny), cancelledCallback),
			),
			nil, nil,
			container.NewCenter(
				rt,
			),
		),
	)

	modalWindow.CenterOnScreen()
	modalWindow.FixedSize()
	modalWindow.Resize(fyne.Size{Width: 500, Height: 300})

	return modalWindow
}
