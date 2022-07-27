// https://blog.golang.org/go-imagedraw-package

package main

import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "os"
)

func main() {

    new_png_file := "./tmp/two_rectangles1.png" // output image will live here

    myimage := image.NewRGBA(image.Rect(0, 0, 220, 220)) // x1,y1,  x2,y2 of background rectangle
    // mygreen := color.RGBA{0, 100, 0, 255}  //  R, G, B, Alpha
    mywhite := color.RGBA{0, 0, 0, 0}  //  R, G, B, Alpha
    

    // backfill entire background surface with color mygreen
    draw.Draw(myimage, myimage.Bounds(), &image.Uniform{mywhite}, image.ZP, draw.Src)

    red_rect := image.Rect(110, 0, 220, 220) //  geometry of 2nd rectangle which we draw atop above rectangle
    myred := color.RGBA{200, 0, 0, 255}

    // create a red rectangle atop the green surface
    draw.Draw(myimage, red_rect, &image.Uniform{myred}, image.ZP, draw.Src)

    myfile, err := os.Create(new_png_file)     // ... now lets save output image
    if err != nil {
        panic(err)
    }
    defer myfile.Close()
    png.Encode(myfile, myimage)   // output file /tmp/two_rectangles.png
}


// https://stackoverflow.com/questions/28992396/draw-a-rectangle-in-golang

// https://golang.hotexamples.com/examples/github.com.disintegration.gift/-/Crop/golang-crop-function-examples.html

// https://stackoverflow.com/questions/32544927/cropping-and-creating-thumbnails-with-go