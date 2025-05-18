package ex02widget

import (
	"image/color"

	"github.com/ipoluianov/nuiforms/ui"
)

func Run() {
	form := ui.NewForm()
	form.SetTitle("Hello, World!")
	form.SetSize(800, 600)

	{
		myWidget := NewMyWidget(color.RGBA{0, 100, 50, 100})
		myWidget.SetPosition(10, 10)
		myWidget.SetSize(300, 480)
		myWidget.SetAnchors(true, true, false, true)
		form.AddWidget(myWidget)
	}
	{
		myWidget := NewMyWidget(color.RGBA{0, 50, 100, 100})
		myWidget.SetPosition(320, 10)
		myWidget.SetSize(470, 480)
		myWidget.SetAnchors(true, true, true, true)
		form.AddWidget(myWidget)
	}

	{
		myWidget := NewMyWidget(color.RGBA{100, 0, 50, 100})
		myWidget.SetPosition(10, 500)
		myWidget.SetSize(780, 90)
		myWidget.SetAnchors(true, false, true, true)
		form.AddWidget(myWidget)
	}

	form.Exec()
}
