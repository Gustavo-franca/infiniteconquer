package position

type Value struct {
	x int64
	y int64
}

func New(x int64, y int64) Value {
	return Value{x: x, y: y}
}

func (v Value) Get() (x int64, y int64) {
	return v.x, v.y
}
