package snake

import (
	"image"
	"math/rand"
	"time"

	"github.com/marcusolsson/tui-go"
)

type Block struct {
	x int
	y int
}

type Game struct {
	snake      Snake
	block      Block
	width      int
	height     int
	clockSpeed time.Duration
	*tui.Box
}

const blockSize = 1

func NewGame() (g Game) {
	g.Box = tui.NewHBox()
	g.Box.SetBorder(true)

	g.clockSpeed = 100 * time.Millisecond

	return g
}

func (g *Game) Draw(p *tui.Painter) {
	// Do the default stuff
	g.Box.Draw(p)

	// Draw snake
	for idx := 0; idx < len(g.snake.elements); idx += 2 {
		p.DrawRune(g.snake.elements[idx], g.snake.elements[idx+1], '.')
	}

	// Draw block
	p.DrawRune(g.block.x, g.block.y, 'X')
}

func (g *Game) Resize(size image.Point) {
	if (g.width != size.X) || (g.height != size.Y) {
		g.initialize(size.X, size.Y)
	}

	g.Box.Resize(size)
}

func (g *Game) Start(ui tui.UI) {
	g.setKeyBindings(ui)

	for {
		time.Sleep(g.clockSpeed)

		if g.block.isAt(g.snake.front()) {
			g.snake.move(true)
			g.generateBlock()
		} else {
			g.snake.move(false)
		}

		if g.snake.hasSelfCollision() {
			g.lose(ui)
			break
		}

		if g.isOutOfBounds(g.snake.front()) {
			g.lose(ui)
			break
		}

		ui.Repaint()
	}
}

func (g *Game) initialize(w, h int) {
	g.width = w
	g.height = h

	x := g.width / 2
	y := g.height / 2
	g.snake = Snake{
		elements:  []int{x, y, x + 1, y},
		direction: "right",
	}

	g.generateBlock()
}

func (g *Game) setKeyBindings(ui tui.UI) {
	ui.SetKeybinding("Up", func() { g.snake.SetDirection("up") })
	ui.SetKeybinding("Down", func() { g.snake.SetDirection("down") })
	ui.SetKeybinding("Left", func() { g.snake.SetDirection("left") })
	ui.SetKeybinding("Right", func() { g.snake.SetDirection("right") })
}

func (g Game) isOutOfBounds(x, y int) bool {
	if x < 0 || x >= g.width {
		return true
	}

	if y < 0 || y >= g.height {
		return true
	}

	return false
}

func (g *Game) generateBlock() {
	g.block = Block{
		x: rand.Intn(g.width),
		y: rand.Intn(g.height),
	}
}

func (g *Game) lose(ui tui.UI) {
	ui.Quit()
	// ui.Repaint()
}

func (b Block) isAt(x, y int) bool {
	return b.x == x && b.y == y
}
