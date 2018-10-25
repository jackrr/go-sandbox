package main

import (
	"fmt"

	"github.com/jackrr/visuals/snake"
	tui "github.com/marcusolsson/tui-go"
)

func main() {
	game := snake.NewGame()
	ui, err := tui.New(&game)
	if err != nil {
		fmt.Println(err)
	}

	t := tui.NewTheme()
	normal := tui.Style{Bg: tui.ColorBlack, Fg: tui.ColorWhite}
	t.SetStyle("normal", normal)
	ui.SetTheme(t)

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	go game.Start(ui)
	run(ui)
}

func run(ui tui.UI) {
	if err := ui.Run(); err != nil {
		fmt.Println(err)
	}
}
