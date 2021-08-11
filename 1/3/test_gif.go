package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0, 214, 133, 255}}

const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) { //out是io.Writer类型，可进行输出写入
	const (
		cycles  = 5
		res     = 0.01
		size    = 100
		nframes = 64
		delay   = 8 // 80ms
	)

	freq := rand.Float64() * 3.0        //随机数0-3
	anim := gif.GIF{LoopCount: nframes} //struct
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // image的大小
		img := image.NewPaletted(rect, palette)      // 实例化
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex) // 设置颜色
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay) // 延迟
		anim.Image = append(anim.Image, img) // 追加照片
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		return 
	} //EncodeAll将anim写入out
}
