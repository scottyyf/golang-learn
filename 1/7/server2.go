package main
// echo server

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex
var palette = []color.Color{color.White, color.Black, color.RGBA{231,222,124,245}}
const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
)
var count int

func main()  {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/test", handler2)
	http.HandleFunc("/img", handler3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request)  {
	mu.Lock() //获取锁
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", time.Now())
}

func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	count++
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func handler2(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header{
		fmt.Fprintf(w, "Header [%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	err := r.ParseForm()
	if err != nil{
		log.Print(err)
	}
	for k, v := range r.Form{
		fmt.Fprintf(w, "Form[%q = %q\n", k, v )
	}
}

func handler3(w http.ResponseWriter, r *http.Request)  {
	lissajous(w)

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
