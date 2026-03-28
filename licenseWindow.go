package main

import (
	"com/rkbx_launch/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

func newLicenseWindow(a *fyne.App, registeredCallback func(string), cancelledCallback func()) fyne.Window {
	licenseWindow := (*a).NewWindow("Register rkbx_link")

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
				widget.NewButton("Register", func() {
					if licenseEntry.Validate() == nil {
						registeredCallback(licenseEntry.Text)
					}
				}),
				widget.NewButton("I'm only testing rkbx_link", cancelledCallback),
			),
			nil, nil,
			widgets.NewHero("Register rkbx_link"),
		),
	)

	licenseWindow.CenterOnScreen()
	licenseWindow.FixedSize()
	licenseWindow.Resize(fyne.Size{Width: 500, Height: 300})

	return licenseWindow
}
