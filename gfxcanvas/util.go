package gfxcanvas

import (
	"local/AM/gfxinterface"
	"local/AM/gfxobjects/rectangle"
)

func (c *Canvas) FindRectOverlap(x, y float64) gfxinterface.Figure {

	L2x, L2y := x, y
	R2x, R2y := x+20, y+10

	for _, f := range c.Figures.Values() {
		i := f.(*rectangle.Rectangle)

		L1x, L1y := i.X, i.Y
		R1x, R1y := i.X+20, i.Y+10

		if L2x < R1x && R2x > L1x && L2y < R1y && R2y > L1y {
			return i
		}
	}

	return nil
}

func (c *Canvas) FindRectBeneath(x, y float64) gfxinterface.Figure {

	for _, f := range c.Figures.Values() {
		i := f.(*rectangle.Rectangle)

		if x >= i.X && x <= i.X+20 && y >= i.Y && y <= i.Y+10 {
			return i
		}
	}

	return nil
}

func (c *Canvas) FindRectBeneath2(x, y, iX, iY float64) bool {

	if x >= iX+15 && x <= iX+20 && y >= iY && y <= iY+5 {
		return true
	} else {
		return false
	}

}
