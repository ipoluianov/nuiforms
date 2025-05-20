package ui

import (
	"image"
	"image/color"

	"github.com/ipoluianov/nui/nui"
	"github.com/ipoluianov/nui/nuikey"
	"github.com/ipoluianov/nui/nuimouse"
)

type Form struct {
	wnd    nui.Window
	title  string
	width  int
	height int

	topWidget *Widget
}

var MainForm *Form

func NewForm() *Form {
	var c Form
	c.title = "Form"
	c.width = 800
	c.height = 600
	c.topWidget = NewWidget()
	c.topWidget.SetPosition(0, 0)
	c.topWidget.SetSize(c.width, c.height)
	c.topWidget.SetAnchors(true, true, true, true)
	c.topWidget.SetOnPaint(func(cnv *Canvas) {
		cnv.SetColor(color.RGBA{0, 50, 0, 255})
		for i := 0; i < c.topWidget.w; i++ {
			for j := 0; j < c.topWidget.h; j++ {
				cnv.SetPixel(i, j)
			}
		}
	})
	return &c
}

func (c *Form) SetTitle(title string) {
	c.title = title
	if c.wnd != nil {
		c.wnd.SetTitle(title)
	}
}

func (c *Form) SetSize(width, height int) {
	c.width = width
	c.height = height
	if c.wnd != nil {
		c.wnd.Resize(width, height)
	}
}

func (c *Form) Panel() *Widget {
	return c.topWidget
}

func (c *Form) Exec() {
	c.wnd = nui.CreateWindow(c.title, c.width, c.height, true)
	MainForm = c
	c.wnd.OnPaint(c.onPaint)
	c.wnd.OnResize(c.onResize)
	c.wnd.OnMouseButtonDown(c.onMouseDown)
	c.wnd.OnMouseButtonUp(c.onMouseUp)
	c.wnd.OnMouseMove(c.onMouseMove)
	c.wnd.OnMouseLeave(c.onMouseLeave)
	c.wnd.OnMouseEnter(c.onMouseEnter)
	c.wnd.OnKeyDown(c.onKeyDown)
	c.wnd.OnKeyUp(c.onKeyUp)
	c.wnd.OnMouseWheel(c.onMouseWheel)
	c.wnd.Show()
	c.wnd.EventLoop()
}

func (c *Form) Update() {
	if c.wnd != nil {
		c.wnd.Update()
	}
}

func (c *Form) onPaint(rgba *image.RGBA) {
	cnv := NewCanvas(rgba)
	c.topWidget.processPaint(cnv)
}

func (c *Form) onResize(width, height int) {
	c.topWidget.SetSize(width, height)
	c.width = width
	c.height = height
}

func (c *Form) onMouseDown(button nuimouse.MouseButton, x int, y int) {
	c.topWidget.processMouseDown(button, x, y)
}

func (c *Form) onMouseUp(button nuimouse.MouseButton, x int, y int) {
	c.topWidget.processMouseUp(button, x, y)
}

func (c *Form) onMouseMove(x int, y int) {
	c.topWidget.processMouseMove(x, y)
	c.Update()
}

func (c *Form) onMouseLeave() {
	c.topWidget.processMouseLeave()
}

func (c *Form) onMouseEnter() {
	c.topWidget.processMouseEnter()
}

func (c *Form) onKeyDown(keyCode nuikey.Key, mods nuikey.KeyModifiers) {
	c.topWidget.processKeyDown(keyCode)
}

func (c *Form) onKeyUp(keyCode nuikey.Key, mods nuikey.KeyModifiers) {
	c.topWidget.processKeyUp(keyCode)
}

func (c *Form) onMouseWheel(deltaX int, deltaY int) {
	c.topWidget.processMouseWheel(deltaX, deltaY)
}
