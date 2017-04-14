package main

import (
	"azul3d.org/engine/gfx/window"
)

const (
	camFar = 1000.0
	QuadNum = 36
)


func main() {
	props := window.NewProps()
	props.SetTitle("Seabed")
	props.SetSize(1600, 1080)
	props.FramebufferSize()
	//seaplane := seaPlaneinit()
	window.Run(gfxLoop, props)
}