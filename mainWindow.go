package main

import (
	"com/rkbx_launch/globalisation"
	"com/rkbx_launch/helpers"
	"com/rkbx_launch/widgets"
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	fynetooltip "github.com/dweymouth/fyne-tooltip"
)

func newMainWindow(a fyne.App, config *helpers.RkbxLinkConfig) (fyne.Window, context.CancelFunc) {
	w := a.NewWindow(globalisation.Get(globalisation.MainWindowTitle))

	appOptions := widgets.NewAppOptions(config)
	oscOptions := widgets.NewOscOptions(config)
	ablOptions := widgets.NewAblOptions(config)
	sacnOptions := widgets.NewSacnOptions(config)
	fileOptions := widgets.NewFileOptions(config)
	setlistOptions := widgets.NewSetlistOptions(config)

	configuration := container.NewVScroll(
		container.NewBorder(nil, widget.NewLabel("\r\n"), widget.NewLabel("     "), widget.NewLabel("     "), // Hacky padding
			container.NewVBox(
				widgets.NewHeader(globalisation.ConfigurationTitle),
				widgets.NewSubheader(globalisation.GeneralHeading),
				appOptions,
				widgets.NewVerticalSpacer(), // Hacky Spacer
				widgets.NewSubheader(globalisation.ModulesHeading),
				widgets.NewTitle(globalisation.LinkModuleHeading),
				widgets.NewBoolConfigWithSubmenu(globalisation.EnabledLabel, config.Link_enabled, ablOptions),
				ablOptions,
				widgets.NewVerticalSpacer(), // Hacky Spacer
				widgets.NewTitle(globalisation.OscModuleHeading),
				widgets.NewBoolConfigWithSubmenu(globalisation.EnabledLabel, config.Osc_enabled, oscOptions),
				oscOptions,
				widgets.NewVerticalSpacer(), // Hacky Spacer
				widgets.NewTitle(globalisation.SacnModuleHeading),
				widgets.NewBoolConfigWithSubmenu(globalisation.EnabledLabel, config.Sacn_enabled, sacnOptions),
				sacnOptions,
				widgets.NewVerticalSpacer(), // Hacky Spacer
				widgets.NewTitle(globalisation.FileModuleHeading),
				widgets.NewBoolConfigWithSubmenu(globalisation.EnabledLabel, config.File_enabled, fileOptions),
				fileOptions,
				widgets.NewVerticalSpacer(), // Hacky Spacer
				widgets.NewTitle(globalisation.SetlistModuleHeading),
				widgets.NewBoolConfigWithSubmenu(globalisation.EnabledLabel, config.Setlist_enabled, setlistOptions),
				setlistOptions,
			),
		))

	configuration.SetMinSize(fyne.Size{Width: 400, Height: 500})

	runningDisplay := container.NewHScroll( // Hacky way to get a minsize-able container
		container.NewCenter(
			widget.NewLabel(
				globalisation.Get(globalisation.StopToConfigureLabel)),
		))

	runningDisplay.SetMinSize(fyne.Size{Width: 400, Height: 500})
	runningDisplay.Hide()

	offLogo := widgets.NewLogoImage(resourceLinkLogoGrayPng)
	onLogo := widgets.NewLogoImage(resourceLinkLogoGlowingPng)

	onLogo.Hide()

	stateConnected := widgets.NewLogoImage(resourceIconRekordboxConnectedPng)
	stateDisconnected := widgets.NewLogoImage(resourceIconRekordboxDisconnectedPng)

	stateConnected.Hide()
	stateDisconnected.Hide()

	running := false

	ctx, cancel := context.WithCancel(context.Background())
	cmd, c := setupRkbxLinkProcess(ctx, stateConnected, stateDisconnected, &w)

	runButton := widget.NewButton(globalisation.Get(globalisation.StartLabel), func() {})
	saveButton := widget.NewButton(globalisation.Get(globalisation.SaveLabel), func() {
		fmt.Println(config.Sacn_targets.Get())
		go helpers.StoreConfigFile(config, "./rkbx_link/config")
	})

	runButton.OnTapped = func() {
		if !running {
			helpers.StoreConfigFile(config, "./rkbx_link/config")
			cmd.Start()

			runButton.SetText(globalisation.Get(globalisation.StopLabel))
			fmt.Println("[rkbx_launch] Running...")
			running = true
			saveButton.Hide()
			configuration.Hide()
			runningDisplay.Show()
			onLogo.Show()
			offLogo.Hide()
		} else {
			cancel()
			<-c

			runButton.SetText(globalisation.Get(globalisation.StartLabel))
			fmt.Println("[rkbx_launch] Stopped.")
			ctx, cancel = context.WithCancel(context.Background())
			cmd, c = setupRkbxLinkProcess(ctx, stateConnected, stateDisconnected, &w)
			running = false
			saveButton.Show()
			configuration.Show()
			runningDisplay.Hide()
			stateConnected.Hide()
			stateDisconnected.Hide()
			onLogo.Hide()
			offLogo.Show()
		}
	}

	logoStack := container.NewStack(offLogo, onLogo)
	stateStack := container.NewStack(stateConnected, stateDisconnected)

	vbox := container.NewBorder(
		container.NewBorder(
			nil, nil,
			logoStack,
			stateStack,
			nil,
		),
		container.NewVBox(
			saveButton,
			runButton,
		),
		nil,
		nil,
		container.NewStack(
			configuration,
			runningDisplay,
		),
	)

	config.HasUnsavedChanges.AddListener(binding.NewDataListener(func() {
		hasUnsavedChanges, _ := config.HasUnsavedChanges.Get()

		if hasUnsavedChanges {
			runButton.Hide()
			w.SetTitle(globalisation.Get(globalisation.MainWindowTitleUnsaved))
		} else {
			runButton.Show()
			w.SetTitle(globalisation.Get(globalisation.MainWindowTitle))
		}
	}))

	w.SetCloseIntercept(func() {
		w.Close()
		fynetooltip.DestroyWindowToolTipLayer(w.Canvas())
	})

	w.SetContent(fynetooltip.AddWindowToolTipLayer(vbox, w.Canvas()))
	w.SetMaster()
	return w, cancel
}
