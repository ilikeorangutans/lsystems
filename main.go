package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

func loadFont() *ttf.Font {
	var fontPath string
	if runtime.GOOS == "darwin" {
		fontPath = "/Library/Fonts/Verdana.ttf"
	} else if runtime.GOOS == "linux" {
		fontPath = "/usr/share/fonts/truetype/ttf-bitstream-vera/Vera.ttf"
	}
	font, err := ttf.OpenFont(fontPath, 12)
	if err != nil {
		panic(err)
	}

	return font
}

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()
	ttf.Init()
	defer ttf.Quit()

	//font := loadFont()

	viewport := &sdl.Rect{0, 0, 800, 600}

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int(viewport.W), int(viewport.H), sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, 0, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	input := "FX"
	var productions Productions = make(map[rune]string)
	productions['X'] = "X+YF+"
	productions['Y'] = "-FX-Y"
	productions['F'] = ""

	result := productions.ApplyTimes(input, 10)

	x := 0
	y := 0
	angle := 0
	stepsize := 5

	timer := time.NewTicker(time.Second * 1)
	defer timer.Stop()

	running := true
	var event sdl.Event
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false

			case *sdl.KeyDownEvent:
				switch t.Keysym.Sym {
				case sdl.K_ESCAPE:
					running = false
				case sdl.K_LEFT:
					viewport.X -= 10
				case sdl.K_RIGHT:
					viewport.X += 10
				case sdl.K_UP:
					viewport.Y -= 10
				case sdl.K_DOWN:
					viewport.Y += 10
				case sdl.K_KP_MINUS:
					stepsize--
					if stepsize < 1 {
						stepsize = 1
					}
				case sdl.K_KP_PLUS:
					stepsize++

				}

			case *sdl.WindowEvent:
				if t.Event == sdl.WINDOWEVENT_RESIZED {
					viewport.W = t.Data1
					viewport.H = t.Data2
					log.Printf("Window resized")
				}
			}
		}

		start := time.Now()
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()

		angle = 0
		x = 0
		y = 0

		for i := range result {
			cur := result[i]

			dx := 0
			dy := 0
			switch angle {
			case 0:
				dy = -stepsize
			case 90:
				dx = stepsize
			case 180:
				dy = stepsize
			case 270:
				dx = -stepsize
			}

			newX := 0
			newY := 0

			switch cur {
			case '+':
				angle += 90
			case '-':
				angle -= 90

			default:
				newX = x + dx
				newY = y + dy

				if PointInRect(viewport, x, y) || PointInRect(viewport, newX, newY) {

					renderer.SetDrawColor(255, 255, 255, 255)
					renderer.DrawLine(x-int(viewport.X), y-int(viewport.Y), newX-int(viewport.X), newY-int(viewport.Y))
				}

				x = newX
				y = newY

			}

			angle = angle % 360
			if angle < 0 {
				angle += 360
			}
		}

		select {
		case <-timer.C:
			duration := time.Since(start)
			fmt.Printf("render took %s, view: %v\n", duration, viewport)
		default:
		}

		renderer.Present()
		sdl.Delay(33)
	}

}

func PointInRect(rect *sdl.Rect, x int, y int) bool {
	return rect.X < int32(x) && int32(x) < rect.X+rect.W && rect.Y < int32(y) && int32(y) < rect.Y+rect.H
}
