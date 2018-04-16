package main

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	winW = 800
	winH = 800
)

func main() {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Hi", 200, 200, winW, winH, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	window.GLCreateContext()
	defer window.Destroy()

	gl.Init()
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Print("OpenGL version: ", version, "\n")
}
