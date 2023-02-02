package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type Brush = int

type pxCanvasConfig struct {
	DrawingArea fyne.Size
	CanvasOffset fyne.Position
	PxRows, PxCols int
	PxSize int
}

type State struct {
	BrushColor color.Color
	BrushType int
	SwatchSelected int
	FilePath string
}

func (state *State) SetFilePath(path string) {
	state.FilePath = path
}
