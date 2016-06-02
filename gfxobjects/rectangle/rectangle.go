package rectangle

import (
	"local/AM/gfxinterface"
	"local/AM/graph"
	"strconv"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type Rectangle struct {
	*widgets.QGraphicsItem
	Type          string
	Node          *graph.Node
	IncomingEdges *treemap.Map //[]gfxinterface.Link
	OutgoingEdges *treemap.Map //[]gfxinterface.Link
	X, Y          float64
	W, H          float64
	over          bool
	Over2         bool
	fgcolor       Color
	bgcolor       Color
	statusbar     *widgets.QStatusBar
	cursor        *gui.QCursor
}

type Color struct{ R, G, B, A int }

func (r *Rectangle) GetX() float64        { return r.X }
func (r *Rectangle) SetX(x float64)       { r.X = x }
func (r *Rectangle) GetY() float64        { return r.Y }
func (r *Rectangle) SetY(y float64)       { r.Y = y }
func (r *Rectangle) GetNode() *graph.Node { return r.Node }

func (r *Rectangle) AddEdgeIncoming(l gfxinterface.Link) {
	e := l.GetEdge()
	r.IncomingEdges.Put(e.Id, l)
}

func (r *Rectangle) AddEdgeOutgoing(l gfxinterface.Link) {
	e := l.GetEdge()
	r.OutgoingEdges.Put(e.Id, l)
}

func NewRectangle(x, y, w, h float64, g *graph.Graph, sb *widgets.QStatusBar) (int, *Rectangle) {

	r := &Rectangle{
		QGraphicsItem: widgets.NewQGraphicsItem(nil),
		Type:          "Rectangle",
		Node:          g.AddNode(),
		IncomingEdges: treemap.NewWithIntComparator(),
		OutgoingEdges: treemap.NewWithIntComparator(),
		X:             x,
		Y:             y,
		W:             w,
		H:             h,
		fgcolor:       Color{0, 0, 0, 255},
		bgcolor:       Color{0, 0, 0, 100},
		statusbar:     sb,
		cursor:        gui.NewQCursor()}

	r.SetAcceptHoverEvents(true)
	r.ConnectHoverEnterEvent(r.hoverEnterEvent)
	r.ConnectHoverLeaveEvent(r.hoverLeaveEvent)

	//	r.ConnectMousePressEvent(r.mousePressEvent)
	//	r.ConnectMouseReleaseEvent(r.mouseReleaseEvent)

	//r.ConnectMouseMoveEvent(r.mouseMoveEvent)

	r.ConnectBoundingRect(r.boundingRect)
	r.ConnectPaint(r.paint)

	return r.Node.Id, r
}

func (r *Rectangle) paint(p *gui.QPainter, o *widgets.QStyleOptionGraphicsItem, w *widgets.QWidget) {

	var color = gui.NewQColor3(r.fgcolor.R, r.fgcolor.G, r.fgcolor.B, r.fgcolor.A) // r,g,b,a

	var pen = gui.NewQPen3(color)
	pen.SetWidth(1)

	if r.Over2 {
		pen.SetStyle(0)
	}

	p.SetRenderHint(1, true) // Antiailiasing
	var path = gui.NewQPainterPath()
	path.AddRoundedRect2(r.X, r.Y, r.W, r.H, 1, 1, 0) // Qt::AbsoluteSize
	p.SetPen(pen)
	p.DrawPath(path)

	if r.over {
		color = gui.NewQColor3(r.bgcolor.R, r.bgcolor.G, r.bgcolor.B, r.bgcolor.A) // r,g,b,a
		var brush = gui.NewQBrush3(color, 1)
		p.FillPath(path, brush)

		color = gui.NewQColor3(0, 0, 0, 60) // r,g,b,a
		var pen2 = gui.NewQPen2(0)          // no pen

		//		if r.Over2 {
		//			color = gui.NewQColor3(255, 255, 255, 255)
		//			pen2.SetStyle(1)
		//		}

		var brush2 = gui.NewQBrush3(color, 1)
		var path2 = gui.NewQPainterPath()

		p.SetPen(pen2)
		path2.AddRoundedRect2(r.X+15, r.Y, 5, 5, 1, 1, 0)
		p.FillPath(path2, brush2)
		p.DrawPath(path2)
	}

	// if selected {}

	//p.DrawPath(path)
}

func (r *Rectangle) boundingRect() *core.QRectF {
	return core.NewQRectF4(r.X, r.Y, r.W, r.H)
}

func (r *Rectangle) hoverEnterEvent(e *widgets.QGraphicsSceneHoverEvent) {
	r.over = true
	r.Update(core.NewQRectF4(r.X, r.Y, r.W, r.H))
}

func (r *Rectangle) hoverLeaveEvent(e *widgets.QGraphicsSceneHoverEvent) {
	r.over = false
	r.Update(core.NewQRectF4(r.X, r.Y, r.W, r.H))
}

//func (r *Rectangle) mousePressEvent(e *widgets.QGraphicsSceneMouseEvent) {

//	var x, y, iX, iY = e.ScenePos().X(), e.ScenePos().Y(), r.X, r.Y

//	if x >= iX+15 && x <= iX+20 && y >= iY && y <= iY+5 {
//		r.Over2 = true
//	}

//}

//func (r *Rectangle) mouseReleaseEvent(e *widgets.QGraphicsSceneMouseEvent) {
//	//r.Over2 = false
//}

//func (r *Rectangle) mouseMoveEvent(e *widgets.QGraphicsSceneMouseEvent) {
//	//r.MouseMoveEventDefault(e)

//}

func (r *Rectangle) FindRectBeneath(x, y float64) bool {
	var iX, iY = r.X, r.Y
	if x >= iX+15 && x <= iX+20 && y >= iY && y <= iY+5 {
		return true
	} else {
		return false
	}

}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

//------------------------------------------------------------------------------

//func (r *Rectangle) mousePressEvent(e *widgets.QGraphicsSceneMouseEvent) {
//	if e.Button() == 2 { // right click
//		r.canvas.Scene.RemoveItem(r.QGraphicsItem)
//		//r.canvas.
//		// Lines, Edges, Connectors -- Rectangles, Nodes, Figures

//		// delete (image, rectangle, node), (connectors, edges, lines)

//	}
//}

//r.SetFlag(2, true) // selectable
//r.ConnectMousePressEvent(r.mousePressEvent)
//	r.ConnectMouseReleaseEvent(r.mouseReleaseEvent)
