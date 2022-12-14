package utils

type Vec2 struct {
	X int
	Y int
}

func NewVec2(x int, y int) *Vec2 {
	return &Vec2{X: x, Y: y}
}

func (v *Vec2) Equals(vec Vec2) bool {
	return v.X == vec.X && v.Y == vec.Y
}

func (v *Vec2) AddTo(vec Vec2) {
	v.X += vec.X
	v.Y += vec.Y
}

func (v *Vec2) Copy() *Vec2 {
	return &Vec2{X: v.X, Y: v.Y}
}

func (v *Vec2) Add(vec Vec2) Vec2 {
	return Vec2{X: (v.X + vec.X), Y: (v.Y + vec.Y)}
}
