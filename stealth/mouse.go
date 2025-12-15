package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// Point represents a 2D point
type Point struct {
	X float64
	Y float64
}

// cubicBezier calculates a Bezier curve point
func cubicBezier(t float64, p0, p1, p2, p3 Point) Point {
	u := 1 - t
	tt := t * t
	uu := u * u
	uuu := uu * u
	ttt := tt * t

	x := uuu*p0.X + 3*uu*t*p1.X + 3*u*tt*p2.X + ttt*p3.X
	y := uuu*p0.Y + 3*uu*t*p1.Y + 3*u*tt*p2.Y + ttt*p3.Y

	return Point{x, y}
}

// MoveMouseHumanLike simulates human-like mouse movement
func MoveMouseHumanLike(page *rod.Page, fromX, fromY, toX, toY float64) {
	p0 := Point{fromX, fromY}

	p1 := Point{
		fromX + rand.Float64()*100,
		fromY + rand.Float64()*100,
	}

	p2 := Point{
		toX - rand.Float64()*100,
		toY - rand.Float64()*100,
	}

	p3 := Point{toX, toY}

	steps := 25 + rand.Intn(15)

	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)
		point := cubicBezier(t, p0, p1, p2, p3)

		// micro jitter
		jitterX := rand.Float64()*2 - 1
		jitterY := rand.Float64()*2 - 1

		page.MustEval(`
			(x, y) => {
				const evt = new MouseEvent("mousemove", {
					clientX: x,
					clientY: y,
					bubbles: true
				});
				document.dispatchEvent(evt);
			}
		`, point.X+jitterX, point.Y+jitterY)

		time.Sleep(time.Duration(10+rand.Intn(30)) * time.Millisecond)
	}
}
