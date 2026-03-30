package main

import (
	"com/rkbx_launch/globalisation"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const titleSuffix = "/title"
const contentSuffix = "/content"
const acceptSuffix = "/accept"
const denySuffix = "/deny"

func NewModalWindow(a *fyne.App, modalId string, acceptedCallback func(), cancelledCallback func()) fyne.Window {
	modalWindow := (*a).NewWindow(globalisation.Get(modalId + titleSuffix))

	hStyle := widget.RichTextStyleHeading
	hStyle.Alignment = fyne.TextAlignCenter

	pStyle := widget.RichTextStyleParagraph
	pStyle.Alignment = fyne.TextAlignCenter

	rt := widget.NewRichText(
		&widget.TextSegment{Text: globalisation.Get(modalId + titleSuffix), Style: hStyle},
		&widget.TextSegment{Text: globalisation.Get(modalId + contentSuffix), Style: pStyle})

	modalWindow.SetContent(
		container.NewBorder(
			nil,
			container.NewVBox(
				widget.NewButton(globalisation.Get(modalId+acceptSuffix), acceptedCallback),
				widget.NewButton(globalisation.Get(modalId+denySuffix), cancelledCallback),
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
