package ex02widget

import (
	"image/color"

	"github.com/ipoluianov/nuiforms/ui"
)

func NewMyWidget(col color.RGBA) *ui.Widget {
	var c ui.Widget
	c.SetOnPaint(func(cnv *ui.Canvas) {
		cnv.SetColor(col)
		for i := c.X(); i < c.X()+c.W(); i++ {
			for j := c.Y(); j < c.Y()+c.H(); j++ {
				cnv.SetPixel(i, j)
			}
		}
	})
	return &c
}
