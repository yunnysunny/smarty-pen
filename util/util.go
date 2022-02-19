package util

import (
	"image/png"
	"os"

	"github.com/tdewolff/canvas"
)

func DrawBackgroundImage(ctx *canvas.Context, filename string) error {
	// Load the image data
	lenna, err := os.Open(filename)
	if err != nil {
		return err
	}

	// Create a new image object
	img, err := png.Decode(lenna)
	if err != nil {
		return err
	}

	ctx.DrawImage(0, 0, img, 1)
	return nil
}
