package main

import (
	"com/rkbx_launch/globalisation"
	"com/rkbx_launch/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

func NewLicenseWindow(a *fyne.App, registeredCallback func(string), cancelledCallback func()) fyne.Window {
	licenseWindow := (*a).NewWindow(globalisation.Get(globalisation.RegisterRkbxLink))

	licenseValidator := validation.NewRegexp("([A-Z0-9]{8}-){3}[A-Z0-9]{8}", "Invalid License Key Format.")

	licenseEntry := widget.NewEntry()
	licenseEntry.PlaceHolder = "xxxxxxxx-xxxxxxxx-xxxxxxxx-xxxxxxxx"
	licenseEntry.AlwaysShowValidationError = true
	licenseEntry.Validator = licenseValidator

	licenseWindow.SetContent(
		container.NewBorder(
			nil,
			container.NewVBox(
				licenseEntry,
				widget.NewButton(globalisation.Get(globalisation.RegisterLabel), func() {
					if licenseEntry.Validate() == nil {
						registeredCallback(licenseEntry.Text)
					}
				}),
				widget.NewButton(globalisation.Get(globalisation.OnlyTestingLabel), cancelledCallback),
			),
			nil, nil,
			widgets.NewHero(globalisation.RegisterRkbxLink),
		),
	)

	licenseWindow.CenterOnScreen()
	licenseWindow.FixedSize()
	licenseWindow.Resize(fyne.Size{Width: 500, Height: 300})

	return licenseWindow
}
