package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	width  = 640.0
	height = 640.0
)

var (
	maze = [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1},
		{1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1},
		{1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1},
		{1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	oneBlockWidthLength  = width / len(maze[0])
	oneBlockHeightLength = height / len(maze)
)

type Point struct {
	x, y float64
}

type Game struct {
	pos  Point
	dir  Point
	maze *ebiten.Image
}

func isOnScrean(x, y int) bool {
	return x < 0 || x > width || y < 0 || y > height
}
func DrawLineDDA(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	if math.Abs(float64(x2-x1)) >= math.Abs(float64(y2-y1)) {
		if x2 < x1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(y2-y1) / float64(x2-x1)
		for x, y := x1, float64(y1)+0.5; x <= x2; x, y = x+1, y+k {

			if isOnScrean(x, int(y)) || maze[int(y)/int(oneBlockHeightLength)][x/int(oneBlockWidthLength)] != 0 {
				return
			}
			img.Set(x, int(y), c)
		}
	} else {
		if y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(x2-x1) / float64(y2-y1)
		for x, y := float64(x1)+0.5, y1; y <= y2; x, y = x+k, y+1 {
			if isOnScrean(int(x), y) || maze[y/int(oneBlockHeightLength)][int(x)/int(oneBlockWidthLength)] != 0 {
				return
			}
			img.Set(int(x), y, c)
		}
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		a := int(g.pos.x+g.dir.x) / int(oneBlockWidthLength)
		b := int(g.pos.y+g.dir.y) / int(oneBlockHeightLength)
		if maze[b][a] == 0 {
			g.pos.x += g.dir.x
			g.pos.y += g.dir.y
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		a := int(g.pos.x-g.dir.x) / int(oneBlockWidthLength)
		b := int(g.pos.y-g.dir.y) / int(oneBlockHeightLength)
		if maze[b][a] == 0 {
			g.pos.x -= g.dir.x
			g.pos.y -= g.dir.y
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		tmp := Rotate(g.dir, -math.Pi/2)
		a := int(g.pos.x+tmp.x) / int(oneBlockWidthLength)
		b := int(g.pos.y+tmp.y) / int(oneBlockHeightLength)
		if maze[b][a] == 0 {
			g.pos.x += tmp.x
			g.pos.y += tmp.y
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		tmp := Rotate(g.dir, math.Pi/2)
		a := int(g.pos.x+tmp.x) / int(oneBlockWidthLength)
		b := int(g.pos.y+tmp.y) / int(oneBlockHeightLength)
		if maze[b][a] == 0 {
			g.pos.x += tmp.x
			g.pos.y += tmp.y
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.dir = Rotate(g.dir, -math.Pi/180)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.dir = Rotate(g.dir, math.Pi/180)
	}

	return nil
}
func Rotate(a Point, angle float64) Point {
	a.x, a.y = a.x*math.Cos(angle)-a.y*math.Sin(angle), a.x*math.Sin(angle)+a.y*math.Cos(angle)
	return a
}

func DrawMap(screen *ebiten.Image) {
	for i := range maze {
		for j, WallType := range maze[i] {
			if WallType != 0 {
				ebitenutil.DrawRect(screen, float64(j*oneBlockWidthLength), float64(i*oneBlockHeightLength), float64(oneBlockWidthLength), float64(oneBlockHeightLength), color.RGBA{255, 0, 0, 255})
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.DrawImage(g.maze, nil)
	ebitenutil.DrawCircle(screen, g.pos.x, g.pos.y, 3, color.RGBA{255, 255, 0, 255})

	for i := -30.0; i < 31; i += 0.5 {
		tmp := Rotate(g.dir, i*math.Pi/180)
		DrawLineDDA(screen, int(g.pos.x), int(g.pos.y), int(g.pos.x+tmp.x*1000), int(g.pos.y+tmp.y*1000), color.RGBA{255, 255, 0, 255})
	}

	// ebitenutil.DrawRect(screen, float64(a*oneBlockWidthLength), float64(b*oneBlockHeightLength), float64(oneBlockWidthLength), float64(oneBlockHeightLength), color.RGBA{0, 0, 255, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	a := ebiten.NewImage(width, height)
	DrawMap(a)
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{Point{320, 350}, Point{0, -1}, a}); err != nil {
		log.Fatal(err)
	}
}
