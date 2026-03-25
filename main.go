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

	scrollBox := container.NewVScroll(
		container.NewVBox(
			widgets.NewBoolConfigWithSubmenu("OSC", &config.Osc_enabled, oscOptions),
			oscOptions,
			widgets.NewBoolConfig("sACN", &config.Sacn_enabled),
			widgets.NewBoolConfig("Ableton Link", &config.Link_enabled),
			widgets.NewBoolConfig("File Output", &config.File_enabled),
			widgets.NewBoolConfig("Setlist Logging", &config.Setlist_enabled),
		),
	)

	scrollBox.SetMinSize(fyne.Size{Width: 400, Height: 500})

	offLogo := widgets.NewLogoImage(asset_logoGray, "LinkLogoGray.png")
	onLogo := widgets.NewLogoImage(asset_logoGlowing, "LinkLogoGlowing.png")

	onLogo.Hide()

	stateConnected := widgets.NewStateImage(asset_stateConnected, "state_connected.png")
	stateDisconnected := widgets.NewStateImage(asset_stateDisconnected, "state_disconnected.png")

	running := false

	ctx, cancel := context.WithCancel(context.Background())
	cmd, c := setupRkbxLinkProcess(ctx, stateConnected, stateDisconnected, &w)

	runButton := widget.NewButton("Start", func() {})
	runButton.OnTapped = func() {
		if !running {
			cmd.Start()

			runButton.SetText("Stop")
			fmt.Println("[rkbx_launch] Running...")
			running = true
			onLogo.Show()
			offLogo.Hide()
		} else {
			cancel()
			<-c

			runButton.SetText("Start")
			scrollBox.ScrollToTop()
			fmt.Println("[rkbx_launch] Stopped.")
			ctx, cancel = context.WithCancel(context.Background())
			cmd, c = setupRkbxLinkProcess(ctx, stateConnected, stateDisconnected, &w)
			running = false
			stateConnected.Hide()
			stateDisconnected.Hide()
			onLogo.Hide()
			offLogo.Show()
		}
	}

	saveButton := widget.NewButton("Save", func() { helpers.StoreConfigFile(config, "./rkbx_link/config") })

	logoStack := container.NewStack(offLogo, onLogo)
	stateStack := container.NewStack(stateConnected, stateDisconnected)

	vbox := container.NewVBox(
		container.NewBorder(
			nil, nil,
			logoStack,
			stateStack,
			nil,
		),
		scrollBox,
		saveButton,
		runButton,
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

		if strings.Contains(line, "Ensure Rekordbox is running!") {
			fyne.Do(func() {
				connectedWidget.Hide()
				disconnectedWidget.Show()
				(*w).Content().Refresh()
			})
		} else if strings.Contains(line, "Connected") {
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
