package ex02widget

import (
	"image/color"

	"github.com/ipoluianov/nuiforms/ui"
)

func NewMyWidget(col color.RGBA) *ui.Widget {
	var c ui.Widget
	c.SetOnPaint(func(cnv *ui.Canvas) {
		cnv.SetColor(col)
		for i := 0; i < c.W(); i++ {
			for j := 0; j < c.H(); j++ {
				cnv.SetPixel(i, j)
			}
		}
	})
	return &c
}
