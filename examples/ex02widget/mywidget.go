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

	secWidget := NewSecondWidget(color.RGBA{col.R, col.G + 50, col.B + 20, 255})
	secWidget.SetPosition(10, 10)
	secWidget.SetSize(50, 50)
	c.AddWidget(secWidget)
	return &c
}

func NewSecondWidget(col color.RGBA) *ui.Widget {
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
