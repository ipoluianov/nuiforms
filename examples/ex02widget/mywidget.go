package ex02widget

import (
	"image"
	"image/color"

	"github.com/ipoluianov/nuiforms/ui"
)

func NewMyWidget(col color.RGBA) *ui.Widget {
	var c ui.Widget
	c.SetOnPaint(func(rgba *image.RGBA) {
		for i := c.X(); i < c.X()+c.W(); i++ {
			for j := c.Y(); j < c.Y()+c.H(); j++ {
				rgba.Set(i, j, col)
			}
		}
	})
	return &c
}
