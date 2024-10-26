package geometry

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Scale(factor float64) (float64, float64) {
	r.width *= factor
	r.height *= factor
	return r.width, r.height
}
