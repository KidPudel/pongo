package customTypes

type Vector2[U int | float64] struct {
	X U
	Y U
}

func (v Vector2[U]) Add(anotherV Vector2[U]) Vector2[U] {
	return Vector2[U]{v.X + anotherV.X, v.Y + anotherV.Y}
}
