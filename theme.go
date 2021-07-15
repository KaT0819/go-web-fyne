package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct {
}

// var _ fyne.Theme = (*myTheme)(nil)

func (*myTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return resourceCWindowsFontsMeiryoTtc
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return resourceCWindowsFontsMeiryoTtc
}

func (t *myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	colors := theme.darkPalette
	if v == theme.VariantLight {
		colors = theme.lightPalette
	}

	if n == theme.ColorNamePrimary {
		return theme.PrimaryColorNamed(fyne.CurrentApp().Settings().PrimaryColor())
	} else if n == theme.ColorNameFocus {
		return theme.focusColorNamed(fyne.CurrentApp().Settings().PrimaryColor())
	}

	if c, ok := colors[n]; ok {
		return c
	}
	return color.Transparent
}
