package gfxcanvas

import (
	"local/AM/controlbox"
	"local/AM/gfxinterface"
	"local/AM/gfxobjects/line"
	"local/AM/gfxobjects/rectangle"
	"local/AM/graph"

	"strconv"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type Canvas struct {
	*widgets.QWidget
	Scene      *widgets.QGraphicsScene
	View       *widgets.QGraphicsView
	Statusbar  *widgets.QStatusBar
	Controlbox *controlbox.Controlbox
	Graph      *graph.Graph
	*Grid
	Figures     *treemap.Map //[]gfxinterface.Figure
	drag        bool
	drawingline *line.Line
	move        bool
	movingrect  *rectangle.Rectangle
}

type Color struct {
	R, G, B, A int
}

func NewCanvas(statusbar *widgets.QStatusBar) *Canvas {
	canvas := &Canvas{
		Statusbar: statusbar,
		QWidget:   widgets.NewQWidget(nil, 0),
		Scene:     widgets.NewQGraphicsScene(nil),
		View:      widgets.NewQGraphicsView(nil),
		Figures:   treemap.NewWithIntComparator()}

	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(canvas.View, 0, 0)

	canvas.Grid = canvas.NewGrid(0, 0, 0, 24, 64, 5, 160, 160, 160, 120)

	canvas.Scene.ConnectKeyPressEvent(canvas.keyPressEvent)
	canvas.Scene.ConnectMousePressEvent(canvas.mousePressEvent)
	canvas.Scene.ConnectWheelEvent(canvas.wheelEvent)
	canvas.Scene.ConnectMouseMoveEvent(canvas.mouseMoveEvent)

	canvas.SetLayout(layout)

	canvas.Graph = graph.NewGraph()

	//canvas.View.Viewport().SetMouseTracking(true)

	canvas.View.Scale(5.00, 5.00)

	canvas.View.SetViewportUpdateMode(0)

	//canvas.Scene.SetSceneRect2(0.0, 0.0, 64.0*5, 24.0*5)

	return canvas
}

func (c *Canvas) mouseMoveEvent(e *widgets.QGraphicsSceneMouseEvent) {

	if c.move {
		var px, py = e.ScenePos().X(), e.ScenePos().Y()
		var x, y = float64((int(px) / 5) * 5), float64((int(py) / 5) * 5) // round to nearest fit

		c.movingrect.SetX(x)
		c.movingrect.SetY(y)
		c.movingrect.PrepareGeometryChange()
		//c.Statusbar.ShowMessage(FloatToString(target.GetX())+" "+FloatToString(target.GetY()), 0)

	}

	if c.drag {
		var l = c.drawingline
		target := l.Target // Point
		target.SetX(e.ScenePos().X())
		target.SetY(e.ScenePos().Y())
		l.PrepareGeometryChange()

		c.Statusbar.ShowMessage(FloatToString(target.GetX())+" "+FloatToString(target.GetY()), 0)
	}
	c.View.Viewport().Repaint()
	c.Scene.MouseMoveEventDefault(e)
}

func (c *Canvas) SetControlbox(cb *controlbox.Controlbox) {
	c.Controlbox = cb
}

func (c *Canvas) wheelEvent(e *widgets.QGraphicsSceneWheelEvent) {
	if e.Modifiers() == core.Qt__ControlModifier {
		if e.Delta() > 0 {
			c.View.Scale(1.25, 1.25)
		} else {
			c.View.Scale(0.8, 0.8)
		}
	}
}

func (c *Canvas) mousePressEvent(e *widgets.QGraphicsSceneMouseEvent) {
	var px, py = e.ScenePos().X(), e.ScenePos().Y()
	switch c.Controlbox.Mode {
	case "draw":
		switch c.Controlbox.Tool {
		case "square":

			switch e.Button() {
			case 1: // left button

				if c.move {
					c.movingrect.Over2 = false
					c.move = false
					//c.movingrect.SetX(x)
					//c.movingrect.SetY(y)
					c.movingrect = nil
					return
				}

				var x, y = float64((int(px) / 5) * 5), float64((int(py) / 5) * 5) // round to nearest fit

				var m = c.FindRectOverlap(x, y)
				if m == nil { // no node

					r := c.AddRectangle(x, y, 20, 10)
					c.Statusbar.ShowMessage(FloatToString(r.X)+" "+FloatToString(r.Y), 0)

				} else {

					var m = c.FindRectBeneath(px, py)
					if m != nil { // found node
						if c.drawingline == nil {

							r := m.(*rectangle.Rectangle)
							if r.FindRectBeneath(px, py) {
								r.Over2 = true
								c.move = true
								c.movingrect = r

							} else {

								c.DrawLine(m, px, py)
								c.Statusbar.ShowMessage("drag start", 0)
							}
						} else {

							c.AddLine(c.drawingline.Source, m)
							c.Statusbar.ShowMessage("drag end", 0)

						}
					}
				}
			case 2: // right button
				var m = c.FindRectBeneath(px, py)
				if m != nil { // found node
					c.RemoveRectangle(m)
				}

			}

			c.ListNodes()

		case "line":

		}
	case "move":
		c.Statusbar.ShowMessage("move", 0)
	case "pan":
		c.Statusbar.ShowMessage("pan", 0)

	}

	c.Scene.MousePressEventDefault(e)
	c.View.Viewport().Repaint()
}

func (c *Canvas) RemoveRectangle(t gfxinterface.Figure) {

	var target = t.(*rectangle.Rectangle)
	c.RemoveEdges(target)

	c.Scene.RemoveItem(target) // remove from scene
	target.PrepareGeometryChange()
	c.Figures.Remove(target.GetNode().Id) // remove from treemap

	c.View.Viewport().Repaint()
}

func (c *Canvas) RemoveEdges(r *rectangle.Rectangle) {

	for _, f := range r.IncomingEdges.Values() {
		l := f.(*line.Line)
		c.Scene.RemoveItem(l) // remove from scene
		l.PrepareGeometryChange()
	}

	for _, f := range r.OutgoingEdges.Values() {
		l := f.(*line.Line)
		c.Scene.RemoveItem(l) // remove from scene
		l.PrepareGeometryChange()
	}

	r.IncomingEdges.Clear()
	r.OutgoingEdges.Clear() // remove from treemap
}

func (c *Canvas) AddRectangle(x, y, w, h float64) *rectangle.Rectangle {

	i, r := rectangle.NewRectangle(x, y, w, h, c.Graph, c.Statusbar)
	c.Figures.Put(i, r)
	c.Scene.AddItem(r)
	return r
}

func (c *Canvas) AddLine(source gfxinterface.Figure, target gfxinterface.Figure) {

	l := line.AddLine(c.Graph, source, target)

	l.Source.AddEdgeOutgoing(l)
	l.Target.AddEdgeIncoming(l)

	c.Scene.AddItem(l)

	c.Scene.RemoveItem(c.drawingline)
	c.drawingline.PrepareGeometryChange()
	c.drawingline = nil
	c.drag = false
	c.View.Viewport().Repaint()
	//c.Scene.Update(c.Scene.SceneRect())

}

func (c *Canvas) DrawLine(source gfxinterface.Figure, tx, ty float64) {
	l := line.DrawLine(source, tx, ty)
	c.Scene.AddItem(l)
	c.drawingline = l
	c.drag = true
}

func (c *Canvas) keyPressEvent(e *gui.QKeyEvent) {

	if e.Modifiers() == core.Qt__ControlModifier {
		switch int32(e.Key()) {
		case int32(core.Qt__Key_Equal):
			c.View.Scale(1.25, 1.25)

		case int32(core.Qt__Key_Minus):
			c.View.Scale(0.8, 0.8)
		}
	}

	if e.Key() == int(core.Qt__Key_Escape) {
		if c.drag {
			c.Scene.RemoveItem(c.drawingline)
			c.drawingline = nil
			c.drag = false
			c.View.Viewport().Repaint()
		}
	}
}

func (c *Canvas) ClearScene() {
	c.Scene.Clear()
	c.View.SetScene(c.Scene)
	c.View.Show()

}

func (c *Canvas) ShowPic(filepath, filetype string) {

	ir := gui.NewQImageReader3(filepath, filetype)
	img := ir.Read()

	pix := gui.QPixmap_FromImage(img, 0)

	c.Scene.Clear()
	c.Scene.AddPixmap(pix)

	c.View.SetScene(c.Scene)
	c.View.Show()

}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func (c *Canvas) ListNodes() {
	c.Controlbox.List.Clear()

	for _, f := range c.Figures.Values() {
		i := f.(*rectangle.Rectangle)
		c.Controlbox.List.AddItem("node: " + FloatToString(i.X) + " " + FloatToString(i.Y))
	}
}

//-------------------------------------------------------------------------Notes

//pix := gui.NewQPixmap4(FilePath, "", 0)

//pix := gui.QPixmap.ConvertFromImage(img,0) // conversion flags 0

//	f := gui.NewQFont2("courier", 10, 1, false)
//	t := c.Scene.AddSimpleText("Wish you were here", f)
//	t.SetPos2(x, y)

//			for f, n := range c.Graph.Nodes { // index, node
//				if (n.Y <= n.Y+10) && (n.X <= n.X+20) {
//					var i = c.Scene.ItemAt3(x, y, nil)
//					c.Scene.RemoveItem(i)
//					c.Graph.RemoveNode(f)
//				} else {
//					r := NewRectangle(x, y, 20, 10, c.Scene)
//					c.Scene.AddItem(r)
//					n := c.Graph.AddNode(x, y)
//					c.Statusbar.ShowMessage(FloatToString(n.X)+" "+FloatToString(n.Y), 0)
//				}
//			}

// graph lookup
//c.Graph.Nodes[x,y]
//			var i = c.Scene.ItemAt3(x, y, nil)

//			if i != nil {
//				r := NewRectangle(x, y, 20, 10, c.Scene)
//				c.Scene.AddItem(r)
//				n := c.Graph.AddNode(x, y)
//				c.Statusbar.ShowMessage(FloatToString(n.X)+" "+FloatToString(n.Y), 0)
//			} else {
//				c.Statusbar.ShowMessage("Item found", 0)
//			}

//l.Update(core.NewQRectF4(l.SX, l.SY, l.TX, l.TY))
//c.Scene.Update(core.NewQRectF4(l.SX/2, l.SY/2, l.TX/2, l.TY/2))
//c.Scene.Update(c.Scene.SceneRect())

//func RemoveItem(a treemap.Map, n gfxinterface.Figure) {

//	//	for i, v := range a { // to remove 1 item from slice wtf?!
//	//		if v == n {
//	//			a[i] = a[len(a)-1]
//	//			a[len(a)-1] = nil
//	//			a = a[:len(a)-1]
//	//		}
//	//	}

//}

//c.Scene.SceneRect() core.QRectF
//c.View.Viewport().Repaint() // update()

//	for _, f := range c.Figures {
//		i := f.(*rectangle.Rectangle)
//		c.Controlbox.List.AddItem("node: " + FloatToString(i.X) + " " + FloatToString(i.Y))
//	}

// erase rectangle -- right click
//					c.Statusbar.ShowMessage("remove item: "+FloatToString(n[0].X)+" "+FloatToString(n[0].Y), 0)
//					c.Scene.RemoveItem(n[0].QGraphicsItem)
//					c.Graph.Nodes = c.Graph.RemoveNode(n[0])
