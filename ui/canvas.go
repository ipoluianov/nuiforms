package ui

import (
	"image"
	"image/color"
)

type canvasState struct {
	col color.RGBA

	translateX int
	translateY int

	clipX int
	clipY int
	clipW int
	clipH int
}

type Canvas struct {
	rgba  *image.RGBA
	state *canvasState
	stack []*canvasState
}

func NewCanvas(rgba *image.RGBA) *Canvas {
	var c Canvas
	c.rgba = rgba
	c.state = &canvasState{}
	c.stack = make([]*canvasState, 0)
	return &c
}

func (c *Canvas) Save() {
	c.stack = append(c.stack, c.state)
	c.state = &canvasState{}
}

func (c *Canvas) Restore() {
	if len(c.stack) == 0 {
		return
	}
	c.state = c.stack[len(c.stack)-1]
	c.stack = c.stack[:len(c.stack)-1]
}

func (c *Canvas) SetColor(col color.RGBA) {
	c.state.col = col
}

func (c *Canvas) SetClip(x, y, w, h int) {
	c.state.clipX += x
	c.state.clipY += y
	c.state.clipW = w
	c.state.clipH = h
}

func (c *Canvas) Translate(x, y int) {
	c.state.translateX += x
	c.state.translateY += y
}

func (c *Canvas) SetPixel(x, y int) {
	x += c.state.translateX
	y += c.state.translateY
	if c.state.clipX != 0 || c.state.clipY != 0 || c.state.clipW != 0 || c.state.clipH != 0 {
		if x < c.state.clipX || y < c.state.clipY || x >= c.state.clipX+c.state.clipW || y >= c.state.clipY+c.state.clipH {
			return
		}
	}
	if x < 0 || y < 0 || x >= c.rgba.Rect.Dx() || y >= c.rgba.Rect.Dy() {
		return
	}
	c.rgba.Set(x, y, c.state.col)
}
