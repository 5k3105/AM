package controlbox

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

const dir = "C:/goworksp/src/local/AM"

type Controlbox struct {
	*widgets.QWidget
	Mode            string
	Tool            string
	Canvas          *widgets.QWidget
	Statusbar       *widgets.QStatusBar
	PropertiesStack *widgets.QStackedLayout
	List            *widgets.QListWidget
}

func NewControlbox(canvas *widgets.QWidget, statusbar *widgets.QStatusBar) *Controlbox {

	cb := &Controlbox{}

	cb.QWidget = widgets.NewQWidget(nil, 0)

	cb.Canvas = canvas
	cb.Statusbar = statusbar

	layout := widgets.NewQVBoxLayout()

	//	f := widgets.NewQLabel2("File", cb, 0)
	//	f.SetFrameStyle(6)

	//	layout.AddWidget(f, 0, 0)
	//	layout.AddWidget(cb.newFile(), 0, 0)

	m := widgets.NewQLabel2("Mode", cb, 0)
	m.SetFrameStyle(6)

	layout.AddWidget(m, 0, 0)
	layout.AddWidget(cb.newMode(), 0, 0)

	t := widgets.NewQLabel2("Tools", cb, 0)
	t.SetFrameStyle(6)

	layout.AddWidget(t, 0, 0)
	layout.AddWidget(cb.newTools(), 0, 0)

	p := widgets.NewQLabel2("Properties", cb, 0)
	p.SetFrameStyle(6)

	layout.AddWidget(p, 0, 0)
	layout.AddWidget(cb.newProperties(), 0, 0)

	var spacerPushUp = widgets.NewQSpacerItem(1, 1, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__MinimumExpanding)
	layout.AddItem(spacerPushUp)

	cb.SetLayout(layout)

	cb.Mode = "draw"
	cb.Tool = "square"

	return cb
}

func (c *Controlbox) newFile() *widgets.QWidget {
	w := widgets.NewQWidget(nil, 0)

	save := widgets.NewQPushButton3(gui.NewQIcon5(dir+"/images/save.png"), "", nil)
	save.SetObjectName("save")

	saveas := widgets.NewQPushButton3(gui.NewQIcon5(dir+"/images/saveas.png"), "", nil)
	saveas.SetObjectName("saveas")

	load := widgets.NewQPushButton3(gui.NewQIcon5(dir+"/images/load.png"), "", nil)
	load.SetObjectName("load")

	layout := widgets.NewQHBoxLayout()
	layout.AddWidget(save, 0, 0)
	layout.AddWidget(saveas, 0, 0)
	layout.AddWidget(load, 0, 0)

	w.SetLayout(layout)
	return w
}

func (c *Controlbox) newMode() *widgets.QWidget {

	w := widgets.NewQWidget(nil, 0)

	// erase
	erase := widgets.NewQPushButton3(gui.NewQIcon5(dir+"/images/eraser.png"), "", nil)
	erase.SetCheckable(true)
	erase.SetObjectName("erase")

	// move
	move := widgets.NewQPushButton3(gui.NewQIcon5(dir+"/images/sizeall.png"), "", nil)
	move.SetCheckable(true)
	move.SetObjectName("move")

	// draw
	draw := widgets.NewQPushButton3(gui.NewQIcon5(dir+"/images/pen.png"), "", nil)
	draw.SetCheckable(true)
	draw.SetObjectName("draw")
	draw.SetChecked(true)

	// pan
	pan := widgets.NewQPushButton3(gui.NewQIcon5(dir+"/images/hand.png"), "", nil)
	pan.SetCheckable(true)
	pan.SetObjectName("pan")

	layout := widgets.NewQHBoxLayout()
	layout.AddWidget(erase, 0, 0)
	layout.AddWidget(move, 0, 0)
	layout.AddWidget(draw, 0, 0)
	layout.AddWidget(pan, 0, 0)

	bg := widgets.NewQButtonGroup(w)
	bg.AddButton(erase, 0)
	bg.AddButton(move, 0)
	bg.AddButton(draw, 1)
	bg.AddButton(pan, 2)

	bg.ConnectButtonClicked(c.modeClicked)

	w.SetLayout(layout)
	return w
}

func (c *Controlbox) modeClicked(b *widgets.QAbstractButton) {
	c.Mode = b.ObjectName()
	c.Statusbar.ShowMessage(c.Mode+" "+c.Tool, 0)

	switch c.Mode {
	case "erase":
		ir := gui.NewQImageReader3(dir+"/images/eraser.png", "png")
		img := ir.Read()
		pix := gui.QPixmap_FromImage(img, 0)
		c.Canvas.SetCursor(gui.NewQCursor4(pix, -1, -1))

	case "move":
		c.Canvas.SetCursor(gui.NewQCursor2(9))

	case "draw":
		c.Canvas.SetCursor(gui.NewQCursor2(0))

	case "pan":
		c.Canvas.SetCursor(gui.NewQCursor2(17))

	}

}

func (c *Controlbox) newTools() *widgets.QWidget {

	w := widgets.NewQWidget(nil, 0)

	layout := widgets.NewQHBoxLayout()
	layout.SetContentsMargins(0, 0, 0, 0)

	g := widgets.NewQPushButton3(gui.NewQIcon5(dir+"/images/grid.png"), "", nil)
	g.SetCheckable(true)
	g.SetObjectName("grid")

	t := widgets.NewQPushButton3(gui.NewQIcon5(dir+"/images/type.png"), "", nil)
	t.SetCheckable(true)
	t.SetObjectName("type")

	e := widgets.NewQPushButton3(gui.NewQIcon5(dir+"/images/circle.png"), "", nil)
	e.SetCheckable(true)
	e.SetObjectName("circle")

	s := widgets.NewQPushButton3(gui.NewQIcon5(dir+"/images/square.svg"), "", nil)
	s.SetCheckable(true)
	s.SetObjectName("square")
	s.SetChecked(true)

	l := widgets.NewQPushButton3(gui.NewQIcon5(dir+"/images/line.png"), "", nil)
	l.SetCheckable(true)
	l.SetObjectName("line")

	layout.AddWidget(g, 0, 0)
	layout.AddWidget(t, 0, 0)
	layout.AddWidget(e, 0, 0)
	layout.AddWidget(s, 0, 0)
	layout.AddWidget(l, 0, 0)

	bg := widgets.NewQButtonGroup(w)
	bg.AddButton(g, 0)
	bg.AddButton(t, 1)
	bg.AddButton(e, 2)
	bg.AddButton(s, 1)
	bg.AddButton(l, 2)

	bg.ConnectButtonClicked(c.toolClicked)

	w.SetLayout(layout)

	return w
}

func (c *Controlbox) toolClicked(b *widgets.QAbstractButton) {
	c.Tool = b.ObjectName()
	c.Statusbar.ShowMessage(c.Mode+" "+c.Tool, 0)

	switch c.Tool {
	case "grid":
		c.PropertiesStack.SetCurrentIndex(0)
	case "type":
		c.PropertiesStack.SetCurrentIndex(3)
	case "circle":
		c.PropertiesStack.SetCurrentIndex(1)
	case "square":
		c.PropertiesStack.SetCurrentIndex(2)
	case "line":
		c.PropertiesStack.SetCurrentIndex(4)

	}

}

func (c *Controlbox) newProperties() *widgets.QWidget {

	w := widgets.NewQWidget(nil, 0)

	c.PropertiesStack = widgets.NewQStackedLayout2(w)

	c.PropertiesStack.AddWidget(newGridProps())     // 0
	c.PropertiesStack.AddWidget(newCircleProps())   // 1
	c.PropertiesStack.AddWidget(c.newSquareProps()) // 2
	c.PropertiesStack.AddWidget(newTypeProps())     // 3
	c.PropertiesStack.AddWidget(newLineProps())     // 4

	c.PropertiesStack.SetCurrentIndex(2)

	w.SetLayout(c.PropertiesStack)

	return w
}

func newGridProps() *widgets.QWidget {

	w := widgets.NewQWidget(nil, 0)

	layout := widgets.NewQHBoxLayout()

	w.SetLayout(layout)

	return w
}

func newCircleProps() *widgets.QWidget {

	w := widgets.NewQWidget(nil, 0)

	layout := widgets.NewQHBoxLayout()

	w.SetLayout(layout)

	return w
}

func (c *Controlbox) newSquareProps() *widgets.QWidget {

	w := widgets.NewQWidget(nil, 0)
	layout := widgets.NewQHBoxLayout()

	g := widgets.NewQPushButton2("Clear", nil)
	g.ConnectClick(c.listClick)

	var lw = widgets.NewQListWidget(w)
	c.List = lw

	layout.AddWidget(lw, 0, 0)
	layout.AddWidget(g, 0, 0)
	w.SetLayout(layout)

	return w
}

func (c *Controlbox) listClick() {

}

func newTypeProps() *widgets.QWidget {

	w := widgets.NewQWidget(nil, 0)

	layout := widgets.NewQHBoxLayout()

	vlayout := widgets.NewQVBoxLayout()

	f := widgets.NewQLabel2("Font: ", w, 0)
	layout.AddWidget(f, 0, 0)

	//  # combo set font
	var comboFontFamily = widgets.NewQFontComboBox(nil)
	//self.comboFontFamily.currentFontChanged.connect(self.onFontFamilyChanged)
	layout.AddWidget(comboFontFamily, 0, 0)

	//  # spin font-size
	var spinFontSize = widgets.NewQSpinBox(nil)
	//spinFontSize.valueChanged.connect(self.onFontSizeChanged)
	layout.AddWidget(spinFontSize, 0, 0)

	//  # button fg
	var buttonColorFg = widgets.NewQPushButton(nil)
	//self.buttonColorFg.clicked.connect(self.onButtonColorFgClicked)
	buttonColorFg.SetText("Fg Color")
	buttonColorFg.SetMaximumWidth(50)
	layout.AddWidget(buttonColorFg, 0, 0)

	p := widgets.NewQLabel2("[ ]", w, 0)
	layout.AddWidget(p, 0, 0)

	vlayout.AddLayout(layout, 0)

	var spacerPushUp = widgets.NewQSpacerItem(1, 1, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__MinimumExpanding)
	vlayout.AddItem(spacerPushUp)

	w.SetLayout(vlayout)

	return w
}

func newLineProps() *widgets.QWidget {

	w := widgets.NewQWidget(nil, 0)

	layout := widgets.NewQHBoxLayout()

	w.SetLayout(layout)

	return w
}

//-------------------------------------------------------------------------Notes
