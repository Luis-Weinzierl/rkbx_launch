package main

import (
	"bufio"
	"com/rkbx_launch/helpers"
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
)

const configFilePath = "./rkbx_link/config"
const linkDir = "./rkbx_link/"

func main() {
	fmt.Println("[rkbx_launch] Starting...")

	config := helpers.NewBoundRkbxConfig()
	a := app.NewWithID("rkbx_launch_app")

	a.Settings().SetTheme(&RkbxTheme{})

	mainWindow, cancel := newMainWindow(a, &config)

	var licenseWindow fyne.Window
	licenseWindow = NewLicenseWindow(&a,
		func(key string) {
			config.App_licenseKey.Set(key)
			licenseWindow.Hide()
			helpers.StoreConfigFile(&config, configFilePath)
		},
		func() {
			licenseWindow.Hide()
			mainWindow.Show()
			mainWindow.CenterOnScreen()
		})

	licenseWindow.SetCloseIntercept(func() {
		mainWindow.Show()
		licenseWindow.Hide()
		mainWindow.CenterOnScreen()
	})

	fmt.Println("[rkbx_launch] Checking for updates...")
	version, err := getInstalledVersion()

	if err != nil {
		// rkbx_link is not (properly) installed, download
		var modal fyne.Window
		modal = NewModalWindow(&a, "No Installation Found", "No installation of rkbx_link was found. \nWould you like to download the latest version?", "Download", "Exit",
			func() {
				// Download latest version and open license window as the new copy cannot be registered
				downloadLatestVersion()
				modal.Hide()

				go helpers.LoadConfigFile(configFilePath, &config)
				licenseWindow.Show()
			},
			func() {
				// Quit application as rkbx_launch cannot be used without rkbx_link

				modal.Hide()
				a.Quit()
			},
		)
		modal.Show()
	} else if isUpdateAvailable(version) {
		// update available, download
		var modal fyne.Window
		modal = NewModalWindow(&a, "Update available", "A new version of rkbx_link is available!", "Update", "Continue without updating",
			func() {
				// Download latest version and continue to main / license window
				downloadLatestVersion()
				modal.Hide()
				mainLoop(&config, licenseWindow, mainWindow)
			},
			func() {
				modal.Hide()
				mainLoop(&config, licenseWindow, mainWindow)
			},
		)
		modal.Show()
	} else {
		mainLoop(&config, licenseWindow, mainWindow)
	}

	a.Run()

	cancel()
}

func mainLoop(config *helpers.BoundRkbxConfig, licenseWindow fyne.Window, mainWindow fyne.Window) {
	helpers.LoadConfigFile(configFilePath, config)

	if config.IsEvaluation() {
		licenseWindow.Show()
	} else {
		mainWindow.Show()
		mainWindow.CenterOnScreen()
	}
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
	fmt.Println("[rkbx_launch] Downloading latest version...")
	os.Remove("rkbx_link")
	helpers.HttpDownloadFile("https://raw.githubusercontent.com/grufkork/rkbx_link/9113cbba11822f689af561f8b393016d8ba9093b/version_exe", "version_exe")
	helpers.HttpDownloadFile("https://github.com/grufkork/rkbx_link/releases/latest/download/rkbx_link_win.zip", "latest.temp.zip")
	helpers.Unzip("latest.temp.zip", linkDir)
	os.Remove("latest.temp.zip")
}

func updateLink() {
	version, err := getInstalledVersion()

	if err != nil {

	} else if isUpdateAvailable(version) {

	} else {

	}

}
