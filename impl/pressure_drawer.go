package impl

import (
	"fmt"
	"github.com/smarty-pen/drawer/config"
	"github.com/smarty-pen/drawer/util"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/rasterizer"
)

type PressureDrawer struct {

}



func calculatePressure( force int) float64 {
	var pressure float64 = 1;
	if (force >= 0 && force <= 20) {
		pressure = 100;
	} else if (force > 20 && force <= 40) {
		pressure = 105;
	} else if (force > 40 && force <= 60) {
		pressure = 110;
	} else if (force > 60 && force <= 90) {
		pressure = 120;
	} else if (force > 90 && force <= 150) {
		pressure = 125;
	} else {
		pressure = 130;
	}
	return pressure;
}

func drawLine(ctx *canvas.Context, x1, y1, x2, y2, penWidth float64) {
	ctx.MoveTo(x1, y1)
	ctx.SetStrokeWidth(float64(penWidth))
	ctx.LineTo(x2,y2)	
}

func (drawer *PressureDrawer) DoDraw(dots *[]config.Info, background string, output string) error {
	c := canvas.New(config.IMAGE_WIDTH, config.IMAGE_HEIGHT)

	dotArray := *dots;
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
	
	
	ctx.SetFillColor(canvas.White)//画布背景为白色
	ctx.SetStrokeColor(canvas.Black)
	var lastX float64
	var lastY float64
	for i:=0; i<len(dotArray);i++ {
		dot := dotArray[i]
		point := dot.Dot
		x := (imageWidthFloat * (float64(point.X) + float64(point.Fx) / 100) / config.MAX_X)
		y := imageHeightFloat - (imageHeightFloat * (float64(point.Y) + float64(point.Fy) / 100) / config.MAX_Y)
		penWidth := calculatePressure(point.Force) / config.TRANSFORM_SCALE

		ctx.SetStrokeWidth(float64(1))
		if point.Type == "PEN_DOWN" {
			lastX = x
			lastY = y
		} else if point.Type == "PEN_MOVE" {
			drawLine(ctx, lastX, lastY, x, y, penWidth)
			lastX = x
			lastY = y
		} else if point.Type == "PEN_UP" {
			drawLine(ctx, lastX, lastY, x, y, penWidth)
		} else {
			fmt.Println("非法的笔画类型", point.Type)
		}
		ctx.FillStroke()
	}

	c.WriteFile(output, rasterizer.PNGWriter(1.2))
	return nil
}