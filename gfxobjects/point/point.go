package point

import (
	"local/AM/gfxinterface"
	"local/AM/graph"
)

type Point struct {
	X, Y float64
}

func (p *Point) GetX() float64        { return p.X }
func (p *Point) SetX(x float64)       { p.X = x }
func (p *Point) GetY() float64        { return p.Y }
func (p *Point) SetY(y float64)       { p.Y = y }
func (p *Point) GetNode() *graph.Node { return nil }

func (p *Point) AddEdgeIncoming(edge gfxinterface.Link) {}

func (p *Point) AddEdgeOutgoing(edge gfxinterface.Link) {}
