package main

import (
	"bufio"
	"com/rkbx_launch/helpers"
	"com/rkbx_launch/interfaces"
	"com/rkbx_launch/widgets"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	a.Settings().SetTheme(&RkbxTheme{})

	w := a.NewWindow("rkbx_link")
	w.SetFixedSize(true)

	var config helpers.RkbxConfig
	var licenseWindow fyne.Window
	licenseWindow = interfaces.NewLicenseWindow(&a,
		func(key string) {
			config.App_licenseKey = key
			licenseWindow.Hide()
			helpers.StoreConfigFile(config, "./rkbx_link/config")
		},
		func() {
			licenseWindow.Hide()
			w.Show()
		})

	version, err := getInstalledVersion()

	if err != nil {
		dia := dialog.NewConfirm("Install rkbx_link?", "rkbx_launch couldn't find rkbx_link. Install now?", func(b bool) {
			if b {
				downloadLatestVersion()
				version, err = getInstalledVersion()
				config = helpers.ParseConfigFile("./rkbx_link/config")
				fmt.Println(config)
			} else {
				panic("cancelled")
			}
		}, licenseWindow)

		dia.Show()
		licenseWindow.Show()
	} else if isUpdateAvailable(version) {
		dia := dialog.NewConfirm("Update rkbx_link?", "An update for rkbx_link is available. Install now?", func(b bool) {
			if b {
				downloadLatestVersion()
				version, err = getInstalledVersion()
			}
			config = helpers.ParseConfigFile("./rkbx_link/config")
		}, licenseWindow)

		dia.Show()
		licenseWindow.Show()
	} else {
		config = helpers.ParseConfigFile("./rkbx_link/config")
		if config.App_licenseKey == "evaluation" {
			licenseWindow.Show()
		} else {
			w.Show()
		}
	}

	oscOptions := widgets.NewOscOptions(&config)
	ablOptions := widgets.NewAblOptions(&config)
	sacnOptions := widgets.NewSacnOptions(&config)
	fileOptions := widgets.NewFileOptions(&config)
	setlistOptions := widgets.NewSetlistOptions(&config)

	availVersions := []string{"7.2.10", "7.2.8", "7.2.6", "7.2.4", "7.2.3", "7.2.2", "7.1.4"}

	if config.App_licenseKey == "evaluation" {
		availVersions = []string{"7.2.2"}
	}

	configuration := container.NewVScroll(
		container.NewVBox(
			widgets.NewHeader("Configuration"),
			widgets.NewSubheader("General"),
			widgets.NewSelectEntry("Rekordbox Version", &config.Keeper_rekordboxVersion, availVersions),
			widgets.NewBoolConfig("Auto-Update", &config.App_autoUpdate),
			widgets.NewBoolConfig("Debug Mode", &config.App_debug),
			widgets.NewBoolConfig("Keep non-master decks warm", &config.Keeper_keepWarm),
			widgets.NewEntrySlider("Update rate (Hz)", 10, 500, &config.Keeper_updateRate),
			widgets.NewEntrySlider("Slow Update every n-th", 5, 20, &config.Keeper_slowUpdateEveryNth),
			widgets.NewEntrySlider("Delay compensation (ms)", -5, 5, &config.Keeper_delayCompensation),
			widgets.NewEntrySlider("Active Decks", 2, 4, &config.Keeper_decks),
			widgets.NewSubheader(""), // Hacky Spacer
			widgets.NewSubheader("Modules"),
			widgets.NewTitle("Ableton® Link"),
			widgets.NewBoolConfigWithSubmenu("Enabled", &config.Link_enabled, ablOptions),
			ablOptions,
			widgets.NewSubheader(""), // Hacky Spacer
			widgets.NewTitle("Open Sound Control"),
			widgets.NewBoolConfigWithSubmenu("Enabled", &config.Osc_enabled, oscOptions),
			oscOptions,
			widgets.NewSubheader(""), // Hacky Spacer
			widgets.NewTitle("sACN"),
			widgets.NewBoolConfigWithSubmenu("Enabled", &config.Sacn_enabled, sacnOptions),
			sacnOptions,
			widgets.NewSubheader(""), // Hacky Spacer
			widgets.NewTitle("File Output"),
			widgets.NewBoolConfigWithSubmenu("Enabled", &config.File_enabled, fileOptions),
			fileOptions,
			widgets.NewSubheader(""), // Hacky Spacer
			widgets.NewTitle("Setlist Logging"),
			widgets.NewBoolConfigWithSubmenu("Enabled", &config.Setlist_enabled, setlistOptions),
			setlistOptions,
		),
	)

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

	w.SetContent(vbox)

	w.CenterOnScreen()

	windowSize := w.Canvas().Size()
	windowSize.Width += 32
	w.Resize(windowSize)

	a.Run()

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

func getInstalledVersion() (string, error) {
	stream, err := os.Open("version_exe")

	if err != nil {
		return "", err
	}

	content, err := io.ReadAll(stream)

	return string(content), err
}

func isUpdateAvailable(installedVersion string) bool {
	resp, err := http.Get("https://raw.githubusercontent.com/grufkork/rkbx_link/9113cbba11822f689af561f8b393016d8ba9093b/version_exe")

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return false
	}

	latest := string(body)

	fmt.Println(latest)

	return installedVersion != latest
}

func downloadLatestVersion() {
	os.Remove("rkbx_link")
	helpers.HttpDownloadFile("https://raw.githubusercontent.com/grufkork/rkbx_link/9113cbba11822f689af561f8b393016d8ba9093b/version_exe", "version_exe")
	helpers.HttpDownloadFile("https://github.com/grufkork/rkbx_link/releases/latest/download/rkbx_link_win.zip", "latest.temp.zip")
	helpers.Unzip("latest.temp.zip", "./rkbx_link/")
	os.Remove("latest.temp.zip")
}
