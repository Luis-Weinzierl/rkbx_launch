package main

import (
	"bufio"
	"context"
	"fmt"
	"image/color"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("Start")

	a := app.New()

	a.Settings().SetTheme(&RkbxTheme{})

	w := a.NewWindow("Rkbx Launch")
	w.SetFixedSize(true)

	consoleLabel := widget.NewLabel("Hello Fyne!")
	scrollBox := container.NewVScroll(consoleLabel)

	scrollBox.SetMinSize(fyne.Size{Width: 700, Height: 300})

	running := false

	ctx, cancel := context.WithCancel(context.Background())
	cmd, c := setupRkbxLinkProcess(ctx, consoleLabel, scrollBox)

	logo := canvas.NewImageFromFile("./assets/LinkLogo.png")
	logo.SetMinSize(fyne.Size{Width: 100, Height: 100})
	logo.FillMode = canvas.ImageFillContain

	runButton := widget.NewButton("Run", func() {
		if !running {
			cmd.Start()

			consoleLabel.SetText("Running...")
			fmt.Println("[rkbx_launch] Running...")
			running = true
		} else {
			cancel()
			<-c

			consoleLabel.SetText("Stopped.")
			scrollBox.ScrollToTop()
			fmt.Println("[rkbx_launch] Stopped.")
			ctx, cancel = context.WithCancel(context.Background())
			cmd, c = setupRkbxLinkProcess(ctx, consoleLabel, scrollBox)
			running = false
		}
	})

	vbox := container.NewVBox(
		container.NewHBox(
			logo,
			scrollBox,
		),
		runButton,
	)

	w.SetContent(container.NewCenter(vbox))

	w.ShowAndRun()

	cancel()
}

func setupRkbxLinkProcess(ctx context.Context, outputLabel *widget.Label, scrollBox *container.Scroll) (*exec.Cmd, chan int) {
	wd, _ := os.Getwd()

	cmd := exec.CommandContext(ctx, wd+"/rkbx_link/rkbx_link.exe")
	cmd.Dir = wd + "/rkbx_link"
	cmd.Stderr = os.Stderr

	c := make(chan int)
	go attachScanner(cmd, outputLabel, scrollBox, c)

	return cmd, c
}

func attachScanner(cmd *exec.Cmd, label *widget.Label, scrollWrap *container.Scroll, c chan int) {
	stdout, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		runes := []rune(scanner.Text())

		if len(runes) > 9 {
			runes = runes[9:]
		}

		if len(runes) > 90 {
			runes = append(runes[:90], []rune("...")...) // Truncate long lines
		}

		line := string(runes)

		fmt.Println(line)

		fyne.Do(func() {
			label.SetText(label.Text + "\r\n" + line)
			scrollWrap.ScrollToBottom()
		})
	}
	c <- 1
}

type RkbxTheme struct{}

func (m *RkbxTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m *RkbxTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m *RkbxTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m *RkbxTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
