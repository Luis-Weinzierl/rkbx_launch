package main

import (
	"bufio"
	"com/rkbx_launch/helpers"
	"com/rkbx_launch/widgets"
	"context"
	"embed"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var assets embed.FS

//go:embed assets/LinkLogoGlowing.png
var asset_logoGlowing []byte

//go:embed assets/LinkLogoGray.png
var asset_logoGray []byte

//go:embed assets/state_connected.png
var asset_stateConnected []byte

//go:embed assets/state_disconnected.png
var asset_stateDisconnected []byte

func main() {
	fmt.Println("Start")

	config := helpers.ParseConfigFile("./rkbx_link/config")

	fmt.Println(config)

	config.App_debug = true

	a := app.New()

	a.Settings().SetTheme(&RkbxTheme{})

	w := a.NewWindow("Rkbx Launch")
	w.SetFixedSize(true)

	oscOptions := widgets.NewOscOptions(&config)
	ablOptions := widgets.NewAblOptions(&config)
	sacnOptions := widgets.NewSacnOptions(&config)
	fileOptions := widgets.NewFileOptions(&config)
	setlistOptions := widgets.NewSetlistOptions(&config)

	header := canvas.NewText(" Configuration", a.Settings().Theme().Color(theme.ColorNameForeground, 0))
	header.TextSize = 32
	header.TextStyle.Bold = true

	divider := canvas.NewText(" Modules", a.Settings().Theme().Color(theme.ColorNameForeground, 0))
	divider.TextSize = 24
	divider.TextStyle.Bold = true

	configuration := container.NewVScroll(
		container.NewVBox(
			header,
			widgets.NewBoolConfig("Auto-Update", &config.App_autoUpdate),
			widgets.NewBoolConfig("Debug Mode", &config.App_debug),
			widgets.NewEntrySlider("Update rate (Hz)", 10, 500, &config.Keeper_updateRate),
			widgets.NewEntrySlider("Slow Update every n-th", 5, 20, &config.Keeper_slowUpdateEveryNth),
			widgets.NewEntrySlider("Delay compensation (ms)", -5, 5, &config.Keeper_delayCompensation),
			widgets.NewEntrySlider("Active Decks", 2, 4, &config.Keeper_decks),
			widgets.NewBoolConfig("Keep non-master decks warm", &config.Keeper_keepWarm),
			divider,
			widgets.NewBoolConfigWithSubmenu("Ableton Link", &config.Link_enabled, ablOptions),
			ablOptions,
			widgets.NewBoolConfigWithSubmenu("OSC", &config.Osc_enabled, oscOptions),
			oscOptions,
			widgets.NewBoolConfigWithSubmenu("sACN", &config.Sacn_enabled, sacnOptions),
			sacnOptions,
			widgets.NewBoolConfigWithSubmenu("File Output", &config.File_enabled, fileOptions),
			fileOptions,
			widgets.NewBoolConfigWithSubmenu("Setlist Logging", &config.Setlist_enabled, setlistOptions),
			setlistOptions,
		),
	)

	configuration.SetMinSize(fyne.Size{Width: 400, Height: 500})

	runningDisplay := container.NewCenter(
		widget.NewLabel("Stop rkbx_link to configure."),
	)

	runningDisplay.Hide()

	offLogo := widgets.NewLogoImage(asset_logoGray, "LinkLogoGray.png")
	onLogo := widgets.NewLogoImage(asset_logoGlowing, "LinkLogoGlowing.png")

	onLogo.Hide()

	stateConnected := widgets.NewLogoImage(asset_stateConnected, "state_connected.png")
	stateDisconnected := widgets.NewLogoImage(asset_stateDisconnected, "state_disconnected.png")

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
		container.NewStack(
			configuration,
			runningDisplay,
		),
	)

	w.SetContent(vbox)

	w.ShowAndRun()

	cancel()
}

func setupRkbxLinkProcess(ctx context.Context, connectedWidget *canvas.Image, disconnectedWidget *canvas.Image, w *fyne.Window) (*exec.Cmd, chan int) {
	wd, _ := os.Getwd()

	cmd := exec.CommandContext(ctx, wd+"/rkbx_link/rkbx_link.exe")
	cmd.Dir = wd + "/rkbx_link"

	c := make(chan int)
	go attachScanner(cmd, c, connectedWidget, disconnectedWidget, w)

	return cmd, c
}

func attachScanner(cmd *exec.Cmd, c chan int, connectedWidget *canvas.Image, disconnectedWidget *canvas.Image, w *fyne.Window) {
	stdout, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Ensure Rekordbox is running!") ||
			strings.Contains(line, "Connection to Rekordbox lost") {
			fyne.Do(func() {
				connectedWidget.Hide()
				disconnectedWidget.Show()
				(*w).Content().Refresh()
			})
		} else if strings.Contains(line, "Connected to Rekordbox!") {
			fyne.Do(func() {
				connectedWidget.Show()
				disconnectedWidget.Hide()
				(*w).Content().Refresh()
			})
		}

		fmt.Println(line)
	}
	c <- 1
}
