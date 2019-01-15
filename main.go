package main

import (
    "os"
    "fmt"

    "gopkg.in/gographics/imagick.v3/imagick"
)

func main () {
    imagick.Initialize()
    defer imagick.Terminate()

    mw1 := imagick.NewMagickWand()
    defer mw1.Destroy()
    err := mw1.ReadImage(os.Args[1])
    if err != nil {
        panic(err)
    }


    mw1.SetFormat("PNG8")
    mw1.SetColorspace(imagick.COLORSPACE_RGB)
    mw1.PosterizeImage(8, imagick.DITHER_METHOD_NO)

    color_num := mw1.GetImageColors()
    width := mw1.GetImageWidth()
    height := mw1.GetImageHeight()

    var palettes []*imagick.PixelWand
    for i := 0; i < int(color_num); i++ {
        pix, _ := mw1.GetImageColormapColor(uint(i))
        palettes = append(palettes, pix)
    }

    for w := 0; w < int(width); w++ {
        for h := 0; h < int(height); h++ {
            pix, _ := mw1.GetImagePixelColor(w, h)
            fmt.Println(getIndex(pix, palettes))
        }
    }
    mw1.WriteImage("images/8bit.png")
}

func getIndex(pix *imagick.PixelWand, palette []*imagick.PixelWand) int {
    for i, v := range palette {
        fmt.Println(v.GetColorAsNormalizedString())
        if (pix.IsSimilar(v, 0.00000001)) {
            return i
        }
    }
    return 0
}
