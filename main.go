package main

import (
	"runtime"

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

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
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

	result := productions.ApplyTimes(input, 14)

	x := 400
	y := 300
	angle := 0
	stepsize := 5

	running := true
	var event sdl.Event
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false

			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				}

			}
		}

		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()

		angle = 0
		x = 400
		y = 300

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

				renderer.SetDrawColor(255, 255, 255, 255)
				renderer.DrawLine(x, y, newX, newY)
				x = newX
				y = newY
			}

			angle = angle % 360
			if angle < 0 {
				angle += 360
			}
		}

		renderer.Present()
		sdl.Delay(33)
	}

}
