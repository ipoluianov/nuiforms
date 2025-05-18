package ui

import (
	"image"

	"github.com/ipoluianov/nui/nui"
)

type Form struct {
	wnd    nui.Window
	title  string
	width  int
	height int

	widgets []*Widget
}

func NewForm() *Form {
	var c Form
	c.title = "Form"
	c.width = 800
	c.height = 600
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

func (c *Form) Exec() {
	c.wnd = nui.CreateWindow(c.title, c.width, c.height, true)
	c.wnd.OnPaint(c.onPaint)
	c.wnd.OnResize(c.onResize)
	c.wnd.Show()
	c.wnd.EventLoop()
}

func (c *Form) AddWidget(w *Widget) {
	c.widgets = append(c.widgets, w)
}

func (c *Form) onPaint(rgba *image.RGBA) {
	cnv := NewCanvas(rgba)
	for _, w := range c.widgets {
		w.onPaint(cnv)
	}
}

func (c *Form) onResize(width, height int) {
	for _, w := range c.widgets {
		w.onResize(c.width, c.height, width, height)
	}
	c.width = width
	c.height = height
}
