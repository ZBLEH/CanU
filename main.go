package main

import (
	g "HelperCanU"
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	winW = 1000
	winH = 1000
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

	fmt.Print(g.GetVersion(), "\n")

	shaderProgram, err := g.NewShader("hello.vert", "hello.frag")
if err != nil {
	panic(err)
}
	vertices := []float32{
		0.5, 0.5, 0.0, 1.0, 1.0,
		0.5, -0.5, 0.0, 1.0, 0.0,
		-0.5, -0.5, 0.0, 0.0, 0.0,
		-0.5, 0.5, 0.0, 0.0, 1.0}


		indices := []uint32 {
			0,1,3,
			1,2,3}

	g.GenBindBuffer(gl.ARRAY_BUFFER)
	VAO := g.GenBindVertexArray()
	g.BufferDataFloat(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)
	g.GenBindBuffer(gl.ELEMENT_ARRAY_BUFFER)
	g.BufferDataInt(gl.ELEMENT_ARRAY_BUFFER, indices, gl.STATIC_DRAW)


	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(2*4))

	gl.BindVertexArray(0)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		gl.ClearColor(0.0, 0.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		shaderProgram.Use()
		g.BindVertexArray(VAO)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))

		window.GLSwap()
		shaderProgram.CheckShaderForChanges()
	}
}
