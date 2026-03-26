package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type RkbxTheme struct{}

func (m *RkbxTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.Black
	case theme.ColorNameButton:
		return color.RGBA{20, 20, 20, 255}
	case theme.ColorNameForegroundOnWarning:
		return color.RGBA{224, 38, 38, 255}
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
