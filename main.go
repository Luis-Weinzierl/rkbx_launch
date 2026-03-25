package main

import (
	"bufio"
	"context"
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

func main() {
	fmt.Println("Start")

	config := parseConfigFile("./rkbx_link/config")

	fmt.Println(config)

	config.app_debug = true

	a := app.New()

	a.Settings().SetTheme(&RkbxTheme{})

	w := a.NewWindow("Rkbx Launch")
	w.SetFixedSize(true)

	oscOptions := container.NewVBox(
		widget.NewLabel("Target Address"),
		container.NewBorder(nil, nil, nil, widget.NewButton("Set", func() {}), widget.NewEntry()),
	)
	oscOptions.Hide()

	scrollBox := container.NewVScroll(
		container.NewVBox(
			createWidgetFromBoolWithSubmenu("OSC", &config.osc_enabled, oscOptions),
			oscOptions,
			createWidgetFromBool("sACN", &config.sacn_enabled),
			createWidgetFromBool("Ableton Link", &config.link_enabled),
			createWidgetFromBool("File Output", &config.file_enabled),
			createWidgetFromBool("Setlist Logging", &config.setlist_enabled),
		),
	)

	scrollBox.SetMinSize(fyne.Size{Width: 300, Height: 500})

	offLogo := createLogoImageFromURI("assets/LinkLogoGray.png")
	onLogo := createLogoImageFromURI("assets/LinkLogoGlowing.png")

	onLogo.Hide()

	stateConnected := createStateImageFromURI("assets/state_connected.png")
	stateDisconnected := createStateImageFromURI("assets/state_disconnected.png")

	running := false

	ctx, cancel := context.WithCancel(context.Background())
	cmd, c := setupRkbxLinkProcess(ctx, stateConnected, stateDisconnected)

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
			cmd, c = setupRkbxLinkProcess(ctx, stateConnected, stateDisconnected)
			running = false
			stateConnected.Hide()
			stateDisconnected.Hide()
			onLogo.Hide()
			offLogo.Show()
		}
	}

	saveButton := widget.NewButton("Save", func() { storeConfigFile(config, "./rkbx_link/config") })

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

	w.SetContent(container.NewCenter(vbox))

	w.ShowAndRun()

	cancel()
}

func setupRkbxLinkProcess(ctx context.Context, connectedWidget *canvas.Image, disconnectedWidget *canvas.Image) (*exec.Cmd, chan int) {
	wd, _ := os.Getwd()

	cmd := exec.CommandContext(ctx, wd+"/rkbx_link/rkbx_link.exe")
	cmd.Dir = wd + "/rkbx_link"
	cmd.Stderr = os.Stderr

	c := make(chan int)
	go attachScanner(cmd, c, connectedWidget, disconnectedWidget)

	return cmd, c
}

func attachScanner(cmd *exec.Cmd, c chan int, connectedWidget *canvas.Image, disconnectedWidget *canvas.Image) {
	stdout, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Ensure Rekordbox is running!") {
			connectedWidget.Hide()
			disconnectedWidget.Show()
		} else if strings.Contains(line, "Connected") {
			connectedWidget.Show()
			disconnectedWidget.Hide()
		}

		fmt.Println(line)
	}
	c <- 1
}

func createLogoImageFromURI(uri string) *canvas.Image {
	logo := canvas.NewImageFromFile(uri)
	logo.SetMinSize(fyne.Size{Width: 75, Height: 75})
	logo.FillMode = canvas.ImageFillContain
	return logo
}

func createStateImageFromURI(uri string) *canvas.Image {
	logo := canvas.NewImageFromFile(uri)
	logo.SetMinSize(fyne.Size{Width: 50, Height: 50})
	logo.FillMode = canvas.ImageFillContain
	logo.Hide()
	return logo
}

func createWidgetFromBool(label string, configEntry *bool) *widget.Check {
	w := widget.NewCheck(label, func(b bool) {
		switchCallback(b, configEntry)
	})
	w.SetChecked(*configEntry)
	return w
}

func createWidgetFromBoolWithSubmenu(label string, configEntry *bool, submenu *fyne.Container) *widget.Check {
	w := widget.NewCheck(label, func(b bool) {
		switchCallback(b, configEntry)

		if b {
			(*submenu).Show()
		} else {
			(*submenu).Hide()
		}
	})
	w.SetChecked(*configEntry)
	return w
}

func switchCallback(b bool, configEntry *bool) {
	*configEntry = b
}
