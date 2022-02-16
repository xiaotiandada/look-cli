package main

import (
	"fmt"
	"log"
	"github.com/jroimartin/gocui"
)

type Wallpaper struct {
	url string
	name string
}

func main() {
		// 获取接口数据

		// 展示图片在终端
		g, err := gocui.NewGui(gocui.OutputNormal)
		if err != nil {
			log.Panicln(err)
		}
		defer g.Close()
	
		g.SetManagerFunc(layout)

		go fetchImages(g)
	
		if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
			log.Panicln(err)
		}
	
		if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
			log.Panicln(err)
		}

		// 可以交互？
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if _, err := g.SetView("side", 0, 0, int(0.16*float32(maxX)), maxY-6); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("main", int(0.16*float32(maxX + 4)), 0, maxX-1, maxY-6); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("cmdline", 0, maxY-5, maxX-1, maxY-1); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func fetchImages(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("side")
		if err != nil {
			return err
		}
		v.Clear()

		fmt.Fprintln(v, "Wallpaper")
		fmt.Fprintln(v, "version 0.0.1")
		fmt.Fprintln(v, "https://wallhaven.cc")
		return nil
	})
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("main")
		if err != nil {
			return err
		}
		v.Clear()

		wallpaper := Wallpaper {
			url: "https://w.wallhaven.cc/full/y8/wallhaven-y8vlyk.jpg",
			name: "Wallpaper name",
		}

		for i := 0; i <= 40; i++ {
			fmt.Fprintln(v, i, wallpaper.url)
		}

		return nil
	})
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("cmdline")
		if err != nil {
			return err
		}
		v.Clear()

		fmt.Fprintln(v, "^c: Exit")
		fmt.Fprintln(v, "Help: https://github.com/xiaotiandada")
		return nil
	})
}