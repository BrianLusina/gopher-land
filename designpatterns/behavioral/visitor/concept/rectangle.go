package concept

type rectangle struct {
	width  int
	height int
}

func (r *rectangle) accept(v visitor) {
	v.visitForRectangle(r)
}

func (r *rectangle) getType() string {
	return "Rectangle"
}
