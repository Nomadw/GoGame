package main

import (
	"azul3d.org/engine/gfx"
	math "azul3d.org/engine/lmath"
	"azul3d.org/engine/gfx/gfxutil"
	"log"
)

type BaseObject struct {
	position, rotation, scale math.Vec3
	vertices []gfx.Vec3
	texturecoords []gfx.TexCoordSet
	name, shader, texture string
}

func NewBase(scale float64, shader, texture string) *BaseObject {
	o := new(BaseObject)
	o.position = math.Vec3{0, 0, 0}
	o.rotation = math.Vec3{0, 0, 0}
	o.scale = math.Vec3{scale, scale, scale}
	o.shader = shader
	o.texture = texture
	o.vertices = []gfx.Vec3{
		{-1, 0, -1},
		{1, 0, -1},
		{-1, 0, 1},

		{-1, 0, 1},
		{1, 0, -1},
		{1, 0, 1},
	}
	o.texturecoords = []gfx.TexCoordSet{
		{
			Slice: []gfx.TexCoord{
				{0, 1},
				{1, 1},
				{0, 0},

				{0, 0},
				{1, 1},
				{1, 0},
			},
		},
	}
	return o
}

func InitBase(object BaseObject) *gfx.Object {
	o := gfx.NewObject()
	o.SetPos(object.position)
	o.SetRot(object.rotation)
	o.SetScale(object.scale)
	shader, err := gfxutil.OpenShader(object.shader)
	if err != nil {
		log.Fatal(err)
	}
	o.Shader = shader
	tex, err := gfxutil.OpenTexture(object.texture)
	if err != nil {
		log.Fatal(err)
	}
	tex.Format = gfx.DXT1RGBA
	o.Textures = []*gfx.Texture{tex}
	o.OcclusionTest = true
	o.State = gfx.NewState()
	o.State.FaceCulling = gfx.NoFaceCulling
	mesh := gfx.NewMesh()
	mesh.Vertices = object.vertices
	mesh.TexCoords = object.texturecoords
	o.Meshes = []*gfx.Mesh{mesh}
	return o
}


