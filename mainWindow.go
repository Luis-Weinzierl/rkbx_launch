package main

import (
	"com/rkbx_launch/helpers"
	"com/rkbx_launch/widgets"
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func newMainWindow(a fyne.App, config *helpers.BoundRkbxConfig) (fyne.Window, context.CancelFunc) {
	w := a.NewWindow("rkbx_link")
	w.SetFixedSize(true)

	oscOptions := widgets.NewOscOptions(config)
	ablOptions := widgets.NewAblOptions(config)
	sacnOptions := widgets.NewSacnOptions(config)
	fileOptions := widgets.NewFileOptions(config)
	setlistOptions := widgets.NewSetlistOptions(config)

	availVersions := []string{"7.2.10", "7.2.8", "7.2.6", "7.2.4", "7.2.3", "7.2.2", "7.1.4"}

	if val, err := config.App_licenseKey.Get(); err == nil && val == "evaluation" {
		availVersions = []string{"7.2.2"}
	}

	configuration := container.NewVScroll(
		container.NewBorder(nil, widget.NewLabel("\r\n"), widget.NewLabel("     "), widget.NewLabel("     "), // Hacky padding
			container.NewVBox(
				widgets.NewHeader("Configuration"),
				widgets.NewSubheader("General"),
				widgets.NewSelectEntry("Rekordbox Version", config.Keeper_rekordboxVersion, availVersions),
				widgets.NewBoolConfig("Auto-Update", config.App_autoUpdate),
				widget.NewCheckWithData("Debug Mode", config.App_debug),
				widgets.NewBoolConfig("Keep non-master decks warm", config.Keeper_keepWarm),
				widgets.NewEntrySlider("Update rate (Hz)", 10, 500, config.Keeper_updateRate),
				widgets.NewEntrySlider("Slow Update every n-th", 5, 20, config.Keeper_slowUpdateEveryNth),
				widgets.NewEntrySlider("Delay compensation (ms)", -5, 5, config.Keeper_delayCompensation),
				widgets.NewEntrySlider("Active Decks", 2, 4, config.Keeper_decks),
				widgets.NewSubheader(""), // Hacky Spacer
				widgets.NewSubheader("Modules"),
				widgets.NewTitle("Ableton® Link"),
				widgets.NewBoolConfigWithSubmenu("Enabled", config.Link_enabled, ablOptions),
				ablOptions,
				widgets.NewSubheader(""), // Hacky Spacer
				widgets.NewTitle("Open Sound Control"),
				widgets.NewBoolConfigWithSubmenu("Enabled", config.Osc_enabled, oscOptions),
				oscOptions,
				widgets.NewSubheader(""), // Hacky Spacer
				widgets.NewTitle("sACN"),
				widgets.NewBoolConfigWithSubmenu("Enabled", config.Sacn_enabled, sacnOptions),
				sacnOptions,
				widgets.NewSubheader(""), // Hacky Spacer
				widgets.NewTitle("File Output"),
				widgets.NewBoolConfigWithSubmenu("Enabled", config.File_enabled, fileOptions),
				fileOptions,
				widgets.NewSubheader(""), // Hacky Spacer
				widgets.NewTitle("Setlist Logging"),
				widgets.NewBoolConfigWithSubmenu("Enabled", config.Setlist_enabled, setlistOptions),
				setlistOptions,
			),
		))

	configuration.SetMinSize(fyne.Size{Width: 400, Height: 500})

	runningDisplay := container.NewCenter(
		widget.NewLabel("Stop rkbx_link to configure."),
	)

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

	runButton := widget.NewButton("Start", func() {})
	saveButton := widget.NewButton("Save", func() { helpers.StoreConfigFile(config, "./rkbx_link/config") })

	runButton.OnTapped = func() {
		if !running {
			cmd.Start()

			runButton.SetText("Stop")
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

			runButton.SetText("Start")
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
		container.NewCenter(
			container.NewStack(
				configuration,
				runningDisplay,
			)),
	)

	w.SetCloseIntercept(func() {
		w.Hide()
		a.Quit()
	})

	w.SetContent(vbox)

	return w, cancel
}
