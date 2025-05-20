package ui

import (
	"image/color"

	"github.com/ipoluianov/nui/nuimouse"
)

func NewButton() *Widget {
	c := NewWidget()
	c.SetSize(100, 30)
	c.SetOnPaint(func(cnv *Canvas) {
		cnv.SetColor(color.RGBA{0, 0, 0, 255})
		backColor := color.RGBA{0, 0, 0, 255}
		if c.IsHovered() {
			backColor = color.RGBA{50, 50, 50, 255}
		}

		clicked := c.GetPropBool("clicked", false)
		if clicked {
			backColor = color.RGBA{100, 100, 100, 255}
		}

		cnv.FillRect(0, 0, c.W(), c.H(), backColor)

		text := c.GetPropString("text", "")

		cnv.DrawTextMultiline(0, 0, c.W(), c.H(), HAlignCenter, VAlignCenter, text, color.RGBA{255, 255, 255, 255}, "robotomono", 16, false)
	})

	c.SetOnMouseDown(func(button nuimouse.MouseButton, x int, y int) {
		c.SetProp("clicked", true)
		fnOnClick, ok := c.GetProp("onClick").(func())
		if ok {
			fnOnClick()
		}
	})

	c.SetOnMouseUp(func(button nuimouse.MouseButton, x int, y int) {
		c.SetProp("clicked", false)
	})

	//fnOnClick := c.GetProp("onClick")

	return c
}
