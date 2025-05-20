package ex02widget

import (
	"fmt"
	"image/color"

	"github.com/ipoluianov/nui/nuimouse"
	"github.com/ipoluianov/nuiforms/ui"
)

func NewMyWidget(col color.RGBA) *ui.Widget {
	c := ui.NewWidget()
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
	return c
}

func NewSecondWidget(col color.RGBA) *ui.Widget {
	c := ui.NewWidget()
	inside := false
	c.SetOnPaint(func(cnv *ui.Canvas) {
		if inside {
			cnv.SetColor(col)
		} else {
			cnv.SetColor(color.RGBA{10, 10, 10, 255})
		}
		for i := 0; i < c.W(); i++ {
			for j := 0; j < c.H(); j++ {
				cnv.SetPixel(i, j)
			}
		}
	})
	c.SetOnMouseDown(func(button nuimouse.MouseButton, x int, y int) {
		if button == nuimouse.MouseButtonLeft {
			fmt.Println("Left button clicked", c.Name(), x, y)
		}
	})
	c.SetOnMouseEnter(func() {
		inside = true
		ui.MainForm.Update()
	})
	c.SetOnMouseLeave(func() {
		inside = false
		ui.MainForm.Update()
	})
	return c
}
