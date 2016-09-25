package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

func renderMap(size Point, noize []Point, start Point, finish Point, path []Point) {
	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetLayout(buildLayout(size, noize, start, finish, path))

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func getSymbol(point Point, noizeMap map[string]int, start Point, finish Point, pathMap map[string]int) string {
	var key = point.String()
	if point == start {
		return "S"
	} else if point == finish {
		return "F"
	} else if _, ok := noizeMap[key]; ok {
		return "█"
	} else if _, ok := pathMap[key]; ok {
		return "x"
	} else {
		return " "
	}
}

func buildLayout(size Point, noize []Point, start Point, finish Point, path []Point) func(*gocui.Gui) error {
	noizeMap := Points(noize).ToMap()
	pathMap := Points(path).ToMap()

	return func(g *gocui.Gui) error {
		if v, err := g.SetView("hello", 0, 0, size.X+1, size.Y+1); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			var line string

			for i := 0; i < size.Y; i++ {
				line = ""
				for j := 0; j < size.X; j++ {
					line += getSymbol(Point{j, i}, noizeMap, start, finish, pathMap)
				}
				fmt.Fprintln(v, line)
			}
		}
		return nil
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
