package main

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
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	url := "localhost:8000"

	http.HandleFunc("/", handler)
	http.HandleFunc("/lissajous", lissajous_handler)

	log.Println("listen in port :8000")
	log.Fatalln(http.ListenAndServe(url, nil))
}

func lissajous_handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	_, ok := r.Form["cycles"]
	cycles := 5.0
	if ok {
		v, _ := strconv.Atoi(r.Form["cycles"][0])
		cycles = float64(v)

	}

	log.Println(r.Form["cycles"])
	lissajous(w, cycles)
}

func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintln(w, "Header [", k, "] = ", v)
	}

	fmt.Fprintln(w, "Host:", r.Host)
	fmt.Fprintln(w, "RemoteAddr", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintln(w, "Form[", k, "]", "=", v)
	}
}
