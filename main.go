package main

import (
	"fmt"
	"strings"

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
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Print("OpenGL version: ", version, "\n")

	var status int32

	fragmentShaderSource := `
	#version 300 es
	precision mediump float;
 	out vec4 FragColor;

	void main() {
		FragColor = vec4(0.0, 0.0, 1.0, 1.0);
	}` + "\x00"

	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	csource, free := gl.Strs(fragmentShaderSource)
	gl.ShaderSource(fragmentShader, 1, csource, nil)
	free()
	gl.CompileShader(fragmentShader)
	gl.GetShaderiv(fragmentShader, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLenght int32
		gl.GetShaderiv(fragmentShader, gl.INFO_LOG_LENGTH, &logLenght)
		log := strings.Repeat("\x00", int(logLenght+1))
		gl.GetShaderInfoLog(fragmentShader, logLenght, nil, gl.Str(log))
		panic("Failed to compil\n" + log)
	}

	vertexShaderSource :=
		`#version 300 es
     layout (location = 0) in vec3 aPos;

   void main()
   {
     gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
   }` + "\x00"

	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	csource, free = gl.Strs(vertexShaderSource)
	gl.ShaderSource(vertexShader, 1, csource, nil)
	free()
	gl.CompileShader(vertexShader)
	gl.GetShaderiv(vertexShader, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLenght int32
		gl.GetShaderiv(vertexShader, gl.INFO_LOG_LENGTH, &logLenght)
		log := strings.Repeat("\x00", int(logLenght+1))
		gl.GetShaderInfoLog(vertexShader, logLenght, nil, gl.Str(log))
		panic("Failed to compil\n" + log)
	}

	var shaderProgram uint32
	var success int32
	shaderProgram = gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &success)

	if success == gl.FALSE {
		var logLenght int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLenght)
		log := strings.Repeat("\x00", int(logLenght+1))
		gl.GetProgramInfoLog(shaderProgram, logLenght, nil, gl.Str(log))
		panic("Failed to compil\n" + log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0,
	}

	var VAO uint32
	var VBO uint32

	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)

	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)
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

		gl.UseProgram(shaderProgram)
		gl.BindVertexArray(VAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		window.GLSwap()
	}
}
