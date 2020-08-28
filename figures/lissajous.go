// Defines function that writes out a gif of a Lissajous curve.
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
)

type Parameters struct {
	Cycles int     `json:cycles`
	Freq   float64 `json:freq`
}

func CheckFixParams(ptr *Parameters) {
	if ptr.Cycles <= 0 {
		ptr.Cycles = 3
	}
	if ptr.Freq <= 0 {
		ptr.Freq = 5
	}
}

var palette = []color.Color{color.White, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func DrawGIF(out io.Writer, p Parameters) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   //image canvas covers [-size..)size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(p.Cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(p.Freq*t + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
