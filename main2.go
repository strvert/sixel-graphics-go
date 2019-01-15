package main2

import (
	"github.com/soniakeys/quant/median"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	_ "image/png"
	"os"
)

func main() {
	reader, err := os.Open(os.Args[1])
	if err != nil {
		return
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		return
	}

	q := median.Quantizer(256)
	p := q.Quantize(make(color.Palette, 0, 256), img)
	paletted := image.NewPaletted(img.Bounds(), p)
	draw.FloydSteinberg.Draw(paletted, img.Bounds(), img, image.ZP)

	f, _ := os.Create("median-floyd-steinberg.gif")
	defer f.Close()

	opts := &gif.GIF{
		Image:     []*image.Paletted{paletted},
		Delay:     []int{0},
		LoopCount: 0,
	}
	gif.EncodeAll(f, opts)
}
