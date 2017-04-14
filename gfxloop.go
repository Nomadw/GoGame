package main

import (
	_ "image/png"

	"azul3d.org/engine/gfx"
	"azul3d.org/engine/gfx/camera"
	"azul3d.org/engine/gfx/window"
	"azul3d.org/engine/keyboard"
	math "azul3d.org/engine/lmath"

	"fmt"
)

func gfxLoop(w window.Window, d gfx.Device) {

	props := w.Props()

	cam := camera.New(d.Bounds())
	cam.SetRot(math.Vec3{0,-90,80})
	cam.SetPos(math.Vec3{0, -40,0})


	BaseObject := NewBase(500, "shaders/seaplane", "textures/water.png")
	seaplane := InitBase(*BaseObject)

	events := make(chan window.Event, 256)

	w.Notify(events, window.KeyboardTypedEvents|window.FramebufferResizedEvents)

	for {
		window.Poll(events, func(e window.Event) {
			switch ev := e.(type) {
			case window.FramebufferResized:
				cam.Update(d.Bounds())

			case keyboard.Typed:
				switch ev.S {
				case "f":
					fmt.Println("toggle fullscreen")
					props.SetFullscreen(!props.Fullscreen())
					w.Request(props)

					// Insert Keyboard inputs here
				}
			}
		})

		var v math.Vec2
		// Depending on keyboard state, transform the triangle.
		kb := w.Keyboard()
		if kb.Down(keyboard.Escape) {
			w.Close()
		}
		if kb.Down(keyboard.ArrowLeft) {
			v.X -= 1
		}
		if kb.Down(keyboard.ArrowRight) {
			v.X += 1
		}
		if kb.Down(keyboard.ArrowDown) {
			v.Y -= 1
		}
		if kb.Down(keyboard.ArrowUp) {
			v.Y += 1
		}

		// Apply movement relative to the frame rate.
		v = v.MulScalar(d.Clock().Dt())

		// Update the triangle's transformation matrix.
		if kb.Down(keyboard.R) {
			// Reset transformation.
			oldParent := seaplane.Transform.Parent()
			seaplane.Transform.Reset()
			seaplane.Transform.SetParent(oldParent)
			cam.SetRot(math.Vec3{0,-90,80})
			cam.SetPos(math.Vec3{0, -40,0})

		}

		// Apply movement on X/Z axis.
		p := math.Vec3{v.X, 0, v.Y}
		if kb.Down(keyboard.LeftShift) {
			// Apply movement on X/Y axis.
			p = math.Vec3{v.X, v.Y, 0}
		}
		cam.SetPos(cam.Pos().Add(p.MulScalar(90)))

		// Clear color and depth buffers.
		d.Clear(d.Bounds(), gfx.Color{0.5, 0.9, 1, 1})
		d.ClearDepth(d.Bounds(), 1.0)

		// Draw the triangle to the screen.
		bounds := d.Bounds()
		d.Draw(bounds.Inset(50), seaplane, cam)

		// Render the whole frame.
		d.Render()
	}
}