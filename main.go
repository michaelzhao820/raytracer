package main

import (
	. "github.com/michaelzhao820/raytracer/raytracer"
	"log"
	"math"
)

func main() {
	world := NewWorld()

	// Light source
	world.SetLight(&Light{
		Position:  NewPoint(-10, 10, -10),
		Intensity: NewColor(1, 1, 1),
	})

	// Floor
	floor := NewSphere()
	tfm, _ := ScalingMatrix(10, 0.01, 10)
	floor.SetTransform(tfm)
	floor.GetMaterial().SetColor(1, 0.9, 0.9)
	floor.GetMaterial().SetSpecular(0)

	// Left Wall
	leftWall := NewSphere()
	leftTfm, _ := TranslationMatrix(0, 0, 5)
	ryLeft, _ := RotationYMatrix(-math.Pi / 4)
	rx, _ := RotationXMatrix(math.Pi / 2)
	s, _ := ScalingMatrix(10, 0.01, 10)

	lWallTfm, _ := leftTfm.MultiplyMatrices(ryLeft)
	lWallTfm, _ = lWallTfm.MultiplyMatrices(rx)
	lWallTfm, _ = lWallTfm.MultiplyMatrices(s)

	leftWall.SetTransform(lWallTfm)
	leftWall.SetMaterial(floor.GetMaterial())

	// Right Wall
	rightWall := NewSphere()
	rightTfm, _ := TranslationMatrix(0, 0, 5)
	ryRight, _ := RotationYMatrix(math.Pi / 4)

	rWallTfm, _ := rightTfm.MultiplyMatrices(ryRight)
	rWallTfm, _ = rWallTfm.MultiplyMatrices(rx)
	rWallTfm, _ = rWallTfm.MultiplyMatrices(s)

	rightWall.SetTransform(rWallTfm)
	rightWall.SetMaterial(floor.GetMaterial())

	// Middle Sphere
	middle := NewSphere()
	mtm, _ := TranslationMatrix(-0.5, 1, 0.5)
	middle.SetTransform(mtm)
	middle.GetMaterial().SetColor(0.1, 1, 0.5)
	middle.GetMaterial().SetDiffuse(0.7)
	middle.GetMaterial().SetSpecular(0.3)

	// Right Sphere
	right := NewSphere()
	scl, _ := ScalingMatrix(0.5, 0.5, 0.5)
	trn, _ := TranslationMatrix(1.5, 0.5, -0.5)
	rTx, _ := trn.MultiplyMatrices(scl)
	right.SetTransform(rTx)
	right.GetMaterial().SetColor(0.5, 1, 0.1)
	right.GetMaterial().SetDiffuse(0.7)
	right.GetMaterial().SetSpecular(0.3)

	// Left Sphere
	left := NewSphere()
	scl, _ = ScalingMatrix(0.33, 0.33, 0.33)
	trn, _ = TranslationMatrix(-1.5, 0.33, -0.75)
	lTx, _ := trn.MultiplyMatrices(scl)
	left.SetTransform(lTx)
	left.GetMaterial().SetColor(1, 0.8, 0.1)
	left.GetMaterial().SetDiffuse(0.7)
	left.GetMaterial().SetSpecular(0.3)

	// Add objects to world
	world.AddObject(floor)
	world.AddObject(leftWall)
	world.AddObject(rightWall)
	world.AddObject(middle)
	world.AddObject(right)
	world.AddObject(left)

	// Camera
	camera := NewCamera(800, 400, math.Pi/3)
	camera.SetTransform(ViewTransform(
		NewPoint(0, 1.5, -5),
		NewPoint(0, 1, 0),
		NewVector(0, 1, 0),
	))

	if err := camera.Render(*world); err != nil {
		log.Fatalf("Render failed: %v", err)
	}
}
