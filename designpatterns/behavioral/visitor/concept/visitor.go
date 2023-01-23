package concept

type visitor interface {
	visitForCircle(*circle)
	visitForSquare(*square)
	visitForRectangle(*rectangle)
}
