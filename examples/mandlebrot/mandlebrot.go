package main

import (
	"fmt"
	"math"
	"github.com/japanoise/wuhoo-go/software"
)

func mandlebrotDraw(b *wuhoo.Bitmap) {
	var factorX float64 = -2.5
	var factorY float64 = -2
	var zoom float64 = 4.0

	var Cr float64
	var Ci float64
	var pixelX float64
	var pixelY float64
	iter := 20

	var x, y wuhoo.Size

	width := b.Width()
	height := b.Height()

	for x = 0; x < width; x++ {
		for y = 0; y < height; y++ {
			pixelX = (float64(x) - 0.5)/float64(width)
			pixelY = (float64(y) - 0.5)/float64(height)

			pixelX = pixelX * zoom + factorX
			pixelY = pixelY * zoom + factorY

			Cr = pixelX
			Ci = pixelY

			var Zr float64 = Cr
			var Zi float64 = Ci

			var i int

			for i = 0; i < iter; i++ {
				var px float64 = (Zr * Zr - Zi * Zi) + Cr
				var py float64 = (Zi * Zr + Zr * Zi) + Ci

				Zr = px
				Zi = py

				if (px * px + py * py) > 20.0 {
					break
				}
			}

			col := float64(i) + 1 - math.Log(math.Log(
				math.Sqrt(Zr*Zr + Zi * Zi))) / math.Log(2.0)
			col = col/float64(iter)

			var out byte

			if i == iter {
				out = 255
			} else {
				out = byte(col*255.0)
			}

			b.SetXY(int(x), int(y),
				wuhoo.NewColorRGB(out, out, out))
		}
	}
}

func main() {
	var window wuhoo.Window
	var event wuhoo.Event
	event.Init()
	var bitmap wuhoo.Bitmap

	title := fmt.Sprintf("Mandlebrot Set (%s)", wuhoo.PlatformAPIString)
	running := true
	var posX wuhoo.Position = 300
	var posY wuhoo.Position = 300
	var width wuhoo.Size = 512
	var height wuhoo.Size = 512

	window.Init()
	window.Create(posX, posY, width, height, title,
		wuhoo.FlagTitled|wuhoo.FlagCanvas|
			wuhoo.FlagMouseCapture|
			wuhoo.FlagResizeable|
			wuhoo.FlagClientRegion|wuhoo.FlagCloseable)

	window.Show()

	bitmap.InitBlank(width, height)

	for running {
		window.EventNext(&event)

		switch event.Type() {
		case wuhoo.EventTypeWindow:
			ew := event.Window()

			switch ew.State() {
			case wuhoo.WindowStateInvalidated:
				mandlebrotDraw(&bitmap)

				window.Blit(&bitmap, 0, 0, 0, 0, width, height)
			case wuhoo.WindowStateClosed:
				running = false
			case wuhoo.WindowStateResized:
				window.RegionGet(&posX, &posY, &width, &height);
				fmt.Printf("|Mandelbrot| Resize Event Window Area Size [%d %d] x [%d %d].\n", posX, posY, width, height);
				window.ClientRegionGet(&posX, &posY, &width, &height);
				fmt.Printf("|Mandelbrot| Resize Event Window Client Area Size [%d %d] x [%d %d].\n", posX, posY, width, height);
				bitmap.Delete()

				bitmap.InitBlank(width, height)

				// sprintf(title_buffer, "Window Size (%d, %d)", width, height);
				// WuhooWindowSetTitle(&window, title_buffer);

				window.Blit(&bitmap, 0, 0, 0, 0, width, height)
			}
		case wuhoo.EventTypeKey:
			ek := event.Key()
			if ek.Code() == wuhoo.KeyEscape {
				running = (ek.State() != wuhoo.KeyStateUp)
			}
		}
	}

	bitmap.Delete()
	event.Destroy()
	window.Destroy()
}
