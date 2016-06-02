package gfxinterface

import "local/AM/graph"

type Figure interface {
	GetX() float64
	SetX(x float64)
	GetY() float64
	SetY(y float64)
	GetNode() *graph.Node
	AddEdgeIncoming(l Link)
	AddEdgeOutgoing(l Link)
	//	RemoveEdge()
}

type Link interface {
	GetEdge() *graph.Edge
}
