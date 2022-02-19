package impl

import (
	"github.com/smarty-pen/drawer/config"
)

type BaseDrawer interface {
	DoDraw(dots *[]config.Info, background string, output string) error
}
