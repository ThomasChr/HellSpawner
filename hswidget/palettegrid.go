package hswidget

import (
	"fmt"
	image2 "image"
	"image/color"

	"github.com/OpenDiablo2/giu"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

const (
	gridWidth  = 16
	gridHeight = 16
	cellSize   = 12
)

type PaletteGridWidget struct {
	id     string
	colors [256]d2interface.Color
}

func PaletteGrid(id string, colors [256]d2interface.Color) *PaletteGridWidget {
	result := &PaletteGridWidget{
		id:     id,
		colors: colors,
	}

	return result
}

func (p *PaletteGridWidget) Build() {
	stateId := fmt.Sprintf("PaletteGridWidget_%s", p.id)
	state := giu.Context.GetState(stateId)
	var widget *giu.ImageWidget

	if state == nil {
		widget = giu.Image(nil, gridWidth*cellSize, gridHeight*cellSize)

		//Prevent multiple invocation to LoadImage.
		giu.Context.SetState(stateId, &giu.ImageState{})

		rgb := image2.NewRGBA(image2.Rect(0, 0, gridWidth*cellSize, gridHeight*cellSize))

		for y := 0; y < gridHeight*cellSize; y++ {
			if y%cellSize == 0 {
				continue
			}
			for x := 0; x < gridWidth*cellSize; x++ {
				if x%cellSize == 0 {
					continue
				}
				idx := (x / cellSize) + ((y / cellSize) * gridWidth)
				col := p.colors[idx]
				rgb.Set(x, y, color.RGBA{R: col.R(), G: col.G(), B: col.B(), A: 255})
			}
		}

		go func() {
			texture, err := giu.NewTextureFromRgba(rgb)
			if err == nil {
				giu.Context.SetState(stateId, &giu.ImageState{Texture: texture})
			}
		}()
	} else {
		imgState := state.(*giu.ImageState)
		widget = giu.Image(imgState.Texture, gridWidth*cellSize, gridHeight*cellSize)
	}

	widget.Build()

}
