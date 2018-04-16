package main

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
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

	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0,
	}

	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_COPY)

	vertexShaderSource :=
		`#version 330
     layout (location = 0) in vec3 aPos;

   void main()
   {
     gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
   }` + "\x00"

	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	csource, free := gl.Strs(vertexShaderSource)
	gl.ShaderSource(vertexShader, 1, csource, nil)
	free()
	gl.CompileShader(vertexShader)
	var status int32
	gl.GetShaderiv(vertexShader, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLenght int32
		gl.GetShaderiv(vertexShader, gl.INFO_LOG_LENGTH, &logLenght)
		log := strings.Repeat("\x00", int(logLenght+1))

	}
}
