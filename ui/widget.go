package ui

type Widget struct {
	// position
	x int
	y int

	// size
	w int
	h int

	// anchors
	anchorLeft   bool
	anchorTop    bool
	anchorRight  bool
	anchorBottom bool

	widgets []*Widget

	onCustomPaint func(cnv *Canvas)
}

func NewWidget() *Widget {
	var c Widget
	return &c
}

func (c *Widget) AddWidget(w *Widget) {
	c.widgets = append(c.widgets, w)
}

func (c *Widget) X() int {
	return c.x
}

func (c *Widget) Y() int {
	return c.y
}

func (c *Widget) W() int {
	return c.w
}

func (c *Widget) H() int {
	return c.h
}

func (c *Widget) SetPosition(x, y int) {
	c.x = x
	c.y = y
}

func (c *Widget) SetSize(w, h int) {
	c.w = w
	c.h = h
}

func (c *Widget) SetAnchors(left, top, right, bottom bool) {
	c.anchorLeft = left
	c.anchorTop = top
	c.anchorRight = right
	c.anchorBottom = bottom
}

func (c *Widget) SetOnPaint(f func(cnv *Canvas)) {
	c.onCustomPaint = f
}

func (c *Widget) onPaint(cnv *Canvas) {
	if c.onCustomPaint != nil {
		c.onCustomPaint(cnv)
	}

	for _, w := range c.widgets {
		cnv.Save()
		cnv.Translate(w.x, w.y)
		cnv.SetClip(w.x, w.y, w.w, w.h)
		w.onPaint(cnv)
		cnv.Restore()
	}
}

func (c *Widget) onResize(oldWidth, oldHeight, newWidth, newHeight int) {
	deltaWidth := newWidth - oldWidth
	deltaHeight := newHeight - oldHeight

	newX := c.X()
	newY := c.Y()
	newW := c.W()
	newH := c.H()

	if c.anchorLeft && c.anchorRight {
		newW += deltaWidth
	}
	if !c.anchorLeft && c.anchorRight {
		newX += deltaWidth
	}

	if c.anchorTop && c.anchorBottom {
		newH += deltaHeight
	}

	if !c.anchorTop && c.anchorBottom {
		newY += deltaHeight
	}

	c.x = newX
	c.y = newY
	c.w = newW
	c.h = newH
}
