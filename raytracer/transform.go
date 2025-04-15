package raytracer

func ViewTransform(from, to, up Tuple) Matrix {
	forward, _ := to.Subtract(from)
	forward, _ = forward.Normalize()

	upn, _ := up.Normalize()
	left, _ := Cross(forward, upn)
	trueUp, _ := Cross(left, forward)

	orientation := NewMatrix(4, 4)

	orientation, _ = orientation.Set(0, 0, left[X])
	orientation, _ = orientation.Set(0, 1, left[Y])
	orientation, _ = orientation.Set(0, 2, left[Z])
	orientation, _ = orientation.Set(0, 3, 0)

	orientation, _ = orientation.Set(1, 0, trueUp[X])
	orientation, _ = orientation.Set(1, 1, trueUp[Y])
	orientation, _ = orientation.Set(1, 2, trueUp[Z])
	orientation, _ = orientation.Set(1, 3, 0)

	orientation, _ = orientation.Set(2, 0, -forward[X])
	orientation, _ = orientation.Set(2, 1, -forward[Y])
	orientation, _ = orientation.Set(2, 2, -forward[Z])
	orientation, _ = orientation.Set(2, 3, 0)

	orientation, _ = orientation.Set(3, 0, 0)
	orientation, _ = orientation.Set(3, 1, 0)
	orientation, _ = orientation.Set(3, 2, 0)
	orientation, _ = orientation.Set(3, 3, 1)

	translate, _ := TranslationMatrix(-from[X], -from[Y], -from[Z])
	viewTransform, _ := orientation.MultiplyMatrices(translate)
	return viewTransform
}
