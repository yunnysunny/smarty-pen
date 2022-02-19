package impl

import (
	"fmt"

	"github.com/smarty-pen/drawer/config"
	"github.com/smarty-pen/drawer/util"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/rasterizer"
)

type SimpleDrawer struct {
}

func (drawer *SimpleDrawer) DoDraw(dots *[]config.Info, background string, output string) error {
	c := canvas.New(config.IMAGE_WIDTH, config.IMAGE_HEIGHT)

	dotArray := *dots
	imageWidthFloat := float64(config.IMAGE_WIDTH)
	imageHeightFloat := float64(config.IMAGE_HEIGHT)
	// Create a canvas context used to keep drawing state
	ctx := canvas.NewContext(c)
	if background != "" {
		err := util.DrawBackgroundImage(ctx, background)
		if err != nil {
			return err
		}
	}

	ctx.SetFillColor(canvas.White) //画布背景为白色
	ctx.SetStrokeColor(canvas.Black)

	for i := 0; i < len(dotArray); i++ {
		dot := dotArray[i]
		point := dot.Dot
		x := (imageWidthFloat * (float64(point.X) + float64(point.Fx)/100) / config.MAX_X)
		y := imageHeightFloat - (imageHeightFloat * (float64(point.Y) + float64(point.Fy)/100) / config.MAX_Y)

		if point.Type == "PEN_DOWN" {
			ctx.MoveTo(x, y)
		} else if point.Type == "PEN_MOVE" {
			ctx.LineTo(x, y)
		} else if point.Type == "PEN_UP" {
			ctx.LineTo(x, y)
			ctx.FillStroke()
		} else {
			fmt.Println("非法的笔画类型", point.Type)
		}
	}

	c.WriteFile(output, rasterizer.PNGWriter(1.2))
	return nil
}
